package multi_golang

import (
	"context"
	"fmt"
)

func NoLeakGenNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("produce done")
				close(ch)
				return
			default:
				ch <- i
			}
		}
	}()
	return ch
}

func NoLeakPrimeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case value := <- ctx.Done():
				close(ch)
				fmt.Println("cancel ", value)
				return
			default:
				if v := <- in; v % prime != 0 {
					ch <- v
				}

			}
		}
	}()
	return ch
}