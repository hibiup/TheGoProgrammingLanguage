package goroutine_sample

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

// TCP server
func tcp_server(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close() // defer 关闭连接

	for { // 死循环
		time.Sleep(1 * time.Second)                                     // 睡眠一秒
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n")) // 向客户端发送时间戳
		if err != nil {
			return // e.g., client disconnected
		}
	}
}

// Client
func tcp_client(address string) { // target = "localhost:8000"
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	// 我们需要让主 goroutine 等待后台 goroutine 完成工作后再退出，我们使用了一个 channel 来同步两个 goroutine
	done := make(chan struct{}) // make 一个收发空结构体的 channel

	// 启动一个后台 goroutine 负责接受来自服务端的数据
	go func() {
		io.Copy(os.Stdout, conn) // (阻塞)接受来自服务端的输入写到标准输出，直到连接或标准输出中断
		log.Println("done")
		done <- struct{}{} // 向 main goroutine 发送信号
	}()

	send(conn, os.Stdin) // 尝试(阻塞)从标准输入读取内容发送到服务端，直到连接或标准输入中断
	conn.Close()         // send() 函数返回后调用 conn.Close() 确保关闭读和写方向的网络连接。
	<-done               // 读取 channel，直到确保后台 goroutine 发送了信号才结束 main goroutine。
}

func send(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil { // (阻塞)
		log.Fatal(err)
	}
}
