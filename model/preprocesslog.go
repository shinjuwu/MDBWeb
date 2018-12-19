package model

import (
	"MDBWeb/orm"
	"MDBWeb/tool"
	"fmt"
	"strconv"
)

//CQ9細單
type ResInfoBetDetailFishGetForCQ9 struct {
	BetDetail FishDetailLogCQ9 `json:"gamelogList"`
}

type FishDetailLogCQ9 struct {
	RoundID   string        `json:"tid"`       //單號
	Paccount  string        `json:"paccount"`  //上層代理
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

func GetFishBetDetailForCQ9(betCluster *orm.BetCluster, paccount string) *ResInfoBetDetailFishGetForCQ9 {
	preprocessLog := GetProcessLog(betCluster.ClusterID, betCluster.ServerID)
	if preprocessLog == nil {
		return nil
	}
	fishGamelogList := []FishGameLog{}

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
		fishGamelogList = append(fishGamelogList, fishGameLog)
	}
	logList := assignFratureTypeBetWin(betCluster.ClusterID, betCluster.ServerID, fishGamelogList)
	fishDetailLogCQ9 := FishDetailLogCQ9{
		RoundID:   betCluster.RoundID,
		Paccount:  paccount,
		StartTime: betCluster.StartTime.String(),
		EndTime:   betCluster.EndTime.String(),
		Agent:     betCluster.Agent,
		LobbyID:   betCluster.LobbyID,
		Account:   betCluster.Account,
		Currency:  betCluster.Currency,
		Round:     getTotalRound(betCluster.ClusterID, betCluster.ServerID),
		OrderBet:  betCluster.Bet,
		Win:       getTotalWin(betCluster.ClusterID, betCluster.ServerID),
		WinLose:   betCluster.WinLose,
		GameLog:   logList,
	}
	res := &ResInfoBetDetailFishGetForCQ9{
		BetDetail: fishDetailLogCQ9,
	}
	return res
}

func GetProcessLog(clusterID int64, serverID int) []orm.PreprocessLog {
	db := orm.MysqlDB()
	preprocessLog := make([]orm.PreprocessLog, 0)
	sql := "SELECT Bet,FeatureBet,FeatureType,FishID,SUM(TotalFeatureHit),SUM(TotalRound),SUM(TotalBet),SUM(TotalWin),dis_con_times,dis_con_settle FROM `preprocess_log` WHERE ClusterID=" +
		strconv.Itoa(int(clusterID)) + " AND ServerID=" + strconv.Itoa(serverID) + " GROUP BY Bet,FeatureBet,FeatureType,FishID"
	fmt.Println(sql)
	results, err := db.Query(sql)
	if err != nil {
		tool.Log.Errorf("GetProcessLog error, sql= %s", sql)
		return nil
	}
	for k, v := range results {
		if k == 4 {
			fmt.Println(k)
		}
		bet, _ := strconv.Atoi(string(v["Bet"]))
		featureBet, _ := strconv.Atoi(string(v["FeatureBet"]))
		featureType, _ := strconv.Atoi(string(v["FeatureType"]))
		fishID := string(v["FishID"])
		totalFeatureHit, _ := strconv.Atoi(string(v["SUM(TotalFeatureHit)"]))
		totalRound, _ := strconv.Atoi(string(v["SUM(TotalRound)"]))
		totalBet, _ := strconv.Atoi(string(v["SUM(TotalBet)"]))
		totalWin, _ := strconv.Atoi(string(v["SUM(TotalWin)"]))
		disConnTimes, _ := strconv.Atoi(string(v["dis_con_times"]))
		disConnSettle, _ := strconv.Atoi(string(v["dis_con_settle"]))
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

func assignFratureTypeBetWin(clusterID int64, serverID int, logList []FishGameLog) []FishGameLog {
	for k, v := range logList {
		if v.FeatureType == 0 && v.FishID == "22" {
			totalWin := getSumOfFeatureWin(clusterID, 3, 14, serverID, v.Bet)
			v.TotalWin = totalWin
			logList[k] = v
		}
		if v.FeatureType == 0 && v.FishID == "23" {
			totalWin := getSumOfFeatureWin(clusterID, 4, 14, serverID, v.Bet)
			v.TotalWin = totalWin
			logList[k] = v
		}
	}
	return logList
}

func getSumOfFeatureWin(clusterID int64, featureType int, ps int, serverID int, featureBet int) int64 {
	db := orm.MysqlDB()
	ss := new(orm.PreprocessLog)
	totals, err := db.Where("ClusterID=?", clusterID).And("ServerID=?", serverID).And("FeatureBet=?", featureBet).And("FeatureType=?", featureType).And("Process_Status=?", ps).SumsInt(ss, "TotalWin")
	if err != nil {
		tool.Log.Errorf("getSumOfFeatureWin failed ! ClusterID= %d", clusterID)
		return 0
	}
	return totals[0]
}

func getTotalWin(clusterID int64, serverID int) int64 {
	db := orm.MysqlDB()
	ss := new(orm.PreprocessLog)
	totals, err := db.Where("ClusterID=?", clusterID).And("ServerID=?", serverID).SumsInt(ss, "TotalWin")
	if err != nil {
		tool.Log.Errorf("getTotalWin failed ! ClusterID= %d", clusterID)
		return 0
	}
	return totals[0]
}

func getTotalRound(clusterID int64, serverID int) int64 {
	db := orm.MysqlDB()
	ss := new(orm.PreprocessLog)
	totals, err := db.Where("ClusterID=?", clusterID).And("ServerID=?", serverID).And("FeatureType=?", 0).SumsInt(ss, "TotalRound")
	if err != nil {
		tool.Log.Errorf("getTotalRound failed ! ClusterID= %d", clusterID)
		return 0
	}
	return totals[0]
}
