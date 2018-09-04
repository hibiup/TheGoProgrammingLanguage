package goroutine_sample

/**
naturals    -> squares    -> 打印
0,1,2,3 ... -> 0,1,4,9... -> printer
*/

import "fmt"

func DigitalPipe() {
	naturals := make(chan int)   // 可以不指定缓存
	squares := make(chan int, 3) // 也可以指定缓存

	// Counter
	go func() {
		for x := 0; x < 100; x++ { // 死循环
			naturals <- x // 发送数字
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals { // range 关键词也可以访问 channel
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x, " ", len(squares), "/", cap(squares)) // 打印数字和缓存的剩余使用情况
	}
}
