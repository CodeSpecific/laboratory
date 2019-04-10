package main

// 通过修改底层flag.CommandLine的形式定制命令参数容器

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	// 使用CommandLine进行定制化Usage，其实flag.Usage底层的方法也是使用的CommandLine
	// CommandLine是默认情况下的参数容器
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.StringVar(&name, "name", "vitamin", "everyone` name")
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	fmt.Printf("hello,%s \n", name)
}
