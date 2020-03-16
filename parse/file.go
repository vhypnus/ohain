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



//查最临近
func (f *File) LastCharPos(char int ,pos int) int {
	
	var arr,s,e,lpos =f.rp,0, len(f.rp)-1,-1

	if pos > arr[e] {
		panic(fmt.Sprintf("[%v] 大于数组 %v 最大值 [%v].\n",pos,arr ,arr[e]))
	}

	if pos < arr[s] {
		panic(fmt.Sprintf("[%v] 小于数组 %v 最小值 [%v].\n", pos,arr,arr[s]))
	}


	//middle position,middle value
	for e > s {
		m := (s+e)/2
		mv := arr[m]
		if mv - pos < 0 {
			s = m+1
			lpos = mv
		} else if mv - pos > 0  {
			e = m
		} else {
			lpos = arr[m-1]
			break
		}
		
	}

	return lpos
}


//
func (f *File) NextCharPos(char int) int {
	
}