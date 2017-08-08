package utilgo

import (
	"testing"
	"sina.com.cn/comos/lib"
	"time"
	"fmt"
	"sync"
)

func TestLogger_Init(t *testing.T)  {
	//newFile(t)


	logFile := new(lib.Logger).Init("/var/log/{APP}/{DATE}.log")
	if _log := time.Now().Format("/var/log/lib.test/2006-01-02.log"); logFile != _log {
		t.Fail()
		t.Log("log not match", logFile, _log)
	}
}



func newFile(t *testing.T){
	year, month, day := time.Now().Date()
	now := time.Now()
	loc, err := time.LoadLocation("Asia/Shanghai")
	if nil != err {
		t.Log(err)
		t.Fail()
	}
	_time := time.Date(year, month, day, now.Hour(), now.Minute(), now.Second()+5, 0, loc)
	ticker := time.NewTicker(_time.Sub(time.Now()))
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		select {
		case <-ticker.C:
			fmt.Println("mkfile ", time.Now())
			newFile(t)
			wg.Done()

		}
	}()
	wg.Wait()
}