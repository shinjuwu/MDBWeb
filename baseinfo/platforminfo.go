package baseinfo

import (
	"MDBWeb/orm"

	"github.com/go-xorm/xorm"
)

type PInfo baseinfotable

var PITable PInfo

func (t *PInfo) OnInit(db *xorm.Engine) int {
	PITable.Name = "platforminfo"
	all := make([]orm.Platforminfo, 0)
	err := db.Find(&all)
	if err != nil {
		panic(err)
	}
	for _, v := range all {
		PITable.Data = append(PITable.Data, v)
	}
	return len(PITable.Data)
}

func GetPlatformInfo(platformID int) (isGet bool, res *PlatformInfo) {
	pData := orm.Platforminfo{}
	isGet = false
	for _, v := range PITable.Data {
		data := v.(orm.Platforminfo)
		if platformID == data.PlatformID {
			pData = data
			isGet = true
			break
		}
	}
	res = &PlatformInfo{
		PlatformID:       pData.PlatformID,
		PlatformName:     pData.PlatformName,
		PlatformAccount:  pData.PlatformAccount,
		PlatformPassword: pData.PlatformPassword,
		IP:               pData.IP,
		PlatformToken:    pData.PlatformToken,
		TokenUpdateTime:  pData.TokenUpdateTime.String(),
	}

	return isGet, res
}
