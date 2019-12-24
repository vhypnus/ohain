/*
 * fileId  example java :package +className
 *
 *
 */
package repo

import (
	"os"
	"os/exec"
	"io/ioutil"
	"log"
	"runtime"
	"bytes"
)
const RepoRootDir string = "C:\\Users\\monan\\Desktop\\all in one\\source code"

type Repo struct {

	// unique
	url string

	path string

	projectName string

}


type Diff struct {

	//abs file path
	filepath string

	// 位图 ，表示某一行删除或者新增

}


func NewRepo(url string) Repo{
	path := RepoRootDir + string(os.PathSeparator) + "driving-order"
	var repo = Repo{url,path,url}
	return repo
}

// return path
// repoCode 由group + service组成
func (r Repo) Sync(){

	osName := runtime.GOOS
	var scriptpath string
	if osName == "windows" {
		scriptpath = "C:\\Users\\monan\\Desktop\\all in one\\source code\\ohain\\repo\\script\\sync.bat"
	} else if osName == "linux" {
		scriptpath = "C:\\Users\\monan\\Desktop\\all in one\\source code\\ohain\\repo\\script\\sync.bat"
	}

	content := execute(scriptpath,r.path) 
	log.Println(content)

	//todo error deal
}


// output string
//srcCommit 不是属性，通过函数参数传递
func (r Repo) Diff(srcCommit string,targetCommit string) string{
	osName := runtime.GOOS
	var scriptpath string
	if osName == "windows" {
		scriptpath = "C:\\Users\\monan\\Desktop\\all in one\\source code\\ohain\\repo\\script\\diff.bat"
	} else if osName == "linux" {
		scriptpath = "C:\\Users\\monan\\Desktop\\all in one\\source code\\ohain\\repo\\script\\diff.bat"
	}

	content := execute(scriptpath,r.path) 
	return content
}


// 
func (r Repo) ReadFile(filepath string) []byte {
	path := RepoRootDir + string(os.PathSeparator) + filepath
	log.Println(path)

	data,err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return data

}


// 缓存文件id与文件路径的关系

func execute(scriptpath string,args string) string{
	cmd := exec.Command(scriptpath,args)

	var out bytes.Buffer
    cmd.Stdout = &out

	err := cmd.Run()

	if err != nil{
		panic(err)
	} 

	output := out.String()

	// todo error case deal

	//success return content
	return output
}






