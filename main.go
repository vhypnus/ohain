package main

import (
	"log"
	"fmt"
	"strings"
)

func main(){
	

}

func trimspace(){
	var s = "  hello world   "
	fmt.Println(s)
	fmt.Println(strings.TrimSpace(s))
}


func logtest(){
	log.Printf("hello world %v\n","huangyuewen")
}

func delslice(){
	//定义一个数字切片
    ageList := []int{1, 3, 7, 7, 8, 2, 5}

    //遍历删除6以下的
    for i := 0; i < len(ageList); {
        if ageList[i] < 2 {
        	fmt.Printf("before del:%v\n", ageList[:i])
        	fmt.Printf("before del:%v\n", ageList)
            ageList = append(ageList[:i], ageList[i+1:]...)
            fmt.Printf("after del:%v\n cap \v%", ageList,len(ageList))
        } else {
            i++
        }
    }
    
}