package main

import (
	"MDBWeb/orm"
	"MDBWeb/preprocess"
	"MDBWeb/route"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	orm.OpenDB()
	orm.TableInit()
	go preprocess.PreProcessLog()
	engine := route.RegisterRouter()
	engine.Run(":8888")

}

func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hellow,world")
}
