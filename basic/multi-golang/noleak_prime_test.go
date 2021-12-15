package multi_golang

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
	这里也会有goroutine泄漏问题，这个螺旋goroutine链接，channel全是阻塞的，导致goroutine不会走到 select case代码中
	需要一个一个chan chan的close掉
 */
func TestNoPrime(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := NoLeakGenNatural(ctx)
	for i := 0; i < 100; i++ {
		prime := <-ch
		//fmt.Printf(" %v %v\n",i+1, prime)
		ch = NoLeakPrimeFilter(ctx, ch, prime)
	}
	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
	fmt.Println("goroutine nums ", runtime.NumGoroutine())
	<- ch
	time.Sleep(10 * time.Second)
	fmt.Println("goroutine nums ", runtime.NumGoroutine())
}