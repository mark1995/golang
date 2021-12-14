/*
	wait_group 模仿
*/
package multi_golang

import (
	"fmt"
	"sync"
	"time"
)

func MockWaitGroup() {
	// 带10个缓存
	done := make(chan int, 10)

	for i := 0; i < cap(done); i++ {
		go func(i int) {
			fmt.Println(" hello world ", i)
			done <- 1
		}(i)
	}

	/*
		对通道的第K个接受完成操作发生在第K+c个发生操作完成之前，其中c位通道的缓存大小，
		happen-before, 主要还是为了保住程序的时序
		如果通道的长度为10，对通道的第1个接受完成操作一定是发生在第11个发生操作完成之前，想当然也知道，
	*/
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

func WaitGroup() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(" print ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func NonCacheChannel() {
	done := make(chan int)

	for i := 0; i < 10; i++ {
		go func(i int) {
			done <- i
			fmt.Println(" i send ok ", i, time.Now().UnixNano())
		}(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(" i receive ", <-done, time.Now().UnixNano())
	}
}
