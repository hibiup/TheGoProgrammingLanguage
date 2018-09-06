package goroutine_sample

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

type Error struct {
	Error_code    int
	Error_message string
}

//////////////////////////////////
// 服务类
type Server struct {
	wg    *sync.WaitGroup // 用于注册工作线程
	ready chan struct{}   // 用于通知监控线程开始工作
}

// 主监控函数
func (s Server) monitor() {
	defer close(s.ready) // 保证 finish 关闭

	select {
	case <-s.ready:
		fmt.Println("Waiting for all job finish.")
		s.wg.Wait() // 挂起等待业务线程全部结束
		fmt.Println("WaitGroup reaches to 0")
	}
}

// 业务管理函数(业务函数参数也可以采用interface)
func (s Server) job_manager(job_func func()) {
	defer func() { // 业务监管线程 defer
		switch p := recover(); p.(type) { // 如果业务线程存在异常，处理之
		case Error:
			fmt.Printf("%s (panic)\n%s", p.(Error).Error_message, debug.Stack())
		default:
		}

		// 无论是否有异常，WaitGroup 为结束的业务线程减一
		s.wg.Done()
	}()

	job_func() // 调用业务函数(有些会抛出异常)
}

// 初始化函数
func (s *Server) init() {
	s.wg = &sync.WaitGroup{}
	s.ready = make(chan struct{})
}

// 服务启动函数
func (s Server) Start() {
	(&s).init()

	defer func() { // 主线程 defer
		for range s.ready {
			//退出前清除 channel(channel 必须已经关闭，否则会挂起)
		}
		fmt.Println("System exit!")
	}()

	// 启动监控线程等待所有线程结束
	go s.monitor()

	// 启动业务线程
	for i := 0; i < 10; i++ {
		(s.wg).Add(1)

		client := Client{i}
		indexed_client := client.Run // method values

		go s.job_manager(indexed_client)
		fmt.Println(i, "Business goroutine is being lunched!")
	}

	// 通知监控线程可以退出了（会保持至所有业务线程结束）
	fmt.Println("All job has been lunched. monitor goroutine is ready to work")
	s.ready <- struct{}{}
}

//////////////////////////////////
// main 入口函数
func Run(proc_number int) {
	runtime.GOMAXPROCS(proc_number) // 设置并行线程数
	server := Server{}
	server.Start()
}

//////////////////////////////////
// 业务函数接口
type Job interface {
	Run()
}

type Client struct {
	index int
}

// 业务函数实现
func (c Client) Run() {
	time.Sleep(1 * time.Second)
	fmt.Println(c.index, "Business goroutine is running...")

	if c.index%2 == 0 {
		panic(Error{1, fmt.Sprintf("%d Business goroutine failed!", c.index)}) // 故意制造故障
	} else { // 正常结束
		fmt.Println(c.index, "Business goroutine finished normally")
	}
}
