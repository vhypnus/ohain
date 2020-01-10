package project

import (
	"testing"
	"fmt"
	"os"
)
var dir , _= os.Getwd()
var testpath = dir + string(os.PathSeparator)+"test_data" +string(os.PathSeparator) 


func TestNewFile(t *testing.T) {
	var c string = `{System.out.print{{}}ln("hello world") ;}`
	var f *File = &File{c:c}
	f.mark()
	fmt.Println(f.mp)
	fmt.Printf("mp len %d cap %d\n",len(f.mp),cap(f.mp))
	fmt.Printf("mp len %d cap %d\n",len(f.mc),cap(f.mp))
}

func TestBlock(t *testing.T) {

	var f *File = NewFile(testpath + "block_test.java")
	fmt.Println(f.mp)
	fmt.Println(f.mc)
	fmt.Println('{')
	var s,e = f.Block(1)

	fmt.Printf("first level s %d e %d\n",s,e)

	s,e = f.Block(2)

	fmt.Printf("second level s %d e %d\n",s,e)
}

func TestOffset(t *testing.T) {
	// 边界场景测试，最后一个用户
	var c string = `{System.out.print{{}}ln("hello world") ;}`
	var f *File = &File{c:c}
	f.mark()
	var p = f.mp[len(f.mp)-1]
	fmt.Println(f.mp)
	fmt.Println(p)
	var o = f.Offset(p)
	fmt.Printf("offset %d\n",o)

	// 边界场景 第1个
	fmt.Println("===============case 2================")
	p = f.mp[len(f.mp)-len(f.mp)]
	fmt.Println(f.mp)
	fmt.Println(p)
	o = f.Offset(p)
	fmt.Printf("offset %d\n",o)

	//其它场景
	fmt.Println("===============case 3================")
	p = f.mp[len(f.mp)-len(f.mp)]
	fmt.Println(f.mp)
	fmt.Println(p)
	o = f.Offset(p)
	fmt.Printf("offset %d\n",o)
}

func TestForward(t *testing.T) {
	var c string = `{System.out.print{{}}ln("hello world") ;}`
	var f *File = &File{c:c}
	f.mark()
	var p = f.Forward(0,'{')
	fmt.Printf("forward %d\n",p)

	p = f.Forward(2,'{')
	fmt.Printf("forward %d\n",p)
}

func TestBackward(t *testing.T) {
	var c string = `{System.out.print{{}}ln("hello world") ;}`
	var f *File = &File{c:c}
	f.mark()
	var p = f.Backward(3,'{')
	fmt.Printf("backword %d\n",p)
}

func TestPreLine(t *testing.T) {
	var c = `public class OsBillInfo extends Model {}`
	//case 1:
	var f *File = &File{c:c}
	f.mark()
	var s,e = f.Block(1)
	fmt.Printf("s  %d e %d \n",s,e)

	s,e = f.PrevLine(s)
	fmt.Printf("s  %d e %d \n",s,e)

	fmt.Println("==================case 2============")
	//case 2:
	c = `import xxx.xxx.xxx
		 public class OsBillInfo extends Model {
		 }`
	f = &File{c:c}
	f.mark()
	s,e = -1,-1
	s,e = f.Block(1)

	s,e = f.PrevLine(s)
	fmt.Printf("s  %d e %d \n",s,e)


	fmt.Println("==================case 3============")
	//case 3:
	c = `import xxx.xxx.xxx
		 import aaa.bbb.ccc
		 public class OsBillInfo extends Model {
		 }`
	f = &File{c:c}
	f.mark()
	s,e = -1,-1
	s,e = f.Block(1)

	s,e = f.PrevLine(s)
	fmt.Printf("s  %d e %d \n",s,e)

}

func TestCurrentLine(t *testing.T) {
	var c = `public class OsBillInfo extends Model {}`
	//case 1:
	var f *File = &File{c:c}
	f.mark()
	var s,e = f.Block(1)
	fmt.Printf("s  %d e %d \n",s,e)

	s,e = f.CurrentLine(s)
	fmt.Printf("s  %d e %d \n",s,e)

	fmt.Println("==================case 2============")
	//case 2:
	c = `import xxx.xxx.xxx
		 public class OsBillInfo extends Model {
		 }`
	f = &File{c:c}
	f.mark()
	s,e = -1,-1
	s,e = f.Block(1)

	s,e = f.CurrentLine(s)
	fmt.Printf("s  %d e %d \n",s,e)


	fmt.Println("==================case 3============")
	//case 3:
	c = `import xxx.xxx.xxx
		 import aaa.bbb.ccc
		 public class OsBillInfo extends Model {
		 }`
	f = &File{c:c}
	f.mark()
	s,e = -1,-1
	s,e = f.Block(1)

	s,e = f.CurrentLine(s)
	fmt.Printf("s  %d e %d \n",s,e)
}


func TestNextLine(t *testing.T) {
	var c = `public class OsBillInfo extends Model {}`
	//case 1:
	var f *File = &File{c:c}
	f.mark()
	var s,e = f.Block(1)
	fmt.Printf("s  %d e %d \n",s,e)

	s,e = f.NextLine(s)
	fmt.Printf("s  %d e %d \n",s,e)

	fmt.Println("==================case 2============")
	//case 2:
	c = `import xxx.xxx.xxx
		 public class OsBillInfo extends Model {
		 	private Long money
		 }`
	f = &File{c:c}
	f.mark()
	s,e = -1,-1
	s,e = f.Block(1)

	s,e = f.NextLine(s)
	fmt.Printf("s  %d e %d \n",s,e)


	fmt.Println("==================case 3============")
	//case 3:
	c = `import xxx.xxx.xxx
		 import aaa.bbb.ccc
		 public class OsBillInfo extends Model {
		 	private Long money
		 }`
	f = &File{c:c}
	f.mark()
	s,e = -1,-1
	s,e = f.Block(1)

	s,e = f.NextLine(s)
	fmt.Printf("s  %d e %d \n",s,e)
}


func TestSubContent(t *testing.T) {
	var c = `public class OsBillInfo extends Model {System.out.println("hello world");}`
	//case 1:
	var f *File = &File{c:c}
	f.mark()
	var s,e = f.Block(1)

	fmt.Println(f.SubContent(s+1,e))
}


