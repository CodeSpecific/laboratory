package internal

import "fmt"

func Name() {
	fmt.Print("我是模块级别的可见度，只能被子包访问和直接父节点访问")
}
