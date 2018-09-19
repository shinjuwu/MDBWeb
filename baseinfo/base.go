package baseinfo

import (
	"MDBWeb/orm"
)

type baseinfotable struct {
	Name string
	Data []interface{}
}

func InitialAllTable() {
	db := orm.MysqlDB()
	gameinfo := new(GameInfo)
	gameinfo.OnInit(db)

	pinfo := new(PInfo)
	pinfo.OnInit(db)
}
