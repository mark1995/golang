package multi_golang

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestProducer(t *testing.T) {
	ch := make(chan int, 64)
	go Producer(3, ch)
	go Producer(5, ch)
	go Consumer(ch)
	//time.Sleep(5 * time.Second)

	// 监听信号退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
