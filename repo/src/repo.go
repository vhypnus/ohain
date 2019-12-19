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

	projectName string

}


func NewRepo(url string,srcCommit string,targetCommit string) Repo{
	var repo = Repo{url,"hello"}
	return repo
}

// return path
// repoCode 由group + service组成
func (r Repo) Sync() bool{
	path = ROOT + os.PathSeparator +r.projectName
	log.println
	exists,err := PathExists(path) 
	if !exists {
		os.Mkdir(path)
		exec.Comman("git clone ",path)
		exec.Comman()
	}
	//exec sych.sh
	var cmd = exec.Command(repoUrl,)
	cmd.Run()
	cmd.stdout
}


func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


// output string
func (r Repo) Diff(srcCommit string,targetCommit string) []string{

}


// 
func (r Repo) ReadFile(repoCode string,fileId string) string {
	// get filepath by fileId

	var content 

}


// 缓存文件id与文件路径的关系






