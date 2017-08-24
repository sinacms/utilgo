package utilgo

import (
	"testing"
	"time"
	"fmt"
)

func TestCron(t *testing.T)  {
	cron,err := NewCrontab("Asia/Shanghai")
	CheckErr(err)
	//cron.Repeat(cron.SecondDuration(3), func(duration *time.Duration) {
	//	err,logFile := new(lib.Logger).Init("/var/log/{APP}/{DATE}/{TIME}.log", "interaction")
	//	fmt.Println(time.Now(),logFile,err)
	//})

	cron.Repeat(cron.DurationAtSecond(3), func(duration *time.Duration) {
		err,logFile := new(Logger).Init("/var/log/{APP}/{DATE}/{TIME}.log", "interaction")
		fmt.Println(time.Now(),logFile,err)
	})

}

