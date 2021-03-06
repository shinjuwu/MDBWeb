package model

import (
	"MDBWeb/baseinfo"
	"MDBWeb/orm"
	"MDBWeb/sysconst"
	"MDBWeb/tool"
	"encoding/json"
)

// BetCluster is a info of the bet cluster(注單)
type BetCluster struct {
	ClusterID  int64 `json:"cluster_id"` // 注單ID (ID of bet cluster)
	ServerID   int   `json:"server_id"`  // Server 用的 ID
	PlatformID int   `json:"-"`          // 平台ID (ID of platform)

	MemberCode int `json:"member_code"` //第三方業者編號
	AgentID    int `json:"agid"`        //第三方代理ID
	LobbyID    int `json:"lobby_id"`    //大廳規則編號

	GameID   int    `json:"game_id"`   // 遊戲ID (ID of played game)
	GameName string `json:"game_name"` // 遊戲名稱 (name of the game)
	GameMode int8   `json:"game_mode"` // 遊戲模式 (mode of the game)

	User_ID          int64   `json:"user_id"`             // 玩家帳號編號 (neo1 會員編號 )
	Bet              float64 `json:"bet"`                 // 押注額 accumulated bets of bet details of this cluster
	Win              float64 `json:"win"`                 // 贏分   accumulated wins of bet details of this cluster
	WinLose          float64 `json:"winlose"`             // 輸贏   a derived field calculated by -`Bet` + `Win`
	StartTime        string  `json:"starttime"`           // 注單起始時間 start datetime of this cluster
	EndTime          string  `json:"endtime"`             // 注單結束時間 end datetime of this cluster
	ThirdPartyUserID int64   `json:"third_party_user_id"` // 第三方平台登入用的 userid (例如這邊是存阿波羅會員的userid)
}

// ResponseInfo_BetClusterGet is a response to get list of bet cluster
type ResponseInfo_BetClusterGet struct {
	DataCount   int          `json:"data_count"`   // 資料總數筆數
	BetClusters []BetCluster `json:"bet_clusters"` // list of bet cluster
}

func GetBetDetail(cmdData *baseinfo.PacketCmd_BetDetailGet) (DataMsg interface{}, Code int) {
	DataMsg = "unknow"
	Code = int(sysconst.ERROR_CODE_SUCCESS)
	db := orm.MysqlDB()
	betCluster := &orm.BetCluster{
		PlatformID:       cmdData.PlatformID,
		ServerID:         cmdData.ServerID,
		ThirdPartyUserID: cmdData.ThirdPartyUserID,
		ClusterID:        cmdData.ClusterID,
	}
	_, err := db.Get(betCluster)
	if err != nil {
		panic(err)
	}
	gameMode := baseinfo.GetGameMode(betCluster.PlatformID, betCluster.GameID)
	res := getBetDetailLog(gameMode, betCluster)
	bytes, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	DataMsg = string(bytes)
	return
}

func getBetDetailLog(gameMode int8, betCluster *orm.BetCluster) interface{} {
	switch gameMode {
	case int8(sysconst.GAME_MODE_FISH):
		return getFishBetDetail(betCluster)
	case int8(sysconst.GAME_MODE_SLOT):
		return getSlotBetDetail(betCluster)
	}
	return ""
}

//使用roundID去取得該筆BetCluster
func GetBetCluster(roundID string) *orm.BetCluster {
	db := orm.MysqlDB()
	betCluster := new(orm.BetCluster)
	has, err := db.Where("RoundID=?", roundID).Get(betCluster)
	if err != nil {
		tool.Log.Errorf("DB select failed! err=%+v", err)
		return nil
	}
	if !has {
		return nil
	}
	return betCluster
}
