package preprocess

import (
	"MDBWeb/orm"
	"MDBWeb/tool"
	"encoding/json"
	"fmt"
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

func ProcessCQ9Log() error {
	fmt.Println("Process CQ9 Log-start")
	db := orm.MysqlDB()
	betClusterList := make([]orm.BetCluster, 0)
	err := db.Where("IsProcess = ?", 0).And("RoundID <> ?", "").Find(&betClusterList)
	//err := db.Where("IsProcess = ?", 0).And("ClusterID = ?", 112).Find(&betClusterList)
	if err != nil {
		panic("get betcluster failed!")
	}

	for _, v := range betClusterList {
		createPreprocessLog(&v)
		setProcessed(&v)
	}
	fmt.Println("Process CQ9 Log-End")
	return nil
}

//var count int

func createPreprocessLog(data *orm.BetCluster) {
	db := orm.MysqlDB()
	sql := "SELECT Bet,FeatureBet,FeatureType,FishType,Result,SUM(Bet),SUM(FeatureBet),SUM(Bet_Win),SUM(Round),Count(id),Process_Status FROM `gamelog_fish` WHERE ClusterID=" +
		strconv.Itoa(int(data.ClusterID)) + " AND(Process_Status=5 or Process_Status=12 or Process_Status=13 or Process_Status=14) GROUP BY Bet,FeatureBet,FeatureType,FishType,Process_Status"

	results, err := db.Query(sql)
	if err != nil {
		panic("query fish log failed")
	}
	fretureLogs := map[int]map[int]orm.PreprocessLog{} //map[bet]map[fishID]log

	for _, v := range results {
		bet, _ := strconv.Atoi(string(v["Bet"]))
		featureBet, _ := strconv.Atoi(string(v["FeatureBet"]))
		featureType, _ := strconv.Atoi(string(v["FeatureType"]))
		totalFeatureBet, _ := strconv.Atoi(string(v["SUM(FeatureBet)"]))
		totalRound, _ := strconv.Atoi(string(v["SUM(Round)"]))
		totalBet, _ := strconv.Atoi(string(v["SUM(Bet)"]))
		totalWin, _ := strconv.Atoi(string(v["SUM(Bet_Win)"]))
		processStatus, _ := strconv.Atoi(string(v["Process_Status"]))
		var disConnTimes, disConnSettle int
		if processStatus == 12 {
			//斷線結清
			disConnTimes, _ = strconv.Atoi(string(v["Count(id)"]))
			disConnSettle = totalWin
		}
		processLog := orm.PreprocessLog{
			ClusterID:       data.ClusterID,
			RoundID:         data.RoundID,
			Bet:             bet,
			FeatureBet:      featureBet,
			FeatureType:     featureType,
			FishType:        string(v["FishType"]),
			Result:          string(v["Result"]),
			TotalFeatureBet: int64(totalFeatureBet),
			TotalRound:      int64(totalRound),
			TotalBet:        int64(totalBet),
			TotalWin:        int64(totalWin),
			ProcessStatus:   processStatus,
			DisConTimes:     int64(disConnTimes),
			DisConSettle:    int64(disConnSettle),
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
				panic("Insert into preprocessLog failed!")
			}
		} else {
			processFeatureLog(&processLog, fretureLogs)
		}
	}
	insertFeatureLog(fretureLogs)
}

//統計results的魚得分資料
func processFeatureLog(log *orm.PreprocessLog, featureLogs map[int]map[int]orm.PreprocessLog) {
	result := SaveGameLog_Fish_Feature_Hit{}
	decodeResult := tool.DoZlibUnCompressGetString(log.Result)
	err := json.Unmarshal([]byte(decodeResult), &result)
	if err != nil {
		panic("Json Umarshal failed!")
	}
	var fLogBywinodds map[int]orm.PreprocessLog
	var ok bool
	if fLogBywinodds, ok = featureLogs[log.Bet]; !ok {
		featureLogs[log.Bet] = make(map[int]orm.PreprocessLog)
		fLogBywinodds = make(map[int]orm.PreprocessLog)
	}
	for _, v := range result.Kill_info {
		if v.Status == 1 {
			preLog := orm.PreprocessLog{}
			if preLog, ok = fLogBywinodds[v.Fish_Type]; !ok {
				log.FishID = strconv.Itoa(v.Fish_Type)
				fLogBywinodds[v.Fish_Type] = *log
			} else {
				preLog.TotalFeatureBet = preLog.TotalFeatureBet + log.TotalFeatureBet
				preLog.TotalWin = preLog.TotalWin + v.Win
				fLogBywinodds[v.Fish_Type] = preLog
			}
		}
	}
	featureLogs[log.Bet] = fLogBywinodds
}

func insertFeatureLog(featureLogs map[int]map[int]orm.PreprocessLog) {
	db := orm.MysqlDB()
	for _, v1 := range featureLogs {
		for _, log := range v1 {
			_, err := db.Insert(log)
			if err != nil {
				panic("Insert log to preprocessLog failed!")
			}
		}
	}
}

func setProcessed(log *orm.BetCluster) {
	db := orm.MysqlDB()
	sql := "UPDATE bet_cluster SET IsProcess=1 WHERE ClusterID=" + strconv.Itoa(int(log.ClusterID))
	_, err := db.Query(sql)
	if err != nil {
		panic("Set processed failed!")
	}
}
