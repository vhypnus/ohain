/*
 * fileId  example java :package +className
 *
 *
 */
package core

import (
	"os"
	"io/ioutil"
	"log"
	"gopkg.in/src-d/go-git.v4"
)

const (
	REPO_ROOT_DIR = ""
)

type Repo struct {

	// unique
	url string

	path string

	projectName string

}


func NewRepo(url string,srcCommit string,targetCommit string) Repo{
	var repo = Repo{url,"hello"}
	return repo
}

// return path
// repoCode 由group + service组成
func (r Repo) Sync() bool{
	// get project path by r.url
	projectPath = ROOT + os.PathSeparator +r.projectName
	
	// 
	scriptPath = 
	content := exec(scriptpath,projectPath) 
	
	return content
}


// output string
func (r Repo) Diff(srcCommit string,targetCommit string) []string{

}


// 
func (r Repo) ReadFile(filepath string) string {
	path = ROOT +os.PathSeparator+filepath

	os.Open

	var content 

}


// 缓存文件id与文件路径的关系

func (r Repo) exec(scriptpath string,args...string) string{
	cmd := exec.Command(scriptpath,args)

	var out bytes.Buffer
    cmd.Stdout = &out

	err := cmd.Run()

	if err != nil{
		panic(err)
	} 

	output := out.String()
	return output
}




