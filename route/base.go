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

// 共用的回應結構
type CommonHttpResponseInfo struct {
	Code    int         `json:"code"`    // 回應的代碼
	Message string      `json:"message"` // 回應的訊息
	Ret     string      `json:"ret"`     // 回應的命令種類
	SN      int         `json:"sn"`      // 回應的SN
	Data    interface{} `json:"data"`    // 回應的資料
}

func RegisterRouter() *gin.Engine {
	router := gin.Default()
	router.POST(sysconst.THIRD_PARTY_PLATFROM_APOLLO, apolloController)
	router.POST(sysconst.THIRD_PARTY_PLATFROM_DT, dtController)
	router.POST(sysconst.THIRD_PARTY_PLATFROM_DIOS, diosController)
	router.GET(sysconst.THIRD_PARTY_PLATFROM_CQ9, cq9Controller)
	return router
}
