package main

import (
	"context"
	"log"
	"net/rpc"
	"sync"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("client connect server err ", err)
	}
	var response string
	err = client.Call("HelloService.Hello", "my first rpc ", &response)
	if err != nil {
		log.Println("rpc call error ", err)
	} else {
		log.Println("rpc response ", response)
	}

	var notice chan *rpc.Call
	var asyncResp string
	call := client.Go("HelloService.Hello", "it is a asyc call", &asyncResp, notice)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	var err1 error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			return
		case <-call.Done:
			err1 = call.Error
			return
		}
	}()
	wg.Wait()
	if err1 != nil {
		log.Println("hello service error ", err1)
	} else {
		log.Println("rpc async response ", asyncResp)
	}

}
