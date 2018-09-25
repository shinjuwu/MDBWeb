package baseinfo

// 平台結構
type PlatformInfo struct {
	PlatformID   int    `json:"platformId"`   //第三方平台編號
	PlatformName string `json:"platformName"` //第三方平台名稱

	PlatformAccount  string `json:"platformAccount"`  //第三方平台登入GameServer帳號
	PlatformPassword string `json:"platformPassword"` //第三方平台登入GameServer密碼
	IP               string `json:"ip"`               //第三方平台登入IP Address  (平台的IP 連線白名單使用,在名單內才允許連線)
	PlatformToken    string `json:"platformToken"`    //平台token(跟平台之間的驗證)
	TokenUpdateTime  string `json:"tokenUpdateTime"`  //TOKEN更新時間
}

//前端發送的封包格式
type PacketCmd_BetClusterGet struct {
	PlatformID       int   `json:"platform_id"`         // 平台編號 1:阿波羅 2:dt 3:dios
	MemberCode       int   `json:"member_code"`         // 業者的流水編號
	AgentID          int   `json:"agid"`                // 第三方代理ID
	ThirdPartyUserID int64 `json:"third_party_user_id"` //第三方平台登入用的 userID

	StartTime  string `json:"starttime"`   // start time to search
	EndTime    string `json:"endtime"`     // end time to search
	StartIndex int    `json:"start_index"` // start index to search
	RowNum     int    `json:"row_num"`     // row number to search
}

type PacketCmd_BetDetailGet struct {
	PlatformID int `json:"platform_id"` // 平台編號 1:阿波羅 2:dt 3:dios
	ServerID   int `json:"server_id"`   // Server 用的 ID
	MemberCode int `json:"member_code"` // 業者的流水編號
	AgentID    int `json:"agid"`        // 第三方代理ID
	//Token      string `json:"token"`       // token(跟玩家之間的驗證)
	ThirdPartyUserID int64 `json:"third_party_user_id"` //第三方平台登入用的 userID
	ClusterID        int64 `json:"cluster_id"`          // ID of bet cluster
}

type GameInfo struct {
	PlatformID          int    // 第三方平台編號
	GameID              int    // 遊戲編號
	GameName            string // 遊戲中文名稱
	GameEnName          string // 遊戲英文名稱
	GameMode            int8   // 遊戲模式 1:魚機 2:SLOT 3:撲克 4:麻將
	TableDestoryMode    int8   // 0: unknow 1:散桌後刪除此桌資訊  2: 散桌後保留此桌資訊,等待玩家重新入桌
	OpenTableMax        int    // 該遊戲能開啟的最大桌數
	TablePlayerMax      int    // 桌內人數上限
	DisconnectCleanData int8   // 斷線時候是否要清除資料 0:不要清除資料, 會啟動斷線連回機制 1:一斷線就清算資料, 把座位上的錢返回memberinfo
	AfterKickBefore     int8   // 後踢前 0:沒有後踢前 1:起動後踢前機制

	BetClusterSecs    int16 // 寶可夢使用 interval secs to collect bets into one bet cluster
	PlayTimeMax       int64 // 寶可夢使用 遊戲遊玩總時間
	InPlayTime        int64 // 寶可夢使用 走地發生時間
	SettlementTimeMax int64 // 寶可夢使用 結算結束時間

	Enable int8 // 這個廳館是否啟用

	// 記憶體的資料 ------------------------
	OpenTableNow int  // 該遊戲能開啟的目前桌數
	IsCloes      bool // 遊戲是否要關閉
}
