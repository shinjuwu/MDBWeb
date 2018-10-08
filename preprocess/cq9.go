package preprocess

import (
	"MDBWeb/orm"
	"fmt"
	"strconv"
	"strings"
)

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
			processLog.FishNO = processLog.FishType
			_, err := db.Insert(processLog)
			if err != nil {
				panic("Insert into preprocessLog failed!")
			}
		} else {
			insertFeatureLog()
		}

	}
	fmt.Println(results)
}

func insertFeatureLog(log *orm.PreprocessLog) {

}
