package util

import (
	"fmt"
)

func Assert(expect int,actual int){
	if expect != actual {
		panic(fmt.Sprintf("expect [%v] actual [%v] \n",expect,actual))
	}
}
