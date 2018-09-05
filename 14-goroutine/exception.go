package goroutine_sample

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Error struct {
	Error_code    int
	Error_message string
}

func Exception() {
	runtime.GOMAXPROCS(2) // 设置并行线程数
	wg := sync.WaitGroup{}
	finish := make(chan struct{}) // 用于主线程接受退出通知

	defer func() { // 主线程 defer
		for range finish {
			//退出前清除 channel(channel 必须已经关闭，否则会挂起)
		}
		fmt.Println("System exit!")
	}()

	// 启动监控县城等待所有线程结束
	wg.Add(1) // 为监控线程自己加一
	go monitor(finish, &wg)

	// 启动业务线程
	for i := 0; i < 10; i++ {
		wg.Add(1) // waitGroup 加一
		go job_manager(i, &wg)
	}

	// 通知监控线程可以退出了（会保持至所有业务线程结束）
	fmt.Println("All job has been lunched. main goroutine is eligabled to exit")
	finish <- struct{}{}
}

// 主监控线程
func monitor(finish chan struct{}, wg *sync.WaitGroup) {
	defer close(finish) // 保证 finish 关闭

	select {
	case <-finish:
		(*wg).Done() // 如果收到“结束"消息就给自己减一
	}

	fmt.Println("Waiting for all job finish.")
	(*wg).Wait() // 挂起等待业务线程全部结束
	fmt.Println("WaitGroup reaches to 0")
}

// 业务管理线程
func job_manager(index int, wg *sync.WaitGroup) {
	defer func() { // 业务监管线程 defer
		switch p := recover(); p.(type) { // 如果业务线程存在异常，处理之
		case Error:
			fmt.Printf("panic: %s\n", p.(Error).Error_message)
		default:
		}

		// 无论是否有异常，WaitGroup 为结束的业务线程减一
		(*wg).Done()
	}()

	fmt.Println(index, "Business goroutine has been lunching!")
	client(index) // 调用业务函数(有些会抛出异常)
	fmt.Println(index, "Business goroutine finished by normal")
}

// 业务函数
func client(index int) {
	time.Sleep(1 * time.Second)
	if index%2 == 0 {
		panic(Error{1, fmt.Sprintf("%d Business goroutine failed!", index)}) // 故意制造故障
	} else {
		fmt.Println(index, "Business goroutine is running...") // 正常结束
	}
}
