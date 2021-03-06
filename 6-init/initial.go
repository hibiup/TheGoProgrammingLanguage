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
	n = n+1
}

func init() {
	fmt.Printf("Second init(): sub.getNumber()=%d\n",  getNumber())
	n = n+1
}

func init() {
	// ":=" 会声明一个本地变量 "n" 掩盖（不是覆盖）全局的那个 n
	n := "string 0"
	fmt.Printf("Third init(): n=%s\n", n)
}

func main() {
	fmt.Printf("main() function: n=%d\n", n)
	fmt.Printf("sub.getNumber() function: n=%d\n", getNumber())
}
