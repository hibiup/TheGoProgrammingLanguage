package main

import (
	"fmt"
)

var n int

/**
init初始化函数不能被调用或引用外，其他行为和普通函数类似。init初始化函数在程序开始执行时按照它们声明的顺序被自动调用。
*/
func init() {
	fmt.Printf("First init(): n=%d\n", n)
	n = 1
}

func init() {
	fmt.Printf("Second init(): n=%d\n", n)
	n = 2
}

func main() {
	fmt.Printf("main() function: n=%d\n", n)
	fmt.Printf("sub.getNumber() function: n=%d\n", getNumber())
}
