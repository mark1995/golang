### golang goroutine的生命周期管理
1. context管理goroutine的生命周期，在cancel方法掉用时，如果goroutine处于阻塞，没有办法select到ctx.Done（本质close(channel)）的消息，这个goroutine还是回收不掉
2. channel管理的问题
3. 
