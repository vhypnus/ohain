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
	
	// mark position
	mp []int

	// mark char
	mc []int

}




type Range struct {
	// start index
	s int

	// end index 
	e int

	// type:[),[],(],()
	t int

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
	// tmp mark index,tmp mark char 
	var tmp,tmc = make([]int,100),make([]int,100)
	var tmp = nil
	for i,c := range content {
		if c == '{' ||c == '}' || c == '"' || c == '\'' || c == '/' ||c == '*' || c == '\n' {
			
			tmp = append(tmp,c)
			tmc = append(tmc,c)
		}
	}

	// 
	var f = &File{c:content,mp:tmp,mc:tmc}
	return f
}

// subcontent
func (f *File) SubContent(s int,e int) string{
	return f.c[s:e]
}



// c ==> target char 
// s ==> start index
// ec ==> end char 
// -1 is not found
func (f *File) indexUntil(c int,s int ,ec int) int {

}

//return f.mc offset
func (f *File) index(p int) (o int) {
	//start,middle,end
	var s,o,e := 0,len(f.mi)/2,len(f.mi)

	for {
		if f.mp[o] > p {
			o = (s+o) /2 
			e = m
		} else if f.mp[o] < p {
			index = (o +e)/2
			s = o
		} else if f.mp[o] == p {
			break
		}
	}

}

//return f.mp position
func (f *File) forward(o int,c int)  (p int) {
	var min,max,p = o,len(f.mc),o

	for index := min ; index < max ;index ++ {
		f.mc[index] == c {
			p = index
			break
		}
	}
}

// 获取上一行
// p position
func (f *File) PrevLine(c char,p int) (s int,e int) {
	var offset := index(c,p)

	//start,end
	var s,e = offset,offset
	for index = offset ; index > 0 ; index-- {
		if f.mt[index] == '\n'{
			if e == p {
				e = index
			} else e < p {
				s = index
				break
			}
		}
	}
}

// 获取当前行
func (f *File) CurrentLine(c int,p int) (s int,e int) {

	//offset
	var o := index(c,p)
	//start,end,max
	var s,e,max,min = o,o,len(f.mt),0
	for index := o ; index > min ;index-- {
		if f.mt[index] == '\n'{
			if s == p {
				s = index
				break
			}
		}
	}

	for index = o ;index < max ;index++ {
		if f.c == '\n'{
			if e == p {
				e = index
				break
			}
		}
	}
}

func (f *File) NextLine(c char,p int) (s int,e int) {

	//offset
	var o := index(c,p)
	//start,end
	var s,e = p,p
	for index = o ;index < max ;index++ {
		if f.c == '\n'{
			if s == p {
				s = index
			} else s > p {
				e = index
				break
			}
		}
	}
}


/*
 * kmp算法
 * ss:substring
 */
func (f *File) IndexBackward(ss string ,p int) string {
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

func (f *File) IndexForward(ss string,p int) string {

}




//
func (f *File) Block(level int,blockChar int) (int,int) {
	//start index ,end index,l
	var s,e,l = -1,-1,0
	for index,item = range f.mt {
		if item.type = '{' && l = 0{
			l += 1
			s = mi[index]
		}else if item.type == '}' {
			if level = 1 {
				e = mi[index]
			} 
			
			l -= 1
		}
	}

	return s,e
}


func (f *File) Variable(s int,e int) []Variable {
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





