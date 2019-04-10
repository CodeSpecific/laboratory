package main

// 通过自定义flag.FlagSet的形式定制命令参数容器

import (
	"flag"
	"fmt"
	"os"
)

var name string

// CommandLine其实就是flag包中一个全局变量，默认实现返回了FlagSet的实例
// 跟自己创建实例没有任何影响，这样可以更加灵活的定制命令参数容器
// 例子中的实现完全可以参考源代码中默认Command的实现
var cm *flag.FlagSet

func init() {
	cm = flag.NewFlagSet("question", flag.ExitOnError)
	cm.StringVar(&name, "name", "vitamin", "everyone`s name")
	cm.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", "question")
		cm.PrintDefaults()
	}
}

func main() {
	cm.Parse(os.Args[1:])
	fmt.Printf("Hello,%s\n", name)
}
