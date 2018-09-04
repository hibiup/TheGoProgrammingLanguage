package goroutine_sample

import (
	"fmt"
	"time"
)

func calcu_fib() {
	go spinner(100 * time.Millisecond) // 启动一个 spinner 的 goroutine 然后返回
	const n = 45
	fibN := fib(n) // 计算 fib 函数
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for { // 死循环，直到主程序退出
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
