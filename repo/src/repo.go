/*
 * fileId  example java :package +className
 *
 *
 */

package core

import (
	"os"
	"io/ioutil"
)

const(
	ROOT_PATH = ""

	repoDict = {}
)


type Command interface {

	func Execute(args ... string)
}


// 同步
// return path
func (t *Command) Sync(repoCode string) string{
	// repoCode --> {repoUrl,repoPath,repoDir}
	repodir = repoUrl()

	//exec sych.sh
	var cmd = exec.Command(repoUrl,)
	cmd.Run()
	cmd.stdout
}

// output string
func (t *Command) Diff(srcCommit string,targetCommit string) string{

}


func ReadFile(fileId string) string {
	// get filepath by fileId

	var content 



}







