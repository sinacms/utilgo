package utilgo

import (
	"os"
	"strings"
)

func FindInGoPath(path string, inPackage bool )string{
	dir := ""
	if true == inPackage {
		dir = "/src/"
	}
	goPath := os.Getenv("GOPATH")
	paths := strings.Split(goPath, ":")
	tmpPath := ""
	for p := range paths {
		tmpPath = strings.Join([]string{paths[p],dir,path}, "")
		if true == FileExist(tmpPath){
			return tmpPath
		}
	}
	return tmpPath
}

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}