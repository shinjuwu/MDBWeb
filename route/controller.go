package route

import (
	"MDBWeb/baseinfo"
	"MDBWeb/model"
	"MDBWeb/sysconst"
	"MDBWeb/tool"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func apolloController(context *gin.Context) {
	DataMsg, Code, PacketCmd := contextAnalysis(context)
}

func dtController(context *gin.Context) {
	DataMsg, Code, PacketCmd := contextAnalysis(context)
}

func diosController(context *gin.Context) {
	DataMsg, Code, PacketCmd := contextAnalysis(context)
}

func contextAnalysis(context *gin.Context) (DataMsg interface{}, Code int, PacketCmd *CommonHttpPacketCmd) {
	PlatformAccount := context.Param("PlatformAccount")
	PlatformPassword := context.Param("PlatformPassword")
	data := context.Param("Send_Data")
	strPlatformID := context.Param("PlatformID")

	clientIP := context.Request.RemoteAddr
	ipStr := strings.Split(clientIP, ":")
	IP := ipStr[0]
	PlatformID, err := strconv.Atoi(strPlatformID)
	if err != nil {
		Code = int(sysconst.ERROR_CODE_ERROR_AUTH_PLATFORM)
		DataMsg = ""
		return
	}

	isGet, platform := baseinfo.GetPlatformInfo(PlatformID)
	if isGet == false {
		Code = int(sysconst.ERROR_CODE_ERROR_AUTH_PLATFORM)
		DataMsg = ""
		return
	}

	ipList := strings.Split(platform.IP, ",")
	if ipCheck, _ := tool.Contain(ipList, IP); !ipCheck {
		Code = int(sysconst.ERROR_CODE_ERROR_AUTH_PLATFORM_IP)
		DataMsg = ""
		return
	}

	if platform.PlatformAccount != PlatformAccount || platform.PlatformPassword != PlatformPassword {
		Code = int(sysconst.ERROR_CODE_ERROR_AUTH_PLATFORM_ACCOUNT)
		DataMsg = ""
		return
	}

	receiveMsgByte := []byte(data)
	err = json.Unmarshal(receiveMsgByte, PacketCmd)
	if err != nil {
		Code = int(sysconst.ERROR_CODE_ERROR_JSON_MARSHAL)
		DataMsg = ""
		return
	}

	DataMsg, Code = processCMD(context, PacketCmd, IP)
	return
}

func processCMD(context *gin.Context, PacketCmd *CommonHttpPacketCmd, ip string) (DataMsg interface{}, Code int) {
	switch PacketCmd.Cmd {
	case sysconst.HTTP_CMD_BET_CLUSTER_GET:
		cmdData := &baseinfo.PacketCmd_BetClusterGet{}
		err := json.Unmarshal([]byte(PacketCmd.Data), cmdData)
		if err != nil {
			panic(err)
		}
		return model.GetBetCluster(cmdData)

	case sysconst.HTTP_CMD_BET_DETAIL_GET:
		cmdData := &baseinfo.PacketCmd_BetDetailGet{}
		err := json.Unmarshal([]byte(PacketCmd.Data), cmdData)
		if err != nil {
			panic(err)
		}
		return model.GetBetDetail(cmdData)
	}
	return
}
