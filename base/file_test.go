package base

import (
	"testing"
	"fmt"
	"github.com/vhypnus/ohain/assert"
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

func TestCurrentLine(t *testing.T) {
	var rp = []int{1,8,18,28,49,66,77,88}
	var f = &File {rp:rp}
	var s,e = f.CurrentLine(28)
	fmt.Print(fmt.Sprintf("[%v ,%v]\n",s+1,e))

	//error case:
	s,e = f.CurrentLine(99)

	s,e = f.CurrentLine(-1)
}

func TestLastLine(t *testing.T) {
	var rp = []int{1,8,18,28,49,66,77,88}
	var f = &File {rp:rp}
	var s,e = f.LastLine(28)
	assert.Assert(8,s)
	assert.Assert(18,e)
}


func TestNextLine(t *testing.T) {

	var rp = []int{1,8,18,28,49,66,77,88}
	var f = &File {rp:rp}
	var s,e = f.NextLine(28)
	assert.Assert(49,s)
	assert.Assert(66,e)
}

func TestBlock(t *testing.T) {
	var content = `package helloworld;

	public class HelloWorld {

    	public static void main(String[] args) {
       		System.out.println("hello world") ;
    	}

    	public void test(){

    	}

	}`
	var f = parse(content)
	fmt.Println(f)
	

	var blocks = f.Block(1);
	fmt.Print(fmt.Sprintf("blocks %v\n",blocks))

	blocks = f.Block(2);
	fmt.Print(fmt.Sprintf("blocks %v\n",blocks))

}


func TestVariable(t *testing.T) {

	// case:没有包含子结构
	var content = `package helloworld;

	public class HelloWorld {

		// 名字
    	private String name ;

    	// 年龄
    	// 年龄
    	private int age ;

    	/*生日*/
    	private Date birth ;

    	/* 家乡

    	家乡*/
    	private String hometown ;


    	public void set{
    		
    	}


	}`

	var f = parse(content)

	var blocks = f.Block(1) ;
	fmt.Println(blocks)

	var variable = f.Variable(blocks[0]) ;

	fmt.Println(len(variable))
	fmt.Println(variable)

	// 接口
	content = `package helloworld;

	public interface HelloWorld {

		// 你好
		public String hello() ;

		// 请说人话
		public String say();


	}`

	f = parse(content)

	blocks = f.Block(1) ;

	variable = f.Variable(blocks[0]) ;

	fmt.Println(len(variable))
	fmt.Println(variable)
}

func TestNew(t *testing.T){
	// var c = "hello world"

	// for pos,char := range c {
	// 	 fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	// }
}