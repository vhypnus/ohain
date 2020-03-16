package file

import (
	"testing"
	"fmt"
)

func TestParse(t *testing.T) {
	var content = `package helloworld;

	public class HelloWorld {

    	public static void main(String[] args) {
       		System.out.println("hello world") ;
    	}

	}`
	var file = parse(content)
	fmt.Println(file)
}


func TestLastCharPos(t *testing.T) {
	var rp = []int{1,8,18,28,49,66,77,88}
	var f = &File {rp:rp}
	var pos = f.LastCharPos('\n',28)
	fmt.Println(pos)

	pos = f.LastCharPos('\n',19)
	fmt.Println(pos)

	pos = f.LastCharPos('\n',50)
	fmt.Println(pos)

	pos = f.LastCharPos('\n',49)
	fmt.Println(pos)

	//error case:
	pos = f.LastCharPos('\n',99)
	fmt.Println(pos)

	pos = f.LastCharPos('\n',-1)
	fmt.Println(pos)

}

func TestNextCharPos(t *testing.T) {
	var rp = []int{1,8,18,28,49,66,77,88}
	var f = &File {rp:rp}
	var pos = f.NextCharPos('\n',28)
	fmt.Println(pos)

	pos = f.NextCharPos('\n',19)
	fmt.Println(pos)

	pos = f.NextCharPos('\n',50)
	fmt.Println(pos)

	pos = f.NextCharPos('\n',49)
	fmt.Println(pos)

	//error case:
	pos = f.NextCharPos('\n',99)
	fmt.Println(pos)

	pos = f.NextCharPos('\n',-1)
	fmt.Println(pos)

}

func TestNew(t *testing.T){
	// var c = "hello world"

	// for pos,char := range c {
	// 	 fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	// }
}