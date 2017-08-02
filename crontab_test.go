package lib

import (
	"testing"
	"time"
	"sina.com.cn/cms/lib"
	"fmt"
)

func TestCron(t *testing.T)  {
	cron,err := lib.NewCrontab("Asia/Shanghai")
	lib.CheckErr(err)
	//cron.Repeat(cron.SecondDuration(3), func(duration *time.Duration) {
	//	err,logFile := new(lib.Logger).Init("/var/log/{APP}/{DATE}/{TIME}.log", "interaction")
	//	fmt.Println(time.Now(),logFile,err)
	//})

	cron.Repeat(cron.DurationAtSecond(3), func(duration *time.Duration) {
		err,logFile := new(lib.Logger).Init("/var/log/{APP}/{DATE}/{TIME}.log", "interaction")
		fmt.Println(time.Now(),logFile,err)
	})

}

