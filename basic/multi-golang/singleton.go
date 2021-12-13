package multi_golang

import (
	"sync"
	"sync/atomic"
)

type singletone struct {

}

var (
	instance *singletone
	initialized uint32
	mu sync.Mutex
)

/*
	单例模式的实现方式,细心的人可能发现了，这个代码和标准库中的 sync.Once的实现方式很相识
 */

func Instance() *singletone {
	if atomic.LoadUint32(&initialized) ==  1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.AddUint32(&initialized, 1)
		instance = &singletone{}
	}
	return instance
}

var myOnce MyOnce
func InstanceV2() *singletone  {
	myOnce.Do(func() {
		instance = &singletone{}
	})
	return instance
}

/*
   sync.once
 */

type MyOnce struct {
	m sync.Mutex
	done uint32
}


func (once *MyOnce) Do (f func()) {
	if atomic.LoadUint32(&once.done) == 1 {
		f()
	}

	once.m.Lock()
	defer once.m.Unlock()

	if once.done == 0 {
		defer atomic.AddUint32(&once.done, 1)
		f()
	}
}




