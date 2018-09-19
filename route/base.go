package route

import (
	"MDBWeb/sysconst"

	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine {
	router := gin.Default()
	router.POST(sysconst.THIRD_PARTY_PLATFROM_APOLLO, apolloController)
	router.POST(sysconst.THIRD_PARTY_PLATFROM_DT, dtController)
	router.POST(sysconst.THIRD_PARTY_PLATFROM_DIOS, diosController)
	return router
}
