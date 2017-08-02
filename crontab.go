package lib

import (
	"time"
	"sync"
	"log"
)

type Crontab struct {
	local *time.Location
	task func()
}


func NewCrontab(location string)(*Crontab,error){
	loc, err := time.LoadLocation(location)//"Asia/Shanghai")
	if nil != err {
		return nil, err
	}
	return &Crontab{
		local: loc,
		task: nil,
	}, nil
}
func (t *Crontab)Location()*time.Location{
	return t.local
}
func (t *Crontab)SecondDuration(seconds int) (time.Duration) {
	now := time.Now()
	_time := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()+seconds, 0, t.local)
	return  _time.Sub(now)
}
func (t *Crontab)DurationAtHour(hours int) (time.Duration) {
	now := time.Now()
	_time := time.Date(now.Year(), now.Month(), now.Day(), hours, 0, 0, 0, t.local)
	return  _time.Sub(now)
}
func (t *Crontab)DurationAtSecond(sec int) (time.Duration) {
	now := time.Now()
	_time := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), sec, 0, t.local)
	return  _time.Sub(now)
}
func (t *Crontab)Repeat(duration time.Duration, task func(duration *time.Duration)){
	tick := time.NewTicker(duration)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		select {
		case <-tick.C:
			log.Println("tick ")
			task(&duration)
			t.Repeat(duration, task)
			wg.Done()
		}
	}()
	wg.Wait()
}

//func (t *Crontab)At(year, month, day, hour, minute, second int ,task func())error{
//	slice := []int{year, month, day, hour, minute, second}
//	var repeatUnit int
//	var hasValue bool
//	now := time.Now()
//	m,_ := strconv.Atoi(now.Month().String())
//	nowSlice := []int{now.Year(), m, now.Day(), now.Hour(), now.Minute(), now.Second()}
//	for i, num := range slice {
//		if !hasValue {
//			if num == -1 {
//				slice[i] = nowSlice[i]
//				repeatUnit = i
//				continue
//			}else{
//				slice[i] = num
//			}
//		}else{
//			if num == -1 {
//				return errors.New("数值右边不能再出现-1")
//			}else{
//				slice[i] = num
//			}
//		}
//		hasValue = true
//	}
//	m,ok := slice[1].(time.Month)
//	if !ok {
//		return errors.New("month有误")
//	}
//	desc := time.Date(slice[0], m, slice[2], slice[3], slice[4], slice[5], 0, t.local)
//	dur := desc.Sub(time.Now())
//	if dur > 0 {
//		t.Repeat(dur, task)
//	}else{
//		return errors.New("计划时间必须大于当前时间")
//	}
//}
