package goroutine_sample

import (
	"testing"
)

func TestSpinner(t *testing.T) {
	calcu_fib()
}

func TestTcpServer(t *testing.T) {
	addr := "localhost:8008"
	go tcp_server(addr)
	tcp_client(addr)
}

func TestDigitalPipe(t *testing.T) {
	DigitalPipe()
}

func TestLaunch(t *testing.T) {
	Launch()
}
