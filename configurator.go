package utilgo

import (
	"log"
	"github.com/BurntSushi/toml"
	"os"
)

type Configurator struct {
	dir  string
	mode MODE_TYPE
	debug bool
}

type Attribute interface{}


func (this *Configurator)SetLoadDir(dir string){
	this.dir = dir
}

func (this *Configurator)GetLoadDir()string{
	return this.dir
}

func (this *Configurator)SetMode(mode MODE_TYPE){
	this.mode = mode
}
func (this *Configurator)GetMode(mode MODE_TYPE)MODE_TYPE{
	return this.mode
}
func (this *Configurator)Parse(attributes  Attribute, subPath string)(Attribute){
	_modeDir := ""
	if len(this.ModeDir(this.mode)) > 0 {
		_modeDir =  string(os.PathSeparator) + this.ModeDir(this.mode) + string(os.PathSeparator)
	}else{
		_modeDir = ""
	}
	configPath :=  FindInGoPath(this.dir + _modeDir  + subPath , true)
	log.Println("Main config is loaded from "+configPath)
	_, e := toml.DecodeFile(configPath, attributes)
	if nil != e {
		log.Fatalln("Main config is not exists")
	}
	return attributes
}
func (this *Configurator)ModeDir(mode MODE_TYPE) string{
	var dir string
	switch mode {
	case MODE_DEV:
		dir = MODE_DEV_NAME
	case MODE_TEST:
		dir = MODE_TEST_NAME
	case MODE_SIM:
		dir = MODE_SIM_NAME
	case MODE_PROD:
		dir = MODE_PROD_NAME
	}
	return dir
}
func (this *Configurator)isDebug()(bool){
	return bool(this.debug)
}
func (this *Configurator)setDebug(debug bool){
	this.debug = debug
}