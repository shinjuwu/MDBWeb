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
	LobbyID   int           `json:"lobbyID"`   //遊戲館ID
	Account   string        `json:"account"`   //玩家帳號
	Currency  string        `json:"currency"`  //幣別
	Round     int64         `json:"round"`     //總局數
	OrderBet  int64         `json:"orderbet"`  //總壓分
	Win       int64         `json:"win"`       //總贏分
	WinLose   int64         `json:"winLose"`   //總輸贏
	GameLog   []FishGameLog `json:"gameLog"`   //遊戲紀錄
}

type FishGameLog struct {
	FeatureType     int    `json:"featureType"` //道具種類
	Bet             int    `json:"bet"`         //押注額
	FeatureBet      int    `json:"featureBet"`  //特殊道具壓住額
	FishID          string `json:"fishID"`
	TotalBet        int64  `json:"totalBet"`         //總押注
	TotalFeatureHit int64  `json:"totalFeatureHit"`  //總道具擊中次數
	TotalWin        int64  `json:"totalWin"`         //總贏分
	TotalRound      int64  `json:"totalRound"`       //總局數
	DisConTimes     int64  `json:"disconnectTimes"`  //斷線次數
	DisConSettle    int64  `json:"disconnectSettle"` //斷線結清
}

func GetFishBetDetailForCQ9(betCluster *orm.BetCluster) *ResInfoBetDetailFishGetForCQ9 {
	preprocessLog := GetProcessLog(betCluster.ClusterID)
	fishGamelogList := []FishGameLog{}
	var effectTotalRound int64
	var galtingTotalWin int64
	var drillBomeTotalWin int64
	for _, v := range preprocessLog {
		fishGameLog := FishGameLog{
			FeatureType:     v.FeatureType,
			Bet:             v.Bet,
			FeatureBet:      v.FeatureBet,
			FishID:          v.FishID,
			TotalBet:        v.TotalBet,
			TotalFeatureHit: v.TotalFeatureHit,
			TotalWin:        v.TotalWin,
			TotalRound:      v.TotalRound,
			DisConTimes:     v.DisConTimes,
			DisConSettle:    v.DisConSettle,
		}
		if v.FeatureType == 3 { //蓋特機槍
			galtingTotalWin = galtingTotalWin + v.TotalWin
		} else if v.FeatureType == 4 {
			drillBomeTotalWin = drillBomeTotalWin + v.TotalWin
		}
		effectTotalRound = effectTotalRound + v.TotalRound
		fishGamelogList = append(fishGamelogList, fishGameLog)
	}

	for k, v := range fishGamelogList {
		if v.FeatureType == 3 { //蓋特機槍
			v.TotalRound = galtingTotalWin
		} else if v.FeatureType == 4 { //鑽頭砲
			v.TotalRound = drillBomeTotalWin
		}
		fishGamelogList[k] = v
	}
	fishDetailLogCQ9 := FishDetailLogCQ9{
		RoundID:   betCluster.RoundID,
		StartTime: betCluster.StartTime.String(),
		EndTime:   betCluster.EndTime.String(),
		Agent:     betCluster.Agent,
		LobbyID:   betCluster.LobbyID,
		Account:   betCluster.Account,
		Currency:  betCluster.Currency,
		Round:     effectTotalRound, //改成有效回合數
		OrderBet:  betCluster.Bet,
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
	sql := "SELECT Bet,FeatureBet,FeatureType,FishID,SUM(TotalFeatureHit),SUM(TotalRound),SUM(TotalBet),SUM(TotalWin),dis_con_times,dis_con_settle FROM `preprocess_log` WHERE ClusterID=" +
		strconv.Itoa(int(clusterID)) + " GROUP BY Bet,FeatureBet,FeatureType,FishID"
	results, err := db.Query(sql)
	if err != nil {
		panic("GetProcessLog error")
	}
	for _, v := range results {
		bet, _ := strconv.Atoi(string(v["Bet"]))
		featureBet, _ := strconv.Atoi(string(v["FeatureBet"]))
		featureType, _ := strconv.Atoi(string(v["FeatureType"]))
		fishID := string(v["FishID"])
		totalFeatureHit, _ := strconv.Atoi(string(v["SUM(TotalFeatureHit)"]))
		totalRound, _ := strconv.Atoi(string(v["SUM(TotalRound)"]))
		totalBet, _ := strconv.Atoi(string(v["SUM(TotalBet)"]))
		totalWin, _ := strconv.Atoi(string(v["SUM(TotalWin)"]))
		disConnTimes, _ := strconv.Atoi(string(v["SUM(dis_con_times)"]))
		disConnSettle, _ := strconv.Atoi(string(v["SUM(dis_con_settle)"]))
		processLog := orm.PreprocessLog{
			ClusterID:       clusterID,
			Bet:             bet,
			FeatureBet:      featureBet,
			FeatureType:     featureType,
			FishID:          fishID,
			TotalFeatureHit: int64(totalFeatureHit),
			TotalRound:      int64(totalRound),
			TotalBet:        int64(totalBet),
			TotalWin:        int64(totalWin),
			DisConTimes:     int64(disConnTimes),
			DisConSettle:    int64(disConnSettle),
		}
		preprocessLog = append(preprocessLog, processLog)
	}
	return preprocessLog
}
