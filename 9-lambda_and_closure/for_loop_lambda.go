package main

import (
	"fmt"
	"os"
)

func for_loop_lambda() {
	var rmdirs []func()
	for _, d := range []string{"a", "b", "c"} {
		fmt.Printf("x%X => %s\n", &d, d)
		dir := d                       // NOTE:for 循环的控制体只会初始化(分配)一次控制变量。因此只有一个 d 变量。
		os.MkdirAll(dir, 0755)         // creates parent directories too
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}

	// ...do some work…
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}