package model

import (
	"MDBWeb/orm"
	"strconv"
)

//CQ9細單
type ResInfoBetDetailFishGetForCQ9 struct {
	BetDetail FishDetailLogCQ9 `json:"gamelogList"`
}

type FishDetailLogCQ9 struct {
	RoundID   string        `json:"tid"`       //單號
	StartTime string        `json:"startTime"` //開始時間
	EndTime   string        `json:"endTime"`   //結束時間
	Agent     string        `json:"agent"`     //代理商
	Account   string        `json:"account"`   //玩家帳號
	Currency  string        `json:"currency"`  //幣別
	Round     int64         `json:"round"`     //總局數
	Bet       int64         `json:"bet"`       //總壓分
	Win       int64         `json:"win"`       //總贏分
	WinLose   int64         `json:"winLose"`   //總輸贏
	GameLog   []FishGameLog `json:"gameLog"`   //遊戲紀錄
}

type FishGameLog struct {
	FeatureType     int    `json:"featureType"` //道具種類
	Bet             int    `json:"bet"`         //押注額
	FishID          string `json:"fishID"`
	TotalBet        int64  `json:"totalBet"`         //總押注
	TotalFeatureBet int64  `json:"totalFeatureBet"`  //總道具押注
	TotalWin        int64  `json:"totalWin"`         //總贏分
	TotalRound      int64  `json:"totalRound"`       //總局數
	DisConTimes     int64  `json:"disconnectTimes"`  //斷線次數
	DisConSettle    int64  `json:"disconnectSettle"` //斷線結清
}

func GetFishBetDetailForCQ9(betCluster *orm.BetCluster) *ResInfoBetDetailFishGetForCQ9 {
	preprocessLog := GetProcessLog(betCluster.ClusterID)
	fishGamelogList := []FishGameLog{}
	for _, v := range preprocessLog {
		fishGameLog := FishGameLog{
			FeatureType:     v.FeatureType,
			Bet:             v.Bet,
			FishID:          v.FishID,
			TotalBet:        v.TotalBet,
			TotalFeatureBet: v.TotalFeatureBet,
			TotalWin:        v.TotalWin,
			TotalRound:      v.TotalRound,
			DisConTimes:     v.DisConTimes,
			DisConSettle:    v.DisConSettle,
		}
		fishGamelogList = append(fishGamelogList, fishGameLog)
	}
	fishDetailLogCQ9 := FishDetailLogCQ9{
		RoundID:   betCluster.RoundID,
		StartTime: betCluster.StartTime.String(),
		EndTime:   betCluster.EndTime.String(),
		Agent:     betCluster.Agent,
		Account:   betCluster.Account,
		Currency:  betCluster.Currency,
		Round:     betCluster.Round,
		Bet:       betCluster.Bet,
		Win:       betCluster.Win,
		WinLose:   betCluster.WinLose,
		GameLog:   fishGamelogList,
	}
	res := &ResInfoBetDetailFishGetForCQ9{
		BetDetail: fishDetailLogCQ9,
	}
	return res
}

func GetProcessLog(clusterID int64) []orm.PreprocessLog {
	db := orm.MysqlDB()
	preprocessLog := make([]orm.PreprocessLog, 0)
	sql := "SELECT WinOdds,FeatureType,FishID,SUM(TotalFeatureBet),SUM(TotalRound),SUM(TotalBet),SUM(TotalWin),dis_con_times,dis_con_settle FROM `preprocess_log` WHERE ClusterID=" +
		strconv.Itoa(int(clusterID)) + " GROUP BY WinOdds,FeatureType,FishID"
	results, err := db.Query(sql)
	if err != nil {
		panic("GetProcessLog error")
	}
	for _, v := range results {
		bet, _ := strconv.Atoi(string(v["Bet"]))
		featureType, _ := strconv.Atoi(string(v["FeatureType"]))
		fishID := string(v["FishID"])
		totalFeatureBet, _ := strconv.Atoi(string(v["SUM(TotalFeatureBet)"]))
		totalRound, _ := strconv.Atoi(string(v["SUM(TotalRound)"]))
		totalBet, _ := strconv.Atoi(string(v["SUM(TotalBet)"]))
		totalWin, _ := strconv.Atoi(string(v["SUM(TotalWin)"]))
		processLog := orm.PreprocessLog{
			ClusterID:       clusterID,
			Bet:             bet,
			FeatureType:     featureType,
			FishID:          fishID,
			TotalFeatureBet: int64(totalFeatureBet),
			TotalRound:      int64(totalRound),
			TotalBet:        int64(totalBet),
			TotalWin:        int64(totalWin),
		}
		preprocessLog = append(preprocessLog, processLog)
	}
	return preprocessLog
}
