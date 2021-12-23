package web

import (
	"log"
	"net/http"
	"testing"
)

func TestEcho(t *testing.T) {
	http.HandleFunc("/", Echo)

	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		log.Fatal("listen err ", err)
	}

}

func TestEcho2(t *testing.T) {
	http.Handle("/", timeMiddleware(http.HandlerFunc(Echo)))

	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		log.Fatal("listen err", err)
	}
}
