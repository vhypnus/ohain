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
	var len,tl,rp,lbp,rbp,sp = len(content),0,make([]int,0,10),make([]int,0,10),make([]int,0,10),make([]int,0,10)
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



//小于但最临近的值
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


//大于但最临近的值
func (f *File) NextCharPos(char int,pos int) int {
	var arr,s,e,npos =f.rp,0, len(f.rp)-1,-1

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
		if mv - pos > 0 {
			e = m
			npos = mv
		} else if mv - pos < 0  {
			s = m+1
		} else {
			// 最快的找法
			npos = arr[m+1]
			break
		}
		
	}

	return npos
}


func (f *File) CurrentLine(pos int) (s int,e int) {

	var arr = f.rp

	if pos > arr[len(arr)-1] {
		panic(fmt.Sprintf("[%v] 大于数组 %v 最大值 [%v].\n",pos,arr ,arr[e]))
	}

	if pos < arr[0] {
		panic(fmt.Sprintf("[%v] 小于数组 %v 最小值 [%v].\n", pos,arr,arr[s]))
	}

	s ,e = f.LastCharPos('\n',pos),f.NextCharPos('\n',pos)
	return s,e
}

//
func (f *File) Block(level int) (s int ,e int) {
	// 代码规模，不允许穿插在中间

	// start position,end position
	var sp,ep = f.lbp[level-1],f.rbp[len(f.rbp)-1]
	fmt.Printf("%v  %v\n",sp,ep)

	s,_ = f.CurrentLine(sp)
	_,e = f.CurrentLine(ep)
	return s,e
}



// 两个Block的差
func (f * File) Delta(sl int ,vl int) (s int,e int) {
	return  -1,-1
}