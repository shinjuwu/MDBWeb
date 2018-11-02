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
var Conf settings

type database struct {
	DBName   string `json:"dbname"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
	UserName string `json:"username"`
	Passwd   string `json:"passwd"`
}
type pinterval struct {
	Time int `json:"time"`
}
type serve struct {
	Port int `json:"port"`
}
type data struct {
	Database        database  `json:"database"`
	ProcessInterval pinterval `json:"processInterval"`
	Serve           serve     `json:"HTTPserve"`
}

type settings struct {
	Setting data `json:"settings"`
}

func init() {
	config.Load(file.NewSource(
		file.WithPath("./config/config.json"),
	))
	config.Scan(&Conf)
}

func OpenDB() {
	fmt.Println("mysqldb->Open DB")
	// username := Conf.Setting.Database.UserName
	// passwd := Conf.Setting.Database.Passwd
	// ipaddr := Conf.Setting.Database.Address
	// port := Conf.Setting.Database.Port
	// dbName := Conf.Setting.Database.DBName
	connStr := "root:pass@tcp(127.0.0.1:3306)/one1cloud_main?parseTime=true"
	//connStr := username + ":" + passwd + "@tcp(" + ipaddr + ":" + strconv.Itoa(port) + ")/" + dbName + "?parseTime=true"
	var err error
	db, err = xorm.NewEngine("mysql", connStr)
	if err != nil {
		panic("Connect DB error")
	}

	db.SetMapper(core.GonicMapper{})
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	db.SetDefaultCacher(cacher)

	logSetting(db, "release")
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
		new(Gameinfo),
		new(GamelogError),
		new(GamelogFish),
		new(GamelogSlot),
		new(Platforminfo),
		new(PreprocessLog),
	)

	if err != nil {
		panic(err.Error())
	}
}

func MysqlDB() *xorm.Engine {
	return db
}
