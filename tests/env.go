package utilgo

import (
	"os/exec"
	"os"
	"strings"
)

var instance *env
type MODE_TYPE uint
const  (
	MODE_DEV MODE_TYPE = 1 + iota
	MODE_TEST
	MODE_SIM
	MODE_PROD
)
const (
	MODE_DEV_NAME = "dev"
	MODE_TEST_NAME = "test"
	MODE_SIM_NAME = "sim"
	MODE_PROD_NAME = "prod"
)
var CONF_MODES = map[MODE_TYPE]string{
	MODE_DEV: MODE_DEV_NAME,
	MODE_TEST: MODE_TEST_NAME,
	MODE_SIM: MODE_SIM_NAME,
	MODE_PROD: MODE_PROD_NAME,
}
type env struct {
	user       string
	Mode       MODE_TYPE
	parameters map[string]interface{}
}
func (e *env)User()string{
	return e.user
}
func (e *env)Getex(key string)(interface{},  bool){
	ret,ok := e.parameters[key]
	if ok {
		return ret, true
	}
	return nil,false
}
func (e *env)Get(key string)(interface{}){
	ret,_ := e.parameters[key]
	return ret
}
func (e *env)Set(key string, value interface{}){
	e.parameters[key] = value
}
func (e *env) init()  {
	user := os.Getenv("SUDO_USER")
	if len(user) == 0 {
		cmd := exec.Command("logname")
		_user, _ := cmd.CombinedOutput()
		e.user = strings.Trim(string(_user), "\n\r ")
	}
}



func NewEnv()*env {
	env := env{}
	env.init()
	env.parameters = make(map[string]interface{})
	return &env
}

func SingleEnv()*env{
	if nil == instance {
		instance = NewEnv()
	}
	return instance
}

func (e *env)OverMode(target MODE_TYPE)bool{
	if e.Mode >= target {
		return true
	}else{
		return false
	}
}
func (e *env)UnderMode(target MODE_TYPE)bool{
	if e.Mode <= target {
		return true
	}else{
		return false
	}
}