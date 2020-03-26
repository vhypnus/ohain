package diff

import (
	"strings"
	"fmt"
)


type diff struct {
	methodName string

	codeNo []int

	noteNo []int

	logNo []int

	testCase int
}


func parse(content string) string{
	var lines = strings.Split(content,"\n")

	var diffs = make([]Diff,0,10)
	for _ ,line := range lines {
		fmt.Println(line)
		if len(strings.TrimSpace(line)) > 4 && strings.TrimSpace(line)[0:4] == "diff" {
			item := Diff{FilePath:strings.Split(line," ")[3]}
			diffs = append(diffs,item)
		} 
	}


		
	return ""
}


func Diff(projectName string,srcCommit string,targetCommit) string {
	osName := runtime.GOOS
	var scriptpath string
	if osName == "windows" {
		scriptpath = SCRIPT_ROOT+"diff.bat"
	} else if osName == "linux" {
		scriptpath = SCRIPT_ROOT+"diff.bat"
	}

	var repopath = r.path 

	var cmd *exec.Cmd = exec.Command(scriptpath,repopath,srcCommit,targetCommit)
	content := parserCmd(cmd) 
	return content
}

