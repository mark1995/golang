package design

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestPublisher(t *testing.T) {
	publisher := NewPublisher(5*time.Millisecond, 64)
	// 订阅一个全部
	all := publisher.Subscribe()
	golang := publisher.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	go func() {
		//for {
		//	select {
		//	case msg := <- all:
		//		fmt.Println("all subscribe receive message", msg)
		//	}
		//}
		for msg := range all {
			fmt.Println("all subscribe receive message ", msg)
		}
		fmt.Println("all subscribe be closed")
	}()

	go func() {
		//for {
		//	select {
		//	case msg := <- golang:
		//		fmt.Println("golang subscribe receive message ", msg)
		//	}
		//}
		for msg := range golang {
			fmt.Println(" golang subscribe receive message ", msg)
		}
		fmt.Println("golang subscribe be closed")
	}()

	for i := 0; i < 10; i++ {
		publisher.Publish(fmt.Sprintf("hello world %v", i))
		publisher.Publish(fmt.Sprintf("hello golang %v", i))
	}
	publisher.Close()
	fmt.Printf("%#v", publisher)
	time.Sleep(1 * time.Second)
}
