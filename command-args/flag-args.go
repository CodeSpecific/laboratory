package main

import (
	"flag"
	"fmt"
	"os"
)

// flag的基本用法

var name string

// init 初始化命令行参数的形式一般适用于变量在外部或者是全局
func init() {
	flag.StringVar(&name, "name", "vitamin", "everyone`s name")
	// 自定义Usage信息，即命令源码文件的参数使用说明
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", "question")
		flag.PrintDefaults() //默认情况下Usage中要打印的信息
	}
}

func main() {
	flag.Parse()
	fmt.Printf("Hello,%s\n", name)
}
