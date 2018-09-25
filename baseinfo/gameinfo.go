package baseinfo

import (
	"MDBWeb/orm"

	"github.com/go-xorm/xorm"
)

type GInfo baseinfotable

var GITable GInfo

func (t *GInfo) OnInit(db *xorm.Engine) int {
	GITable.Name = "gameinfo"
	all := make([]orm.Gameinfo, 0)
	err := db.Find(&all)
	if err != nil {
		panic(err)
	}
	for _, v := range all {
		GITable.Data = append(GITable.Data, v)
	}
	return len(GITable.Data)
}

func GetGameNameEN(platformID int, gameID int) string {
	for _, v := range GITable.Data {
		data := v.(orm.Gameinfo)
		if platformID == data.PlatformID && gameID == data.GameID {
			return data.GameEnName
		}
	}
	return ""
}

func GetGameMode(platformID int, gameID int) int8 {
	for _, v := range GITable.Data {
		data := v.(orm.Gameinfo)
		if platformID == data.PlatformID && gameID == data.GameID {
			return int8(data.GameMode)
		}
	}
	return -1
}
