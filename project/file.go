package project

import (
	"io/ioutil"
)


// mark constant
const (

	// left brace '{'
	LB = 1

	//right brace '{'
	RB = 1 << 1

	// sigle quote(') 
	SQ = 1 << 2

	//double quote("") 
	DQ = 1 << 3

	// slash(/) 
	S = 1 << 4 

	// back slash(/) 
	BS = 1 << 5

	// * asterisk  
	A = 1 << 6

	// \n
	R = 1 << 7

	// package key word
	P = 3 

	// import key word

	// new key word
)

// 关系类
/*
 * fileleve root -> import level --> declare leve --> new level --> uselevel
 * 
 *
 */
type File struct{

	//extension 
	e int 

	//file path
	fp string

	//full name
	fn string

	//short name
	sn string

	//content
	c string
	
	// mark
	m []Mark

}


type Mark struct {
	// index 
	i int

	// type
	t int 

	// preview
	prev *Mark

	// next
	next *Mark
}


func NewFile(path string) *File{
	data,err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}


	content := string(data)
	var f *File = mark(content)
	return f
	
}

func mark(content string) *File{
	// index char
	// mark
	var mi,mt = make([]int,100),make([]int,100)
	var tmp = nil
	for i,c := range content {
		if c == '{' ||c == '}' || c == '"' || c == '\'' || c == '/' ||c == '*' || c == '\n' {
			if tmp == nil {

			}
			m = append(m,c)
		}

	}

	// 
	var f = &File{c:content,m:m}
	return f
}


func (f *File) SubContent(s int,e int) string{
	return f.c[s:e]
}




// 获取上一行
func (f *File) GetPrevLine(p int) string {
	//start,end
	var s,e = p,p
	for index := range(p,len(f.content),-1) {
		if f.c == '\n'{
			if e == p {
				e = index
			} else e < p {
				s = index
				break
			}
		}
	}
	return f.c[s:e]
}

// 获取当前行
func (f *File) GetCurrentLine(p int) string {
	//start,end
	var s,e = p,p
	for index := range(p,len(f.content),-1) {
		if f.c == '\n'{
			if s == p {
				s = index
				break
			}
		}
	}

	for index := range(p,len(f.content),1) {
		if f.c == '\n'{
			if e == p {
				e = index
				break
			}
		}
	}
	return f.c[s:e]
}

func (f *File) GetNextLine(p int) string {
	//start,end
	var s,e = p,p
	for index := range(p,len(f.content),1) {
		if f.c == '\n'{
			if s == p {
				s = index
			} else s > p {
				e = index
				break
			}
		}
	}
	return f.c[s:e]
}


/*
 * kmp算法
 * ss:substring
 */
func (f *File)GetPre(ss string ,p int) string{
	var s ,e = -1,-1
	for  i := range(p,0,-1) {
		if string(content[p-1:p]) == '*/' {
			s = i 
		}

		if s >= 0 && string(content[p-1:p]) == '/*' {
			e = p-1
			break
		}
	}

	return s,e
}

func (f *File)GetNext(ss string,p int) string {
	
}




//
func (f *File) GetBlock(level int) (int,int) {
	//start index ,end index,l
	var s,e,l = -1,-1,0
	for _,item = range f.m {
		if item.type = '{' && l = 0{
			l += 1
			s = item.i
		}else if item.type == '}' {
			if level = 1 {
				e = item.i
			} 
			
			l -= 1
		}
	}

	return s,e
}


func (f *File) GetVariable(s int,e int) []Variable {
	var v = make([]Variable ,16)
	//line start ,line end
	var ls ,le = -1 ,-1
	for p := range(s,e,1) {
		if f.content[p] == ';' {
			le = p-1
		}

		if f.content[p] == '\r' {
			ls = p +1
			v = append(v,GetVariable(f.content[ls,le]))
		}

		//reset
		if f.content[p] == '\r' {
			ls ,le = -1 ,-1
		}
	}

	return v
}





