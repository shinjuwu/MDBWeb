package route

import (
	"MDBWeb/sysconst"

	"github.com/gin-gonic/gin"
)

type CommonHttpPacketCmd struct {
	Cmd           string `json:"cmd"`            // 命令種類
	Sys           string `json:"sys"`            // 是否是 system cmd  ( sys:"game" 遊戲封包 sys:"system" 系統封包)
	PlatformToken string `json:"platform_token"` // 平台token(跟平台之間的驗證)
	IsEncode      bool   `json:"isEncode"`       // 是否加密
	Data          string `json:"data"`           // 封包資料
}

func RegisterRouter() *gin.Engine {
	router := gin.Default()
	router.POST(sysconst.THIRD_PARTY_PLATFROM_APOLLO, apolloController)
	router.POST(sysconst.THIRD_PARTY_PLATFROM_DT, dtController)
	router.POST(sysconst.THIRD_PARTY_PLATFROM_DIOS, diosController)
	return router
}
