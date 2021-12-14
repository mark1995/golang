package design

import (
	"sync"
	"time"
)

type (
	// 订阅者为一个通道
	subscriber chan interface{}

	// 主题为一个过滤器
	topicFunc func(v interface{}) bool
)

type Publisher struct {
	m sync.RWMutex
	// 发布者长度
	buffer int32
	// 超时时间
	timeout time.Duration
	//订阅者信息
	subscribers map[subscriber]topicFunc
}

func NewPublisher(publishTimeout time.Duration, buffer int32) *Publisher {
	return &Publisher{
		m:           sync.RWMutex{},
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

func (p *Publisher) SubscribeTopic(topic topicFunc) subscriber {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// 订阅全部 信息
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// 退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	wg := sync.WaitGroup{}
	for sub, topicFunc := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topicFunc, v, &wg)
	}
	wg.Wait()
}

// 发送topic
func (p *Publisher) sendTopic(sub subscriber, f topicFunc, v interface{}, w *sync.WaitGroup) {
	defer w.Done()

	if f != nil && !f(v) {
		return
	}
	// 超时控制
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub, _ := range p.subscribers {
		close(sub)
	}
	p.subscribers = make(map[subscriber]topicFunc)
}
