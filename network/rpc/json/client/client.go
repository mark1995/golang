package main

import (
	"context"
	"errors"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"
	"time"
)

func main() {
	client, err := jsonrpc.Dial("tcp", ":1235")
	if err != nil {
		log.Fatal("client connect error {}", err)
	}
	resp1, resp2 := "", ""
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := client.Call("HelloService.Hello", "goroutine 1", &resp1)
		if err != nil {
			log.Println("err ", err)
		}
	}()
	var err2 error
	go func() {
		defer wg.Done()
		var ch chan *rpc.Call
		call := client.Go("HelloService.Hello", "goroutine 2", &resp2, ch)
		select {
		case <-ctx.Done():
			// 超时了
			err2 = errors.New("timeout")
		case <-call.Done:

		}
	}()

	wg.Wait()
	log.Printf("reuslt %v \t %v", resp1, resp2)
}
