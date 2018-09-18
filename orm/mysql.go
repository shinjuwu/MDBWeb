package orm

import (
	"fmt"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	//編譯但不使用
	_ "github.com/go-sql-driver/mysql"

	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

var db *xorm.Engine
var conf map[string]interface{}

func init() {
	config.Load(file.NewSource(
		file.WithPath("./config/config.json"),
	))
	conf = config.Map()
}

func OpenDB() {
	fmt.Println("mysqldb->Open DB")

	connStr := "root:pass@tcp(127.0.0.1:3306)/one1cloud_main?parseTime=true"
	var err error
	db, err = xorm.NewEngine("mysql", connStr)
	if err != nil {
		panic("Connect DB error")
	}

	db.SetMapper(core.GonicMapper{})
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	db.SetDefaultCacher(cacher)

	logSetting(db, "debug")
}

func logSetting(db *xorm.Engine, status string) {
	switch status {
	case "debug":
		db.ShowSQL(true)
		db.Logger().SetLevel(core.LOG_DEBUG)
	default:
		db.ShowSQL(false)
		db.Logger().SetLevel(core.LOG_ERR)
	}
}

func TableInit() {
	err := db.Sync2(
		new(BetCluster),
	)

	if err != nil {
		panic(err.Error())
	}
}
