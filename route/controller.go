package route

import (
	"MDBWeb/baseinfo"
	"MDBWeb/model"
	"MDBWeb/sysconst"
	"MDBWeb/tool"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var HttpService_SN int = 0 // 回應的SN 累加序號

func apolloController(context *gin.Context) {
	httpResponse := contextAnalysis(context)

	sendMessage(context, httpResponse)
}

func dtController(context *gin.Context) {
	httpResponse := contextAnalysis(context)

	sendMessage(context, httpResponse)
}

func diosController(context *gin.Context) {
	httpResponse := contextAnalysis(context)

	sendMessage(context, httpResponse)
}

func cq9Controller(context *gin.Context) {
	reponse := getOrderInfo(context)

	sendMessageCQ9(context, reponse)
}

func sendMessageCQ9(context *gin.Context, reponse interface{}) {
	//去斜線
	s, ok := reponse.(string)
	if ok {
		rawData := json.RawMessage(s)
		reponse = rawData
	}
	dataMsgByte, err := json.Marshal(reponse)
	if err != nil {
		panic("httpResponse json encode failed")
	}

	context.String(http.StatusOK, string(dataMsgByte))
}

func sendMessage(context *gin.Context, httpReponse *CommonHttpResponseInfo) {
	//去斜線
	s, ok := httpReponse.Data.(string)
	if ok {
		rawData := json.RawMessage(s)
		httpReponse.Data = rawData
	}
	dataMsgByte, err := json.Marshal(httpReponse)
	if err != nil {
		panic("httpResponse json encode failed")
	}

	context.String(http.StatusOK, string(dataMsgByte))
}

func getOrderInfo(context *gin.Context) interface{} {
	token := context.Query("token")
	detailInfo := getDetailOrderInfo(token)
	betCluster := model.GetBetCluster(detailInfo.Data.RoundID)
	if betCluster == nil {
		res := &CommonHttpResponseInfo{
			Code:    int(sysconst.ERROR_CODE_CQ9_ORDER_NOT_FOUND),
			Message: "找不到注單",
		}
		return res
	}
	ResFishBetInfo := model.GetFishBetDetailForCQ9(betCluster)
	if len(ResFishBetInfo.BetDetail.GameLog) == 0 {
		res := &CommonHttpResponseInfo{
			Code:    int(sysconst.ERROR_CODE_CQ9_ORDERDETAIL_NOT_FOUND),
			Message: "找不到細單",
		}
		return res
	}
	return ResFishBetInfo
}

func checkIP(context *gin.Context, platformID int) (msg string, code int) {
	clientIP := context.Request.RemoteAddr
	ipStr := strings.Split(clientIP, ":")
	IP := ipStr[0]
	isGet, platform := baseinfo.GetPlatformInfo(platformID)
	if isGet == false {
		code = int(sysconst.ERROR_CODE_ERROR_AUTH_PLATFORM)
		msg = ""
		return
	}

	ipList := strings.Split(platform.IP, ",")
	if ipCheck, _ := tool.Contain(ipList, IP); !ipCheck {
		code = int(sysconst.ERROR_CODE_ERROR_AUTH_PLATFORM_IP)
		msg = ""
		return
	}
	return
}

//Applo平台頁面解析
func contextAnalysis(context *gin.Context) *CommonHttpResponseInfo {
	var Code int
	var PacketCmd *CommonHttpPacketCmd
	var DataMsg interface{}
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
		res := packageHTTPReponse(DataMsg, Code, PacketCmd)
		return res
	}

	isGet, platform := baseinfo.GetPlatformInfo(PlatformID)
	if isGet == false {
		Code = int(sysconst.ERROR_CODE_ERROR_AUTH_PLATFORM)
		DataMsg = ""
		res := packageHTTPReponse(DataMsg, Code, PacketCmd)
		return res
	}

	ipList := strings.Split(platform.IP, ",")
	if ipCheck, _ := tool.Contain(ipList, IP); !ipCheck {
		Code = int(sysconst.ERROR_CODE_ERROR_AUTH_PLATFORM_IP)
		DataMsg = ""
		res := packageHTTPReponse(DataMsg, Code, PacketCmd)
		return res
	}

	if platform.PlatformAccount != PlatformAccount || platform.PlatformPassword != PlatformPassword {
		Code = int(sysconst.ERROR_CODE_ERROR_AUTH_PLATFORM_ACCOUNT)
		DataMsg = ""
		res := packageHTTPReponse(DataMsg, Code, PacketCmd)
		return res
	}

	receiveMsgByte := []byte(data)
	err = json.Unmarshal(receiveMsgByte, PacketCmd)
	if err != nil {
		Code = int(sysconst.ERROR_CODE_ERROR_JSON_MARSHAL)
		DataMsg = ""
		res := packageHTTPReponse(DataMsg, Code, PacketCmd)
		return res
	}

	DataMsg, Code = processCMD(context, PacketCmd, IP)
	res := packageHTTPReponse(DataMsg, Code, PacketCmd)
	return res
}

func packageHTTPReponse(DataMsg interface{}, Code int, PacketCmd *CommonHttpPacketCmd) *CommonHttpResponseInfo {
	var httpResponse *CommonHttpResponseInfo
	// 儲存cmd 和 累加sn
	HttpService_SN++
	httpResponse.Ret = PacketCmd.Cmd
	httpResponse.SN = HttpService_SN

	// 組合回傳結果
	httpResponse.Code = sysconst.ErrorCode[Code].Code
	httpResponse.Message = sysconst.ErrorCode[Code].Message
	httpResponse.Data = DataMsg
	return httpResponse
}

func processCMD(context *gin.Context, PacketCmd *CommonHttpPacketCmd, ip string) (DataMsg interface{}, Code int) {
	switch PacketCmd.Cmd {
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

type TokenDetailRes struct {
	Data   DetailToken  `json:"data"`
	Status DetailStatus `json:"status"`
}

type DetailToken struct {
	Token string `json:"detailtoken"`
}

type DetailStatus struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Date    string `json:"datetime"`
}

type ResDetailOrder struct {
	Data   OrderInfo    `json:"data"`
	Status DetailStatus `json:"status"`
}

type OrderInfo struct {
	RoundID  string `json:"roundid"`
	Account  string `json:"account"`
	ID       string `json:"id"`
	Gametype string `json:"gametype"`
	Paccount string `json:"paccount"`
}

//Cq9限定，未來有其他平台有類似機制再重購
func getDetailOrderInfo(token string) *ResDetailOrder {
	v := url.Values{}
	v.Set("token", token)
	formBody := ioutil.NopCloser(strings.NewReader(v.Encode()))
	req, err := http.NewRequest("POST", "http://api.cqgame.games/gamepool/cq9/game/detailtoken", formBody)
	if err != nil {
		panic("error")
	}
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJnYW1laGFsbCI6ImNxOSIsInRlYW0iOiJBUCIsImp0aSI6IjUyODM4NDU1MyIsImlhdCI6MTUzNTk2NDM1OSwiaXNzIjoiQ3lwcmVzcyIsInN1YiI6IkdTVG9rZW4ifQ.OtEO9IT3ZgmeM0Kp_fjYE-MaAtGQyGFPLwvDBwbPQCI")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("error")
	}
	res := ResDetailOrder{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic("error")
	}
	return &res
}
