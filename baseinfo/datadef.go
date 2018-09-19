package baseinfo

// 平台結構
type PlatformInfo struct {
	PlatformID   int    `json:"platformId"`   //第三方平台編號
	PlatformName string `json:"platformName"` //第三方平台名稱

	PlatformAccount  string `json:"platformAccount"`  //第三方平台登入GameServer帳號
	PlatformPassword string `json:"platformPassword"` //第三方平台登入GameServer密碼
	IP               string `json:"ip"`               //第三方平台登入IP Address  (平台的IP 連線白名單使用,在名單內才允許連線)
	PlatformToken    string `json:"platformToken"`    //平台token(跟平台之間的驗證)
	TokenUpdateTime  string `json:"tokenUpdateTime"`  //TOKEN更新時間
}
