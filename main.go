package main

import (
	"MDBWeb/orm"
	"MDBWeb/preprocess"
	"MDBWeb/route"
)

func main() {
	orm.OpenDB()
	orm.TableInit()
	go preprocess.PreProcessLog()
	engine := route.RegisterRouter()
	engine.Run(":8888")

}
