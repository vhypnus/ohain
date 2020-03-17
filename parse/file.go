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

	//代码块
	blocks [][3]int

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
	var tl,rp,lbp,sp,blocks = 0,make([]int,0,10),make([]int,0,10),make([]int,0,10),make([][3]int,0,10)
	for pos,char := range content {
		if char == '\n' {
			tl += 1
			rp = append(rp,pos)
		}


		//
		if char == '{' {
			lbp = append(lbp,pos)
		}

		if char == '}' {
			var item = [3]int{len(lbp),lbp[len(lbp)-1],pos}
			lbp = append(lbp[:len(lbp)-1])
			blocks = append(blocks,item)
		}

		if char == ';' {
			sp = append(sp,pos)
		}

	}

	//最后一行
	if content[len(content)-1] != '\n' {
		tl += 1
		rp = append(rp,len(content)-1)
	}

	var f *File = &File{tl:tl,rp:rp,sp:sp,blocks:blocks}
	return f
}



//小于但最临近的值
func (f *File) LastCharPos(char int ,pos int) int {
	
	var arr,s,e,lpos =f.rp,0, len(f.rp)-1,-1

	if pos >= arr[e] {
		fmt.Print((fmt.Sprintf("[%v] 大于数组 %v 最大值 [%v].\n",pos,arr ,arr[e])))
		return arr[e]
	}

	if pos <= arr[s] {
		fmt.Print((fmt.Sprintf("[%v] 小于数组 %v 最小值 [%v].\n", pos,arr,arr[s])))
		return arr[s]
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

	if pos >= arr[e] {
		fmt.Print((fmt.Sprintf("[%v] 大于数组 %v 最大值 [%v].\n",pos,arr ,arr[e])))
		return arr[e]
	}

	if pos <= arr[s] {
		fmt.Print((fmt.Sprintf("[%v] 小于数组 %v 最小值 [%v].\n", pos,arr,arr[s])))
		return arr[s]
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

	s ,e = f.LastCharPos('\n',pos),f.NextCharPos('\n',pos)
	return s,e
}

//
func (f *File) Block(level int) [][]int {
	
	return nil
}



// 两个Block的差
func (f * File) BlockDelta(level int ,olevel int) (s int,e int) {
	return  -1,-1
}