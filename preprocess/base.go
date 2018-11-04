package preprocess

import (
	"MDBWeb/orm"
	"time"
)

func PreProcessLog() {
	runCQ9Log() //首次開機先處理一次
	interval := orm.Conf.Setting.ProcessInterval.Time
	timer := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-timer.C:
			go func() {
				runCQ9Log()
			}()
		}
	}
}

func runCQ9Log() {

	ProcessCQ9Log()

}
