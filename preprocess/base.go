package preprocess

import (
	"MDBWeb/orm"
	"fmt"
	"time"
)

func PreProcessLog() {
	interval := orm.Conf.Setting.ProcessInterval.Time
	timer := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-timer.C:
			go func() {
				runCQ9Log()
				fmt.Println("===========================================================================================================")
			}()
		}
	}

}

func runCQ9Log() {
	err := ProcessCQ9Log()
	if err != nil {
		panic("ProcessCQ9 Log Failed! outside Level")
	}
}
