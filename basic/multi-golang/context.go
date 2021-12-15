package multi_golang

import (
	"context"
	"fmt"
	"sync"
)

// 一般如果有 context传递上下文，或者管理goroutine的生命周期的话，context会放在第一个参数
func Worker(ctx context.Context, wg *sync.WaitGroup, i int) error {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Println("hello ", i)
		}
	}
}




