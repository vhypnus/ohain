

// 关系类型
const (
	DEDENCY = 1 
	DECLARE = 1 >> 1
	NEW = 1 >> 2
	USE =  1 >> 3
)

/*
 * fileleve root -> import level --> declare leve --> new level --> uselevel
 *
 *
 */
type file struct {
	lineType
}

type line struct {

}

type block struct {

	startLine int 

	endLine int
}