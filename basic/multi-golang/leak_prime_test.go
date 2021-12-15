package multi_golang

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestLeakPrime(t *testing.T) {
	ch := GenNatural()
	for i:=0; i<100; i++ {
		prime := <-ch
		fmt.Printf("%v : %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
	}

	go func() {
		for {
			select {
			case value := <- ch:
				fmt.Println("v ",value)
			}
		}
	}()

	time.Sleep(time.Second)
	// 这里会有 goroutine的泄漏问题，生成了100个goroutine,并没有管控goroutine的生命周期，会导致goroutine的泄漏
	fmt.Println("goroutine nums ", runtime.NumGoroutine())
}
