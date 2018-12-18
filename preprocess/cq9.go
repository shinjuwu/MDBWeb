package preprocess

import (
	"MDBWeb/orm"
	"MDBWeb/tool"
	"encoding/json"
	"strconv"
	"strings"
)

type Fish_Kill_Info struct {
	Status       int    `json:"status"`       // 狀態代碼 0:魚沒死 1:魚死
	Fish_id      int    `json:"fish_id"`      // 魚的編id
	Fish_Type    int    `json:"fish_type"`    // 魚種資訊
	Win          int64  `json:"win"`          // 得分金額
	Feature_type string `json:"feature_type"` // 特殊事件類別
	Feature_id   string `json:"feature_id"`   // 特殊事件id
}

type SaveGameLog_Fish_Feature_Hit struct {
	Remain_odds float64 `json:"remain_odds"` // 沒用完的odds, 應該可以不用回傳(client目前沒用到)
	Feature_id  string  `json:"feature_id"`  // 特殊事件id

	Kill_info []Fish_Kill_Info `json:"kill_info"` // 給獎的魚id陣列
}

func ProcessCQ9Log() {
	tool.Log.Info("CQ9 Prepross routine start.")
	db := orm.MysqlDB()
	betClusterList := make([]orm.BetCluster, 0)
	err := db.Where("IsProcess = ?", 0).And("RoundID <> ?", "").Find(&betClusterList)
	if err != nil {
		tool.Log.Errorf("Get betcluster failed! Xorm inner failed! Error: %v", err)
		return
	}
	err = setProcessing(betClusterList)
	if err != nil {
		tool.Log.Errorf("setProcessing(betClusterList) failed, Error:%v", err)
		return
	}
	for _, v := range betClusterList {
		createPreprocessLog(&v)
		setProcessed(&v)
		tool.Log.Infof("PreLog process end, RoundID = %s", v.RoundID)
	}
	tool.Log.Info("CQ9 Prepross routine end.")
}

func ProcessCQ9LogByRoundID(betCluster *orm.BetCluster) error {
	tool.Log.Infof("Process CQ9 Log by RoundID : %s", betCluster.RoundID)
	logs := []orm.BetCluster{}
	logs = append(logs, *betCluster)
	err := setProcessing(logs)
	if err != nil {
		tool.Log.Errorf("setProcessing(betClusterList) failed, Error:%v", err)
		return err
	}
	err = createPreprocessLog(betCluster)
	if err != nil {
		return err
	}
	err = setProcessed(betCluster)
	if err != nil {
		return err
	}
	tool.Log.Infof("Process CQ9 Log by RoundID end, RoundID = %s", betCluster.RoundID)
	return nil
}

func createPreprocessLog(data *orm.BetCluster) error {
	//處理一般子彈
	err := processNoramalBullet(data)
	if err != nil {
		tool.Log.Errorf("ProcessNormalBullet and insert to db failed!, error=%v , betCluster Data = %v", err, data)
		return err
	}
	//處理特殊子彈
	err = processFeatureBullet(data)
	if err != nil {
		tool.Log.Errorf("processFeatureBullet and insert to db failed!,error=%v , betCluster Data = %v", err, data)
		return err
	}
	return nil
}

//處理一般子彈
func processNoramalBullet(data *orm.BetCluster) error {
	db := orm.MysqlDB()
	sql := "SELECT Bet,FeatureBet,FeatureType,FishType,SUM(Bet),SUM(Bet_Win),Count(Round),Process_Status FROM `gamelog_fish`" +
		" WHERE ClusterID=" + strconv.Itoa(int(data.ClusterID)) + " AND ServerID=" + strconv.Itoa(data.ServerID) + " AND(Process_Status=5 or Process_Status=13) GROUP BY Bet,FeatureBet,FeatureType,FishType,Process_Status"
	results, err := db.Query(sql)
	if err != nil {
		tool.Log.Errorf("Query fish log failed, Sql: %s    ,func:processNoramalBullet() , Error = %v", sql, err)
		return err
	}
	insertData := make([]orm.PreprocessLog, 0)
	for _, v := range results {
		queryData, err := checkNoemalQueryField(v)
		if err != nil {
			tool.Log.Errorf("String to int failed, Error = %v", err)
			return err
		}
		processLog := orm.PreprocessLog{
			ClusterID:     data.ClusterID,
			ServerID:      data.ServerID,
			RoundID:       data.RoundID,
			Bet:           queryData.Bet,
			FeatureBet:    queryData.FeatureBet,
			FeatureType:   queryData.FeatureType,
			FishType:      string(v["FishType"]),
			Result:        string(v["Result"]),
			TotalRound:    queryData.TotalRound,
			TotalBet:      queryData.TotalBet,
			TotalWin:      queryData.TotalWin,
			ProcessStatus: queryData.ProcessStatus,
			DisConTimes:   0,
			DisConSettle:  0,
		}

		processLog.FishID = processLog.FishType[1 : len(processLog.FishType)-1]
		insertData = append(insertData, processLog)
	}
	_, err = db.Insert(insertData)
	if err != nil {
		return err
	}
	return nil
}

//處理特殊子彈
func processFeatureBullet(data *orm.BetCluster) error {
	db := orm.MysqlDB()
	sql := "SELECT Bet,FeatureBet,FeatureType,FishType,Bet_Win,Round,Process_Status,Result FROM gamelog_fish WHERE ClusterID=" +
		strconv.Itoa(int(data.ClusterID)) + " AND ServerID=" + strconv.Itoa(data.ServerID) + " AND(Process_Status=12 or Process_Status=14)"
	results, err := db.Query(sql)
	if err != nil {
		tool.Log.Errorf("Query fish log failed, Sql: %s    ,func:processFeatureBullet() , Error = %v", sql, err)
		return err
	}
	for _, v := range results {
		queryData, err := checkFeatureQueryField(v)
		if err != nil {
			tool.Log.Errorf("String to int failed, Error = %v", err)
			return err
		}
		var disConnSettle int64
		var disConnCount int64
		if queryData.ProcessStatus == 12 {
			//斷線結清
			disConnCount = 1
			disConnSettle = queryData.TotalWin
		}
		processLog := orm.PreprocessLog{
			ClusterID:     data.ClusterID,
			ServerID:      data.ServerID,
			RoundID:       data.RoundID,
			Bet:           queryData.Bet,
			FeatureBet:    queryData.FeatureBet,
			FeatureType:   queryData.FeatureType,
			FishType:      string(v["FishType"]),
			Result:        string(v["Result"]),
			TotalRound:    queryData.TotalRound,
			TotalBet:      queryData.TotalBet,
			TotalWin:      queryData.TotalWin,
			ProcessStatus: queryData.ProcessStatus,
			DisConTimes:   disConnCount,
			DisConSettle:  disConnSettle,
		}
		if processLog.FishType == "" {
			if processLog.ProcessStatus == 12 {
				//斷線結清處理
				_, err = db.Insert(processLog)
				if err != nil {
					return err
				}
			}
			continue
		}
		processLog.FishType = processLog.FishType[1 : len(processLog.FishType)-1]
		fretureLogs := processFeatureLog(&processLog)
		if fretureLogs != nil {
			err := batchInsert(fretureLogs)
			if err != nil {
				tool.Log.Errorf("Feature Bullet batch insert failed!, error:%v  , data: %v", err, fretureLogs)
				return err
			}
		}

	}
	return nil
}

//統計results的魚得分資料
func processFeatureLog(log *orm.PreprocessLog) map[int]map[int]orm.PreprocessLog {
	result := SaveGameLog_Fish_Feature_Hit{}
	decodeResult := tool.DoZlibUnCompressGetString(log.Result)
	err := json.Unmarshal([]byte(decodeResult), &result)
	if err != nil {
		tool.Log.Errorf("Json Unmarshal failed! Error= %v , Result= %s", err, decodeResult)
		return nil
	}

	featureLogs := map[int]map[int]orm.PreprocessLog{}
	var fLogByFeatureBet map[int]orm.PreprocessLog //map[featureBet]orm.PreprocessLog
	featureLogs[log.FeatureBet] = make(map[int]orm.PreprocessLog)
	fLogByFeatureBet = make(map[int]orm.PreprocessLog)
	for _, v := range result.Kill_info {
		if v.Status == 1 {
			preLog := orm.PreprocessLog{}
			var ok bool
			if preLog, ok = fLogByFeatureBet[v.Fish_Type]; !ok {
				copyLog := *log
				copyLog.FishID = strconv.Itoa(v.Fish_Type)
				copyLog.TotalFeatureHit = 1
				copyLog.TotalWin = v.Win
				fLogByFeatureBet[v.Fish_Type] = copyLog
			} else {
				preLog.TotalFeatureHit = preLog.TotalFeatureHit + 1
				preLog.TotalWin = preLog.TotalWin + v.Win
				fLogByFeatureBet[v.Fish_Type] = preLog
			}
		}
	}
	featureLogs[log.Bet] = fLogByFeatureBet
	return featureLogs
}

func fishWinisFitFishType(result *SaveGameLog_Fish_Feature_Hit, log *orm.PreprocessLog) bool {
	fishWins := strings.Split(log.FishWin, ",")
	fishTypes := strings.Split(log.FishType, ",")
	if len(fishWins) != len(fishTypes) {
		return false
	}
	index := 0
	for _, v := range result.Kill_info {
		if v.Status == 1 {
			if v.Feature_type != fishTypes[index] {
				return false
			}
			if strconv.Itoa(int(v.Win)) != fishWins[index] {
				return false
			}
			index = index + 1
		}
	}
	return true
}

func batchInsert(featureLogs map[int]map[int]orm.PreprocessLog) error {
	db := orm.MysqlDB()
	insertData := make([]orm.PreprocessLog, 0)
	for _, v1 := range featureLogs {
		for _, log := range v1 {
			insertData = append(insertData, log)
		}
	}
	_, err := db.Insert(insertData)
	if err != nil {
		return err
	}
	return nil
}

func setProcessing(logs []orm.BetCluster) error {
	db := orm.MysqlDB()
	for _, v := range logs {
		sql := "UPDATE bet_cluster SET IsProcess=1 WHERE ClusterID=" + strconv.Itoa(int(v.ClusterID)) + " AND ServerID=" + strconv.Itoa(v.ServerID)
		_, err := db.Query(sql)
		if err != nil {
			tool.Log.Errorf("Set processed failed!, Error: %v , Sql: %s", err, sql)
			return err
		}
	}
	return nil
}

func setProcessed(log *orm.BetCluster) error {
	db := orm.MysqlDB()
	sql := "UPDATE bet_cluster SET IsProcess=2 WHERE ClusterID=" + strconv.Itoa(int(log.ClusterID)) + " AND ServerID=" + strconv.Itoa(log.ServerID)
	_, err := db.Query(sql)
	if err != nil {
		tool.Log.Errorf("Set processed failed!, Error: %v , Sql: %s", err, sql)
		return err
	}
	return nil
}

func checkNoemalQueryField(v map[string][]byte) (*orm.PreprocessLog, error) {
	queryData := &orm.PreprocessLog{}
	var value int
	var err error
	queryData.Bet, err = strconv.Atoi(string(v["Bet"]))
	if err != nil {
		return nil, err
	}
	queryData.FeatureBet, err = strconv.Atoi(string(v["FeatureBet"]))
	if err != nil {
		return nil, err
	}
	queryData.FeatureType, err = strconv.Atoi(string(v["FeatureType"]))
	if err != nil {
		return nil, err
	}
	value, err = strconv.Atoi(string(v["Count(Round)"]))
	if err != nil {
		return nil, err
	}
	queryData.TotalRound = int64(value)

	value, err = strconv.Atoi(string(v["SUM(Bet)"]))
	if err != nil {
		return nil, err
	}
	queryData.TotalBet = int64(value)

	value, err = strconv.Atoi(string(v["SUM(Bet_Win)"]))
	if err != nil {
		return nil, err
	}
	queryData.TotalWin = int64(value)

	queryData.ProcessStatus, err = strconv.Atoi(string(v["Process_Status"]))
	if err != nil {
		return nil, err
	}
	return queryData, nil
}

func checkFeatureQueryField(v map[string][]byte) (*orm.PreprocessLog, error) {
	queryData := &orm.PreprocessLog{}
	var value int
	var err error
	queryData.Bet, err = strconv.Atoi(string(v["Bet"]))
	if err != nil {
		return nil, err
	}
	queryData.TotalBet = int64(queryData.Bet)
	queryData.FeatureBet, err = strconv.Atoi(string(v["FeatureBet"]))
	if err != nil {
		return nil, err
	}
	queryData.FeatureType, err = strconv.Atoi(string(v["FeatureType"]))
	if err != nil {
		return nil, err
	}
	value, err = strconv.Atoi(string(v["Round"]))
	if err != nil {
		return nil, err
	}
	queryData.TotalRound = int64(value)

	value, err = strconv.Atoi(string(v["Bet_Win"]))
	if err != nil {
		return nil, err
	}
	queryData.TotalWin = int64(value)

	queryData.ProcessStatus, err = strconv.Atoi(string(v["Process_Status"]))
	if err != nil {
		return nil, err
	}
	return queryData, nil
}
