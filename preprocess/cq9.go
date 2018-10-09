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
	fmt.Println("Process CQ9 Log")
	db := orm.MysqlDB()
	betClusterList := make([]orm.BetCluster, 0)
	err := db.Where("IsProcess = ?", 0).Find(&betClusterList)
	if err != nil {
		panic("get betcluster failed!")
	}

	for _, v := range betClusterList {
		createPreprocessLog(&v)
	}
	return nil
}

func createPreprocessLog(data *orm.BetCluster) {
	db := orm.MysqlDB()
	sql := "SELECT WinOdds,FeatureType,FishType,Result,SUM(Bet),SUM(FeatureBet),SUM(Bet_Win),SUM(Round) FROM `gamelog_fish` GROUP BY WinOdds,FeatureType,FishType"
	results, err := db.Query(sql)
	if err != nil {
		panic("query fish log failed")
	}
	fretureLogs := map[int]map[int]orm.PreprocessLog{} //map[winadd]map[fishID]log
	for _, v := range results {
		winodds, _ := strconv.Atoi(string(v["WinOdds"]))
		featureType, _ := strconv.Atoi(string(v["FeatureType"]))
		totalFeatureBet, _ := strconv.Atoi(string(v["SUM(FeatureBet)"]))
		totalRound, _ := strconv.Atoi(string(v["SUM(Round)"]))
		totalBet, _ := strconv.Atoi(string(v["SUM(Bet)"]))
		totalWin, _ := strconv.Atoi(string(v["SUM(Bet_Win)"]))
		processLog := orm.PreprocessLog{
			ClusterID:       data.ClusterID,
			RoundID:         data.RoundID,
			WinOdds:         winodds,
			FeatureType:     featureType,
			FishType:        string(v["FishType"]),
			Result:          string(v["Result"]),
			TotalFeatureBet: int64(totalFeatureBet),
			TotalRound:      int64(totalRound),
			TotalBet:        int64(totalBet),
			TotalWin:        int64(totalWin),
		}
		fishs := strings.Split(",", processLog.FishType)
		if len(fishs) == 1 {
			processLog.FishID = processLog.FishType //TOD 轉int
			_, err := db.Insert(processLog)
			if err != nil {
				panic("Insert into preprocessLog failed!")
			}
		} else {
			processFeatureLog(&processLog, fretureLogs)
		}
	}
	insertFeatureLog(fretureLogs)
	fmt.Println(results)
}

//統計results的魚得分資料
func processFeatureLog(log *orm.PreprocessLog, featureLogs map[int]map[int]orm.PreprocessLog) {
	result := SaveGameLog_Fish_Feature_Hit{}
	decodeResult := tool.DoZlibUnCompressGetString(log.Result)
	err := json.Unmarshal([]byte(decodeResult), &result)
	if err != nil {
		panic("Json Umarshal failed!")
	}
	fLogBywinodds := map[int]orm.PreprocessLog{}
	var ok bool
	if fLogBywinodds, ok = featureLogs[log.WinOdds]; !ok {
		featureLogs[log.WinOdds] = map[int]orm.PreprocessLog{}
	}
	for _, v := range result.Kill_info {
		if v.Status == 1 {
			preLog := orm.PreprocessLog{}
			if preLog, ok = fLogBywinodds[v.Fish_id]; !ok {
				fLogBywinodds[v.Fish_id] = orm.PreprocessLog{}
				log.FishID = strconv.Itoa(v.Fish_id)
				fLogBywinodds[v.Fish_id] = *log
			} else {
				preLog.TotalFeatureBet = preLog.TotalFeatureBet + log.TotalFeatureBet
				preLog.TotalWin = preLog.TotalWin + v.Win
				fLogBywinodds[v.Fish_id] = preLog
			}
		}
	}
	featureLogs[log.WinOdds] = fLogBywinodds
	fmt.Println(featureLogs)
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
