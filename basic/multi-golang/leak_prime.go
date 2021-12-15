package multi_golang

/*
	内存泄漏 goroutine 版的掉用
 */
func GenNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}


func PrimeFilter(in <- chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			if v % prime != 0 {
				out <- v
			}
		}
	}()
	return out
}


