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
	ThirdPartyUserID int64     `xorm:"notnull index 'ThirdPartyUserID'"`
}
