package file

import (
	"io/ioutil"
)

func NewFile(path string) *File {

	return nil
}

func readFile(path string) *[]byte {
	data,err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return &data
}

func parser(content string) *File {
	var tl,rl = 0,make([]int,0,10)
	for pos,char := range content {
		if char == '\n' {
			tl += 1
			rl = append(rl,pos)
		}

		if char == '{' {

		}

		if char == '}' {
			
		}

	}
	var f *File = &File{tl:tl,rl:rl}
	return f
}