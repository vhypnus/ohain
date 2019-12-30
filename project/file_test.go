package project

import (
	"testing"
	"fmt"
)


func TestInitFile(t *testing.T) {
	var code = `
		package com.caocao.hermes.order.client.base;

		import java.io.Serializable;

		/**
		 * @author: xuyongjian
		 * @date: 2018/4/27
		 */
		public class BaseQuery implements Serializable {
		    
		    private static final long serialVersionUID = -5375049797955766960L;
		}

	`

	var f *File = InitFile(code)
	fmt.Println(f)
}