package orm

import "time"

type BetCluster struct {
	ClusterID        int64     `xorm:"pk index autoincr notnull 'ClusterID'"`
	ServerID         int       `xorm:"pk notnull 'ServerID' default(0)"`
	PlatformID       int       `xorm:"pk index 'PlatformID' notnull"`
	MemberCode       int       `xorm:"notnull 'MemberCode' default(0)"`
	AgentID          int       `xorm:"notnull 'AgentID' default(0)"`
	LobbyID          byte      `xorm:"notnull 'LobbyID' default(0)"`
	GameID           int       `xorm:"pk 'GameID' notnull"`
	UserID           int64     `xorm:"notnull index 'UserID'"`
	Bet              int64     `xorm:"notnull 'Bet'"`
	Win              int64     `xorm:"notnull 'Win'"`
	WinLose          int64     `xorm:"notnull 'WinLose'"`
	StartTime        time.Time `xorm:"created 'StartTime'"`
	EndTime          time.Time `xorm:"'EndTime'"`
	OrderState       byte      `xorm:"notnull default(0) 'OrderState'"`
	ThirdPartyUserID int64     `xorm:"notnull default(0) index 'ThirdPartyUserID'"`
	Account          string    `xorm:"notnull varchar(45) 'Account'"`
	Agent            string    `xorm:"notnull varchar(45) 'Agent'"`
	Currency         string    `xorm:"notnull varchar(45) default('RD') 'Currency'"`
	RoundID          string    `xorm:"notnull varchar(45) 'RoundID'"`
	Round            int64     `xorm:"notnull default(0) 'Round'"`
	IsProcess        byte      `xorm:"notnull default(0) 'IsProcess'"`
}

type Gameinfo struct {
	ID                  int64  `xorm:"pk autoincr notnull"`
	PlatformID          int    `xorm:"index notnull 'PlatformID'"`
	GameID              int    `xorm:"index notnull 'GameID'"`
	GameName            string `xorm:"notnull varchar(45) 'GameName'"`
	GameEnName          string `xorm:"notnull varchar(4) 'GameEnName'"`
	GameMode            byte   `xorm:"notnull 'GameMode'"`
	TableDestoryMode    int    `xorm:"notnull 'TableDestoryMode'"`
	OpenTableMax        int    `xorm:"notnull default(0) 'OpenTableMax'"`
	TablePlayerMax      int    `xorm:"notnull 'TablePlayerMax'"`
	DisconnectCleanData byte   `xorm:"notnull default(0) 'DisconnectCleanData'"`
	AfterKickBefore     byte   `xorm:"default(1) 'AfterKickBefore'"`
	BetClusterSecs      int    `xorm:"notnull 'BetClusterSecs'"`
	PlayTimeMax         int    `xorm:"notnull default(0) 'PlayTimeMax'"`
	InPlayTime          int    `xorm:"notnull default(0) 'InPlayTime'"`
	SettlementTimeMax   int    `xorm:"notnull default(0) 'SettlementTimeMax'"`
	Enable              byte   `xorm:"notnull default(0) 'Enable'"`
}

type GamelogError struct {
	ID         int64     `xorm:"pk index notnull autoincr"`
	PlatformID int       `xorm:"notnull 'PlatformID'"`
	MemberCode int       `xorm:"'MemberCode'"`
	AgentID    int       `xorm:"'AgentID'"`
	LobbyID    int       `xorm:"notnull 'LobbyID'"`
	GameID     int       `xorm:"notnull 'GameID'"`
	TableID    string    `xorm:"notnull varchar(45) 'TableID'"`
	SeatID     string    `xorm:"notnull varchar(45) 'Seat_ID'"`
	GameMode   byte      `xorm:"notnull 'GameMode'"`
	CreateTime time.Time `xorm:"created 'CreateTime'"`
	UserID     int64     `xorm:"notnull 'User_ID'"`
	Account    string    `xorm:"varchar(45) 'Account'"`
	NickName   string    `xorm:"varchar(45) 'NickName'"`
	BalanceCI  float64   `xorm:"notnull 'Balance_ci'"`
	BalanceWin float64   `xorm:"notnull 'Balance_win'"`
	ErrorLevel byte      `xorm:"notnull 'ErrorLevel'"`
	Result     string    `xorm:"notnull varchar(1000) 'Result'"`
	Memo       string    `xorm:"varchar(45) 'Memo'"`
}

type GamelogFish struct {
	ID               int64     `xorm:"pk notnull autoincr"`
	ServerID         int       `xorm:"pk notnull default(0) 'ServerID'"`
	ClusterID        int64     `xorm:"notnull 'ClusterID'"`
	PlatformID       int       `xorm:"notnull 'PlatformID'"`
	MemberCode       int       `xorm:"'MemberCode'"`
	AgentID          int       `xorm:"'AgentID'"`
	LobbyID          int       `xorm:"notnull 'LobbyID'"`
	GameID           int       `xorm:"notnull 'GameID'"`
	TableID          string    `xorm:"varchar(45) notnull 'TableID'"`
	SeatID           int       `xorm:"notnull 'Seat_ID'"`
	GameMode         byte      `xorm:"notnull 'GameMode'"`
	CreateTime       time.Time `xorm:"created 'CreateTime'"`
	UserID           int64     `xorm:"notnull default(0) 'User_ID'"`
	ThirdPartyUserID int64     `xorm:"notnull default(0) 'ThirdPartyUserID'"`
	Account          string    `xorm:"varchar(45) notnull 'Account'"`
	NickName         string    `xorm:"varchar(45) notnull 'NickName'"`
	Round            int64     `xorm:"notnull default(0) 'Round'"`
	BeforeBalanceCI  int64     `xorm:"notnull default(0) 'Before_Balance_ci'"`
	BeforeBalanceWin int64     `xorm:"notnull default(0) 'Before_Balance_win'"`
	BalanceCI        int64     `xorm:"notnull default(0) 'Balance_ci'"`
	BalanceWin       int64     `xorm:"notnull default(0) 'Balance_win'"`
	Bet              int64     `xorm:"notnull default(0) 'Bet'"`
	WinOdds          int       `xorm:"notnull default(0) 'WinOdds'"`
	BetWin           int64     `xorm:"notnull default(0) 'Bet_Win'"`
	ProcessStatus    int       `xorm:"notnull 'Process_Status'"`
	FishType         string    `xorm:"varchar(200) notnull default('0') 'FishType'"`
	Result           string    `xorm:"varchar(4000) notnull 'Result'"`
	Memo             string    `xorm:"varchar(100)"`
	FeatureBet       int64     `xorm:"notnull default(0) 'FeatureBet'"`
	Currency         string    `xorm:"notnull varchar(45) default('RD') 'Currency'"`
	FeatureType      int       `xorm:"notnull default(0) 'FeatureType'"`
}

type GamelogSlot struct {
	ID               int64     `xorm:"pk notnull autoincr"`
	ServerID         int       `xorm:"pk notnull default(0) 'ServerID'"`
	PlatformID       int       `xorm:"notnull 'PlatformID'"`
	MemberCode       int       `xorm:"'MemberCode'"`
	AgentID          int       `xorm:"'AgentID'"`
	LobbyID          int       `xorm:"notnull 'LobbyID'"`
	GameID           int       `xorm:"notnull 'GameID'"`
	TableID          string    `xorm:"varchar(45) notnull 'TableID'"`
	SeatID           int       `xorm:"notnull 'Seat_ID'"`
	GameMode         byte      `xorm:"notnull 'GameMode'"`
	CreateTime       time.Time `xorm:"created 'CreateTime'"`
	UserID           int64     `xorm:"notnull 'User_ID'"`
	Account          string    `xorm:"varchar(45) notnull 'Account'"`
	NickName         string    `xorm:"varchar(45) notnull 'NickName'"`
	Round            int64     `xorm:"notnull 'Round'"`
	BeforeBalanceCI  int64     `xorm:"notnull 'Before_Balance_ci'"`
	BeforeBalanceWin int64     `xorm:"notnull 'Before_Balance_win'"`
	BalanceCI        int64     `xorm:"notnull 'Balance_ci'"`
	BalanceWin       int64     `xorm:"notnull 'Balance_win'"`
	Bet              int64     `xorm:"notnull 'Bet'"`
	BetWin           int64     `xorm:"notnull 'Bet_Win'"`
	ProcessStatus    int       `xorm:"notnull 'Process_Status'"`
	Result           string    `xorm:"varchar(100) notnull"`
	Memo             string    `xorm:"varchar(45) 'Memo'"`
	ClusterID        int64     `xorm:"notnull 'ClusterID'"`
	ThirdPartyUserID int64     `xorm:"notnull 'ThirdPartyUserID'"`
	BetID            string    `xorm:"varchar(60) notnull 'BetID'"`
	Currency         string    `xorm:"notnull varchar(45) default('RD') 'Currency'"`
}

type Platforminfo struct {
	PlatformID       int       `xorm:"index unique pk notnull 'PlatformID'"`
	PlatformName     string    `xorm:"varchar(45) notnull 'PlatformName'"`
	PlatformAccount  string    `xorm:"varchar(45) notnull 'PlatformAccount'"`
	PlatformPassword string    `xorm:"varchar(45) notnull 'PlatformPassword'"`
	WebAPIMode       int       `xorm:"notnull default(0) 'WebApiMode'"`
	IP               string    `xorm:"varchar(100) 'IP'"`
	PlatformToken    string    `xorm:"varchar(45) 'PlatformToken'"`
	TokenUpdateTime  time.Time `xorm:"'TokenUpdateTime'"`
}

type PreprocessLog struct {
	ID              int64  `xorm:"pk notnull autoincr"`
	ClusterID       int64  `xorm:"notnull 'ClusterID'"`
	RoundID         string `xorm:"notnull varchar(45) 'RoundID'"`
	WinOdds         int    `xorm:"notnull default(0) 'WinOdds'"`
	FeatureType     int    `xorm:"notnull default(0)"`
	FishType        string `xorm:"varchar(200) notnull default('0') 'FishType'"`
	Result          string `xorm:"varchar(100) notnull 'Result'"`
	TotalFeatureBet int64  `xorm:"notnull default(0) 'TotalFeatureBet'"`
	TotalRound      int64  `xorm:"notnull default(0) 'TotalRound'"`
	TotalBet        int64  `xorm:"notnull default(0) 'TotalBet'"`
	TotalWin        int64  `xorm:"notnull default(0) 'TotalWin'"`
	FishNO          string `xorm:"notnull varchar(11) default(0) 'FishNO'"`
}
