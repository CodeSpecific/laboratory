package main

import (
	"github.com/CodeSpecific/laboratory/visibility/testinternal"
	//"github.com/CodeSpecific/laboratory/visibility/testinternal/internal" //这里访问模块级别的包是不被允许的
)

//测试 Internal包的用法
func main() {
	testinternal.Hello()
	//internal.Name()
}
