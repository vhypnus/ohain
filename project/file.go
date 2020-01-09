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


	var content string = string(data)
	var f *File = &File{c:content}
	return f
	
}

func (f *File) mark() {
	// tmp mark index,tmp mark char 
	var tmp,tmc = make([]int,100),make([]int,100)
	for c := range f.c {
		if c == '{' ||c == '}' || c == '"' || c == '\'' || c == '/' ||c == '*' || c == '\n' {
			
			tmp = append(tmp,int(c))
			tmc = append(tmc,int(c))
		}
	}

	// 
	f.mp = tmp
	f.mc = tmc
}

// subcontent
// func (f *File) SubContent(s int,e int) string{
// 	return f.c[s:e]
// }



// c ==> target char 
// s ==> start index
// ec ==> end char 
// -1 is not found
// func (f *File) Until(c int,s int ,ec int) int {
// 	return 0
// }

//return f.mc offset
// func (f *File) Index(p int) int {
// 	//start,middle,end
// 	var s,m,e = 0,len(f.mp)/2,len(f.mp)

// 	for {
// 		if f.mp[m] > p {
// 			m = (s+m) /2 
// 			e = m
// 		} else if f.mp[m] < p {
// 			m = (m +e)/2
// 			s = m
// 		} else if f.mp[m] == p {
// 			break
// 		}
// 	}

// 	return m
// }

// //return f.mp position
// func (f *File) Forward(o int,c int) int {
// 	var min,max,p = o,len(f.mc),o

// 	for index := min ; index < max ;index ++ {
// 		if f.mc[index] == c {
// 			p = index
// 			break
// 		}
// 	}

// 	return p
// }

// func (f *File) Backward(o int,c int) int {
// 	return 0
// }

// // 获取上一行
// // p position
// func (f *File) PrevLine(c int,p int) (int,int) {
// 	var offset = Index(p)

// 	//start,end
// 	var s,e = offset,offset
// 	for index := offset ; index > 0 ; index-- {
// 		if f.mt[index] == '\n'{
// 			if e == p {
// 				e = index
// 			} else if e < p {
// 				s = index
// 				break
// 			}
// 		}
// 	}

// 	return s,e
// }

// // 获取当前行
// func (f *File) CurrentLine(c int,p int) (int,int) {

// 	//offset
// 	var o = Index(c,p)
// 	//start,end,max
// 	var s,e,max,min = o,o,len(f.mt),0
// 	for index := o ; index > min ;index-- {
// 		if f.mt[index] == '\n'{
// 			if s == p {
// 				s = index
// 				break
// 			}
// 		}
// 	}

// 	for index = o ;index < max ;index++ {
// 		if f.c == '\n'{
// 			if e == p {
// 				e = index
// 				break
// 			}
// 		}
// 	}

// 	return s,e
// }

// func (f *File) NextLine(c char,p int) (int,int) {

// 	//offset
// 	var o = index(c,p)
// 	//start,end
// 	var s,e = p,p
// 	for index = o ;index < max ;index++ {
// 		if f.c == '\n'{
// 			if s == p {
// 				s = index
// 			} else if s > p {
// 				e = index
// 				break
// 			}
// 		}
// 	}

// 	return s,e
// }


// /*
//  * kmp算法
//  * ss:substring
//  */
// func (f *File) Fackward(ss string ,p int) (int,int) {
// 	var s ,e,step = -1,-1,len(ss)
// 	for  index := p ;index > 0 ;index-- {
// 		if string(content[index-step:index]) == ss {
// 			s = index - step
// 			e = index 
// 			break
// 		}
// 	}

// 	return s,e
// }

// //
// func (f *File) Block(level int) (int,int) {
// 	//start index ,end index,l
// 	var s,e,l = -1,-1,0
// 	for index,item = range f.mt {
// 		if item == '{' && l == 0{
// 			l += 1
// 			s = mi[index]
// 		}else if item == '}' {
// 			if level == 1 {
// 				e = mi[index]
// 			} 
			
// 			l -= 1
// 		}
// 	}

// 	return s,e
// }


// func (f *File) Variable(s int,e int) []Variable {
// 	var v = make([]Variable ,16)
// 	//line start ,line end
// 	var ls ,le = -1 ,-1
// 	for p := s ; p < e ; p++ {
// 		if f.content[p] == ';' {
// 			le = p-1
// 		}

// 		if f.content[p] == '\r' {
// 			ls = p +1
// 			v = append(v,GetVariable(f.content[ls,le]))
// 		}

// 		//reset
// 		if f.content[p] == '\r' {
// 			ls ,le = -1 ,-1
// 		}
// 	}

// 	return v
// }





