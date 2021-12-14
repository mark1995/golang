package multi_golang

import "fmt"

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- factor * i
	}
}

func Consumer(in <-chan int) {
	for value := range in {
		fmt.Println(value)
	}
}
