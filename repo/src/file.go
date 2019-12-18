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
struct File {

	fileId string

	path string

	// block format (level,start,end)
	blocks array
}


