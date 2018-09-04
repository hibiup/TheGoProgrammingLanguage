package goroutine_sample

import (
	"fmt"
	"os"
	"time"
)

func Launch() {
	// Create abort channel
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	// Read count down from time.Tick channel
	fmt.Println("Commencing countdown.")

	for countdown := 10; countdown > 0; countdown-- {
		select { // select 会监视 block 内所有的 channel，直到任意一个有返回值
		case t := <-time.After(1 * time.Second): // time.After函数会先返回一个 channel，并起一个新的 goroutine 在经过特定的时间后向该channel发送一个独立的值。
			fmt.Println(countdown)
			fmt.Println(t)
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
		//fmt.Println(countdown)
		//fmt.Println(<-tick)
	}
}
