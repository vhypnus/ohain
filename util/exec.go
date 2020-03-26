package util


import (
	"os/exec"
)


func ExecCmd(scriptpath string,args ... string) {

	var cmd *exec.Cmd = exec.Command(scriptpath,args) 

	var out bytes.Buffer
    cmd.Stdout = &out

	err := cmd.Run()

	if err != nil{
		panic(err)
	} 

	output := out.String()
	// todo error

	//success return content
	return output

}