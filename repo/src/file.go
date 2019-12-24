package repo


// 关系类型
const (
	IMPORT = 1 

	DECLARE = 1 >> 1
	NEW = 1 >> 2
	USE =  1 >> 3
)

/*
 * fileleve root -> import level --> declare leve --> new level --> uselevel
 *
 *
 */
type File struct{

	path string

	data []byte
	// block format (level,start,end)
	// blocks array

	//notes format (start,end)
	// notes array  
}


// func (f File) NewFile(filepath string,data []byte) File {

// 	return 
// }




func (f File) GetBlock( level int) {

}


