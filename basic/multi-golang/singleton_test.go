package multi_golang

import (
	"fmt"
	"sync"
	"testing"
)

func TestMyOnce(t *testing.T) {
	myOnce := &MyOnce{}
	myOnce.Do(func() {
		fmt.Println("hello world")
	})
}

func TestInstance(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(10)
	for i:=0;i<10;i++ {
		go func() {
			instance := Instance()
			fmt.Printf("%p\n", instance)
			group.Done()
		}()
	}
	group.Wait()
}

func TestInstanceV2(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(10)
	for i:=0;i<10;i++ {
		go func() {
			instance := InstanceV2()
			fmt.Printf("%p\n", instance)
			group.Done()
		}()
	}
	group.Wait()
}