package project



const (
	ROOT = "/data"
	SCRIPT_ROOT = "/data/script/"

)


type Project struct {

	name string

	path string

	url string

	modelpath string

	apipath string
} 



func NewRepository(projectName string) *Project {
	return nil
}





// 获取整个文件
// func (r *Project) folder(branch string) string {
// 	var fileinfos,err = ioutil.ReadDir(r.apipath)
// 	if err != nil{
// 		//todo error
// 	}

// 	for _, file := range fileinfos {
		
// 	}
// }