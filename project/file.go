package project

import (
	"io/ioutil"
	"fmt"
	"container/list"
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
	mc []byte

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
	f.mark()
	return f
	
}

func (f *File) mark() {
	// tmp mark index,tmp mark char 
	var tmp,tmc = make([]int,0),make([]byte,0)
	fmt.Println()
	for p:= range f.c {
		var c = f.c[p]

		
		if c == '{' ||c == '}' || c == '\'' || c == '/' ||c == '*' || c == '\n' {
			tmp = append(tmp,int(p))
			tmc = append(tmc,c)
		}
	}

	f.mp = tmp
	f.mc = tmc
}

// c ==> target char 
// s ==> start index
// ec ==> end char 
// -1 is not found
// func (f *File) Until(c int,s int ,ec int) int {
// 	return 0
// }

// {{{}{}}}
func (f *File) Block(level int) (int,int) {
	//position channel
	var q = list.New()
	var s ,e = -1 ,-1
	
	for i,c := range f.mc {

		if c == '{' {
			q.PushBack(f.mp[i])
		} else if c == '}' {
			ele := q.Back()
			q.Remove(ele)
			if q.Len() == level - 1 {
				s = ele.Value.(int)
				e = f.mp[i]
			}

		}
	}

	return s ,e 

}

//return f.mc offset
func (f *File) Offset(p int) int {
	//min,max,offset
	var min,max ,o  = 0,len(f.mp),-1

	for max >= min {
		mid := (min + max) / 2
		if f.mp[mid] > p {
			max = mid - 1 
		} else if f.mp[mid] < p {
			min = mid + 1
		} else if f.mp[mid] == p {
			o = mid
			break
		}
	}

	return o
}

// return f.mp position
func (f *File) Forward(o int,c byte) int {
	var min,max,p = o,len(f.mc),-1

	for index := min ; index <= max ;index ++ {
		if f.mc[index] == c {
			p = f.mp[index]
			break
		}
	}

	return p
}

func (f *File) Backward(o int,c byte) int {
	var min,max,p = 0,o,-1

	for index := max ; index >= min ;index-- {
		if f.mc[index] == c {
			p = f.mp[index]
			break
		}
	}

	return p
}

// // 获取上一行
// // p position
func (f *File) PrevLine(p int) (int,int) {
	var o = f.Offset(p)

	//start,end
	var s,e = -1,-1
	for index := o ; index >= 0 ; index-- {
		if f.mc[index] == '\n'{
			if e == -1 {
				e = f.mp[index]
			} else {
				s = f.mp[index]
				break
			}
		}
	}

	if e > 0 && s < 0 {
		s = 0
	}

	return s,e
}

// // 获取当前行
func (f *File) CurrentLine(p int) (int,int) {

	//offset
	var o = f.Offset(p)
	//start,end,max
	var s,e,min,max = -1,-1,0,len(f.mc)-1
	for index := o ; index >= min ;index-- {
		if f.mc[index] == '\n'{
			if s == -1 {
				s = f.mp[index]
				break
			}
		}
	}

	for index := o ;index <= max ;index++ {
		if f.mc[index] == '\n'{
			if e == -1 {
				e = f.mp[index]
				break
			}
		}
	}

	if s ==-1 {
		s = 0
	}

	if e == -1 {
		e = len(f.c) - 1
	}

	return s,e
}



func (f *File) NextLine(p int) (int,int) {

	//offset
	var o = f.Offset(p)
	//start,end
	var s,e,min,max = -1,-1,o,len(f.mc)-1
	for index := min ;index <= max ;index++ {
		if f.mc[index] == '\n'{
			if s == -1 {
				s = f.mp[index]
			} else if e == -1 {
				e = f.mp[index]
				break
			}
		}
	}

	if s > 0 && e == -1 {
		e = len(f.c) - 1
	}

	return s,e
} 


// sub content
func (f *File) Sub(s int,e int) string{
	return f.c[s:e]
}

// 代码文件包含（注释,code）
// case 字符串包含注释符号,如果首先遇到的是代码，则返回-1.-1
// p position
func (f *File) Note(p int) (int,int){
	var o = f.Offset(p)

}

func (f *File) Code(p int)(int,int){

}

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





