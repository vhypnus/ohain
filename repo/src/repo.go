package core

import "os"

struct Repo interface {

	func get(repoCode string ,fp string) string 

	func sync() 

	//fp -> filepath
	func diff(repoCode string,fp string)

}

//
repoDict = {

}

// 同步
func (t *Repo) GitSync(repoCode string){
	// repoCode --> {repoUrl,repoPath,repoDir}
	repodir = repoUrl()

	//exec sych.sh
	var cmd = exec.Command(repoUrl,)
	cmd.Run()
	cmd.stdout
}

// output string
func (t * Repo) GitDiff(repoCode string) string{

}

// ignoreOutputFlag是否需要
func Exec(args,ignoreOutputFlag) string {

}

func AsynExec(args){

}



