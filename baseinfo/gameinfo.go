package baseinfo

import (
	"MDBWeb/orm"

	"github.com/go-xorm/xorm"
)

type GameInfo baseinfotable

var GITable GameInfo

func (t *GameInfo) OnInit(db *xorm.Engine) int {
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
