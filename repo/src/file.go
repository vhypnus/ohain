<<<<<<< HEAD
package repo

import (
	"log"
	"io/ioutil"
)

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

	fileId string

	path string

	// block format (level,start,end)
	blocks array

	//notes format (start,end)
	notes array  
}


func (f File) NewFile(filepath string) File {

}


func (f File) GetBlock(int level) {
	
}


