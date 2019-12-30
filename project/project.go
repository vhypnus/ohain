/*
 * fileId  example java :package +className
 *
 *
 */
package project

import (
	"os"
	"os/exec"
	"io/ioutil"
	"log"
	"runtime"
	"bytes"
)
const PROJECT_ROOT_DIR string = "C:\\Users\\monan\\Desktop\\ohainworkspace"

type Project struct {

	// unique
	url string

	path string

}




func NewProject(url string) Project{
	path := PROJECT_ROOT_DIR + string(os.PathSeparator) + "driving-order"
	log.Printf("repo path %s",path)
	var project = Project{url,path}
	return project
}

// return path
// repoCode 由group + service组成
func (r Project) Sync(){

	osName := runtime.GOOS
	var scriptpath string
	if osName == "windows" {
		scriptpath = "C:\\Users\\monan\\Desktop\\all in one\\source code\\ohain\\repo\\script\\sync.bat"
	} else if osName == "linux" {
		scriptpath = "C:\\Users\\monan\\Desktop\\all in one\\source code\\ohain\\repo\\script\\sync.bat"
	}

	log.Printf("osName %s scriptpath %s",osName,scriptpath)
	// warning 路径不能带有空格
	var cmd *exec.Cmd = exec.Command(scriptpath,"http://git.51caocao.cn/java-newer/driving-order",r.path) 
	ParserCmd(cmd)


	//todo error deal
}


// output string
//srcCommit 不是属性，通过函数参数传递
func (r Project) Diff(srcCommit string,targetCommit string) string{
	osName := runtime.GOOS
	var scriptpath string
	if osName == "windows" {
		scriptpath = "C:\\Users\\monan\\Desktop\\all in one\\source code\\ohain\\repo\\script\\diff.bat"
	} else if osName == "linux" {
		scriptpath = "C:\\Users\\monan\\Desktop\\all in one\\source code\\ohain\\repo\\script\\diff.bat"
	}

	var cmd *exec.Cmd = exec.Command(scriptpath,"C:\\Users\\monan\\Desktop\\ohainworkspace\\driving-order",srcCommit,targetCommit)
	content := ParserCmd(cmd) 
	return content
}


// 
func (r Project) ReadFile(filepath string) []byte {
	path := PROJECT_ROOT_DIR + string(os.PathSeparator) + filepath
	log.Println(path)

	data,err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return data

}

// 缓存文件id与文件路径的关系

func ParserCmd(cmd *exec.Cmd) string{

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






