package deferred

import (
	"time"
	"log"
)

/**
  defer机制常被用于记录何时进入和退出函数。
*/
func bigSlowOperation() {
	defer trace("bigSlowOperation")()    // defer trace 函数返回的闭包
	time.Sleep(3 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)          // 进入时立刻执行执行
	return func() {                      // 返回一个被 deferred 的闭包
		log.Printf("exit %s (%s)", msg,time.Since(start))   // defer 时执行
	}
}