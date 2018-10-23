package main

import (
	"MDBWeb/orm"
	"MDBWeb/preprocess"
	"MDBWeb/route"
	"strconv"
)

func main() {
	port := orm.Conf.Setting.Serve.Port
	orm.OpenDB()
	orm.TableInit()
	go preprocess.PreProcessLog()
	engine := route.RegisterRouter()
	engine.Run(":" + strconv.Itoa(port))

}
