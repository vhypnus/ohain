package file

import (
	"testing"
	"fmt"
)

func TestParser(t *testing.T) {
	var content = "helle\nworld"
	var file = parser(content)
	fmt.Println(file)
}

func TestNew(t *testing.T){
	// var c = "hello world"

	// for pos,char := range c {
	// 	 fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	// }
}