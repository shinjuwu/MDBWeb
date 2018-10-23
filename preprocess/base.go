package preprocess

import (
	"MDBWeb/orm"
	"fmt"
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
				currentTime := time.Now()
				fmt.Println(currentTime)
				runCQ9Log()
				currentTime = time.Now()
				fmt.Println(currentTime)
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
