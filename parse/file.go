package file

import (
	"io/ioutil"
	"fmt"
)

type File struct {
	//文件路径
	fp string

	//文件内容
	c string

	//总行数
	tl int

	//换行位置
	rp []int

	//左括号位置
	lbp []int

	//右括号位置
	rbp []int

	//分号
	sp []int

}



func NewFile(path string) *File {

	return nil
}

func (f *File) readFile(path string) *[]byte {
	data,err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return &data
}


func parse(content string) *File {
	var tl,rp,lbp,rbp,sp = 0,make([]int,0,10),make([]int,0,10),make([]int,0,10),make([]int,0,10)
	for pos,char := range content {
		if char == '\n' {
			tl += 1
			rp = append(rp,pos)
		}

		if char == '{' {
			lbp = append(lbp,pos)
		}

		if char == '}' {
			rbp = append(rbp,pos)
		}

		if char == ';' {
			sp = append(sp,pos)
		}

	}

	//case:最后一行无回车情况
	tl += 1

	var f *File = &File{tl:tl,rp:rp,lbp:lbp,rbp:rbp,sp:sp}
	return f
}



//
func (f *File) LastCharPos(char int ,pos int) int {

	var arr = f.rp

	//二分法查找
	var s,e,lpos =0, len(f.rp)-1,-1

	//middle index,middle value
	for e > s {
		mi := (s+e-1)/2
		mv := arr[mi]
		if mv == lpos {
			break
		}
		if mv > pos {
			e = mi
		}  else if mv < pos {
			s = mi
			lpos = mv
		} 
		fmt.Printf("%v,%v,%v\n",s,e,lpos)
		// if e == 4 {
		// 	fmt.Printf("%v\n",(s+e-1)/2)
		// 	break
		// }
	}

	return lpos
}


//
func (f *File) NextCharPos(char int) int {
	return -1
}