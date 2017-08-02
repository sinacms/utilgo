package lib

import (
	"strings"
	"os"
	"time"
	"log"
	"path"
	"runtime"
)

type Logger struct  {
	handler *os.File
}
func (this *Logger) Init(pathOfLog string, appName string)(error, string){
	pathOfLog = strings.Replace(pathOfLog, "{APP}", appName, -1)
	pathOfLog = strings.Replace(pathOfLog, "{DATE}", time.Now().Format("2006-01-02"), -1)
	pathOfLog = strings.Replace(pathOfLog, "{TIME}", time.Now().Format("03:04:05"), -1)
	err := os.MkdirAll(path.Dir(pathOfLog), 0766)
	if err != nil {
		//log.Fatalf("error opening log dir: %v", this.err)
		return err, ""
	}
	this.handler, err = os.OpenFile(pathOfLog, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		//log.Fatalf("error opening log file: %v", this.err)
		return err, ""
	}
	log.SetOutput(this.handler)
	return nil, pathOfLog
}

func CheckErr(err error){
	if nil != err {
		panic(err)
	}
}

func IgnoreErr(err error){
	if nil != err {
		log.Println(err)
	}
}
func IgnoreErrWithMsg(err error, msg string){
	if nil != err {
		log.Println("%s:%v", msg, err)
	}
}

func CheckErrWithMsg(err error, msg string){
	if nil != err {
		log.Panicf("%s:%v", msg, err)
	}
}
func CheckOk(ok bool, msg string){
	if !ok {
		pc, fn, line := lineInfo(2)
		log.Printf("[fail] in %s[%s:%d]\n", pc, fn, line)
	}
}

func VarLog(v ...interface{}){
	for _,_v := range v {
		log.Printf("%#v \n", _v)
	}
	log.Println()
}

func lineInfo(skip int)(string, string, int){
	pc, fn, line, _ := runtime.Caller(skip)
	return runtime.FuncForPC(pc).Name(), fn, line
}

func TraceTest(v ...interface{}){
	if SingleEnv().UnderMode(MODE_TEST) {
		VarLog(v...)
	}
}

