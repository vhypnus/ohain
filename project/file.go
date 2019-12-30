package project

import (
	"io/ioutil"
)


// 关系类

/*
 * fileleve root -> import level --> declare leve --> new level --> uselevel
 *
 *
 */
type File struct{

	//file path
	fp string

	//full name index
	fn string

	//short name index
	sn string

	//content
	c string
	
	//left brace index
	lb []int

	//right brace index
	rb []int

	//left curves index
	lc []int

	//right curves index
	rc []int

	//dot  index
	d []int

	//'\n' index
	r []int

	// sigle quote index
	sq []int

	//double quote index
	dq []int

	// slash index
	s []int 

	// asterisk index 
	a []int

	//package index
	p []int

	//import index
	i []int
}




func NewFile(path string) *File{
	data,err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}


	content := string(data)
	var f *File = InitFile(content)
	return f
	
}

func InitFile(content string) *File{
	// index char
	// mark
	var lb,rb,lc,rc,d,r,a,s,ds,ss = make([]int,10),make([]int,10),make([]int,10),make([]int,10),make([]int,10),make([]int,10),make([]int,10),make([]int,10),make([]int,10),make([]int,10)
	for i,c := range content {
		if c == '{' {
			lb = append(lb,i)
		} else if c == '}' {
			rb = append(rb,i)
		} else if c == '(' {
			lc = append(lc,i)
		} else if c == ')' {
			rc = append(rc,i)
		} else if c == '.' {
			d = append(d,i)
		} else if c == '\n' {
			r = append(r ,i)
		} else if c == '*' {
			a = append(a ,i)
		} else if c == '/' {
			s = append(s,i)
		} else if c == '"' {
			ds = append(ds,i)
		} else if c == '\'' {
			ss = append(ss,i)
		}

	}

	var f = &File{"","","",content,lb,rb,lc,rc,d,r,a,s,ds,ss,make([]int,10),make([]int,10)}
	return f
}


