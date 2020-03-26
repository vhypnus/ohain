package sync

import (
	"fmt"
	"runtime"
	"bytes"
	"os/exec"
)



func (r *Project) Sync(projectName string){

	// todo if not exist ,git clone
	osName := runtime.GOOS
	var scriptpath string
	if osName == "windows" {
		scriptpath = SCRIPT_ROOT+"sync.bat"
	} else if osName == "linux" {
		scriptpath = SCRIPT_ROOT+"sync.sh"
	}

	log.Printf("osName %s scriptpath %s",osName,scriptpath)
	
	// todo 路径计算
	var path = ""
	var cmd *exec.Cmd = exec.Command(scriptpath,projectName,"") 
	parserCmd(cmd)
}

func parserCmd(cmd *exec.Cmd) string{

	var out bytes.Buffer
    cmd.Stdout = &out

	err := cmd.Run()

	if err != nil{
		panic(err)
	} 

	output := out.String()

	log.Printf("cmd content %s",output)
	// todo error case deal

	//success return content
	return output
}
