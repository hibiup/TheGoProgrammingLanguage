package main

import (
	"fmt"
	"os"
)

func for_loop_lambda() {
	var rmdirs []func()
	for _, d := range []string{"a", "b", "c"} {
		fmt.Printf("x%X => %s\n", &d, d)   // NOTE:for 循环的控制体只会初始化(分配)一次控制变量。因此只有一个 d 变量。
		dir := d                           // 匿名函数共享了调用者的栈，因此匿名函数不会为 d 入栈，
		                                   // 因此必须分配一个新的局部变量来保存旧的值。
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)              // 否则匿名函数会始终得到最后一个 d 的内容。
		})
	}

	// ...do some work…
	for _, rmdir := range rmdirs {
		rmdir()
	}
}