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
	}

	for _, v := range betClusterList {
		createPreprocessLog(&v)
		setProcessed(&v)
		tool.Log.Infof("PreLog process end, RoundID = %s", v.RoundID)
	}
	tool.Log.Info("CQ9 Prepross routine end.")
}

//var count int

func createPreprocessLog(data *orm.BetCluster) {
	db := orm.MysqlDB()
	sql := "SELECT Bet,FeatureBet,FeatureType,FishType,Result,SUM(Bet),SUM(Bet_Win),Count(Round),Count(id),Process_Status FROM `gamelog_fish` WHERE ClusterID=" +
		strconv.Itoa(int(data.ClusterID)) + " AND(Process_Status=5 or Process_Status=12 or Process_Status=13 or Process_Status=14) GROUP BY Bet,FeatureBet,FeatureType,FishType,Process_Status"

	results, err := db.Query(sql)
	if err != nil {
		tool.Log.Errorf("Query fish log failed, Sql: %s , Error = %v", sql, err)
		return
	}
	//fretureLogs := map[int]map[int]orm.PreprocessLog{} //map[bet]map[fishID]log

	for _, v := range results {
		queryData, err := checkQueryFlied(v)
		if err != nil {
			tool.Log.Errorf("String to int failed, Error = %v", err)
			return
		}
		var disConnTimes int
		var disConnSettle int64
		if queryData.ProcessStatus == 12 {
			//斷線結清
			disConnTimes, err = strconv.Atoi(string(v["Count(id)"]))
			if err != nil {
				tool.Log.Errorf("String to int failed, Error = %v", err)
				return
			}
			disConnSettle = queryData.TotalWin
		}
		processLog := orm.PreprocessLog{
			ClusterID:     data.ClusterID,
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
			DisConTimes:   int64(disConnTimes),
			DisConSettle:  disConnSettle,
		}
		if processLog.FishType == "" {
			continue
		}
		processLog.FishType = strings.TrimLeft(processLog.FishType, "[")
		processLog.FishType = strings.TrimRight(processLog.FishType, "]")
		fishs := strings.Split(processLog.FishType, ",")

		if len(fishs) == 1 {
			processLog.FishID = processLog.FishType
			_, err := db.Insert(processLog)
			if err != nil {
				tool.Log.Errorf("Insert into preprocessLog failed! Error = %v  , ProcessLog = %v", err, processLog)
				return
			}
		} else {
			fretureLogs := processFeatureLog(&processLog)
			if fretureLogs != nil {
				insertFeatureLog(fretureLogs)
			}
		}
	}
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

//回傳總壓鑄額給下個update用,違反單一直則原則,但是為了節省迴圈數
func insertFeatureLog(featureLogs map[int]map[int]orm.PreprocessLog) int64 {
	var totalWin int64
	db := orm.MysqlDB()
	for _, v1 := range featureLogs {
		for _, log := range v1 {
			_, err := db.Insert(log)
			if err != nil {
				tool.Log.Errorf("Insert log to preprocessLog failed! Error: %v , log : %v", err, log)
				continue
			}
			totalWin = totalWin + log.TotalWin
		}
	}
	return totalWin
}

func setProcessed(log *orm.BetCluster) {
	db := orm.MysqlDB()
	sql := "UPDATE bet_cluster SET IsProcess=1 WHERE ClusterID=" + strconv.Itoa(int(log.ClusterID))
	_, err := db.Query(sql)
	if err != nil {
		tool.Log.Errorf("Set processed failed!, Error: %v , Sql: %s", err, sql)
		return
	}
}

func checkQueryFlied(v map[string][]byte) (*orm.PreprocessLog, error) {
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
