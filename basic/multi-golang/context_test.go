package multi_golang

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 *time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {

			err := Worker(ctx, &wg, i)
			if err != nil {
				fmt.Println(i, "err", err)
			}
		}(i)
	}
	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}
