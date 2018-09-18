package main

import (
	"MDBWeb/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	orm.OpenDB()
	orm.TableInit()
	engine := gin.Default()
	engine.Any("/", WebRoot)
	engine.Run(":8888")
}

func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hellow,world")
}
