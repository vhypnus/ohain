package project

import (
	"testing"
	"fmt"
	"os"
)
var dir , _= os.Getwd()
var f *File = NewFile(dir + string(os.PathSeparator) + "OsBillInfo_test.java")

func Test(t *testing.T) {
	fmt.Println(f.mp)
}



