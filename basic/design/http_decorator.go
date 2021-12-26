package design

import (
	"log"
	"net/http"
)

/*
	装饰者模式
*/

func WithLog(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("log handler")
		handler(writer, request)
	}
}

func WithServerHeader(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("test-with-server-header", "ok")
		handler(writer, request)
	}
}

func WithAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("Auth")
		if err != nil || cookie.Value != "Pass" {
			writer.WriteHeader(http.StatusForbidden)
			return
		}
		handler(writer, request)
	}
}
