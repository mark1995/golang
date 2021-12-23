package web

import (
	"io/ioutil"
	"log"
	"net/http"
)

// echo 是如何实现了可以作为handle接口实现者传入到ServeMux的muxEntry.handle，
// HandleFunc 类型 func(http.ResponseWriter, http.Request)
// 所以说 Echo函数是 HandleFunc类型，
// 并且这个 HandleFunc 类型 是Handle的鸭子类型，那么Echo也就是 Handle接口实现的一个实例
func Echo(wr http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wr.Write([]byte("echo error"))
		return
	}

	writeLen, err := wr.Write(msg)
	if err != nil || writeLen != len(msg) {
		log.Println(err, "write len: ", writeLen)
	}

	wr.Write([]byte("hello "))
}

// 如果涉及到比较复杂的 rest 类型的uri，原生的ServerMux是不能满足现有的需求的，
// 原生的ServerMux本质上只是一个map结构，不支持带有变量结构的rest风格

// 字典树，压缩动态检索树，检索树：用于字符串检索
