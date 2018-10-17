package model

import (
	"MDBWeb/orm"
	"MDBWeb/tool"
)

// 遊戲記錄 後台結構 (記錄細節)
type BackstageGameLog_Fish struct {
	ID         int `json:"id"`         // 遊戲紀錄id(細單編號)
	ServerID   int `json:"serverId"`   // Server 用的 ID
	PlatformID int `json:"platformId"` // 第三方平台編號
	MemberCode int `json:"memberCode"` // 第三方業者編號
	AgentID    int `json:"agId"`       // 第三方代理ID

	LobbyID  int    `json:"lobbyId"`  // 大廳編號
	GameID   int    `json:"gameId"`   // 遊戲編號
	TableID  string `json:"tableId"`  // 桌號
	Seat_ID  int    `json:"seatId"`   // 座位
	GameMode int8   `json:"gameMode"` // 遊戲模式 1:魚機 2:SLOT 3:撲克 4:麻將

	CreateTime string `json:"createTime"` // 建立時間

	User_ID  int64  `json:"userId"`   // 玩家帳號編號
	Account  string `json:"account"`  // 玩家帳號
	NickName string `json:"nickName"` // 帳號暱稱

	Round int64 `json:"round"` // 在遊戲內的第幾局

	Before_Balance_ci  float64 `json:"before_balance_ci"`  // 之前的銭 (除1萬倍的結果)
	Before_Balance_win float64 `json:"before_balance_win"` // 之前的銭 (除1萬倍的結果)
	Balance_ci         float64 `json:"balance_ci"`         // 玩家分數_投 (除1萬倍的結果)
	Balance_win        float64 `json:"balance_win"`        // 玩家贏的錢   win 先扣,在扣 ci  隨遊戲不斷變動 (除1萬倍的結果)

	Bet     float64 `json:"bet"`     // 單一押注 (除1萬倍的結果)
	WinOdds int     `json:"winOdds"` // 贏分賠率
	Bet_Win float64 `json:"betWin"`  // 玩家贏分 (除1萬倍的結果)

	Process_Status int    // 玩家處理的狀態 0:unknow 1:shoot 2:hit 3:feature_shoot 4:feature_hit
	FishType       string `json:"fishType"` // 魚死的fishtype
	Result         string `json:"-"`        // 玩家spin 的結果
	//Result string `json:"-"`    // 玩家spin 的結果
	Memo string `json:"memo"` // 中文備忘

	ClusterID        int64 `json:"clusterId"`        // ID of bet cluster (0 for none)
	ThirdPartyUserID int64 `json:"thirdPartyUserId"` //第三方平台登入用的 userid
}

// 魚機的細單資料 ResponseInfo_BetDetailFishGet is a response to get list of bet detail of fish game
type ResponseInfo_BetDetailFishGet struct {
	BetDetails []BackstageGameLog_Fish `json:"gamelogList"` // list of bet detail
}

func getFishBetDetail(betCluster *orm.BetCluster) *ResponseInfo_BetDetailFishGet {
	db := orm.MysqlDB()
	gamelogFish := make([]orm.GamelogFish, 0)
	err := db.Where("PlatformID = ? and ServerID = ? and ThirdPartyUserID = ? and ClusterID = ?",
		betCluster.PlatformID,
		betCluster.ServerID,
		betCluster.ThirdPartyUserID,
		betCluster.ClusterID).OrderBy("CreateTime").Find(&gamelogFish)
	if err != nil {
		panic(err)
	}

	res := getResGamelogFish(gamelogFish)
	return res
}

func getResGamelogFish(list []orm.GamelogFish) *ResponseInfo_BetDetailFishGet {
	res := &ResponseInfo_BetDetailFishGet{}
	for _, v := range list {
		betDetail := BackstageGameLog_Fish{
			ID:                 int(v.ID),
			ServerID:           v.ServerID,
			PlatformID:         v.PlatformID,
			MemberCode:         v.MemberCode,
			AgentID:            v.AgentID,
			LobbyID:            v.LobbyID,
			GameID:             v.GameID,
			TableID:            v.TableID,
			Seat_ID:            v.SeatID,
			GameMode:           int8(v.GameMode),
			CreateTime:         v.CreateTime.String(),
			User_ID:            v.UserID,
			Account:            v.Account,
			NickName:           v.NickName,
			Round:              v.Round,
			Before_Balance_ci:  tool.ConvertPrecision(v.BeforeBalanceCI),
			Before_Balance_win: tool.ConvertPrecision(v.BeforeBalanceWin),
			Balance_ci:         tool.ConvertPrecision(v.BalanceCI),
			Balance_win:        tool.ConvertPrecision(v.BalanceWin),
			Bet:                tool.ConvertPrecision(v.Bet),
			Bet_Win:            tool.ConvertPrecision(v.BetWin),
			WinOdds:            v.WinOdds,
			Process_Status:     v.ProcessStatus,
			FishType:           v.FishType,
			Result:             v.Result,
			Memo:               v.Memo,
			ClusterID:          v.ClusterID,
			ThirdPartyUserID:   v.ThirdPartyUserID,
		}
		res.BetDetails = append(res.BetDetails, betDetail)
	}
	return res
}
