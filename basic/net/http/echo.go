package main

import (
	"net/http"
)

func echo(wr *http.ResponseWriter, r *http.Reqeust) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wr.Write([]byte("echo error"))
		return
	}

	writeLen, err := wr.Write(msg)
}

func main() {
	http.HandleFunc("/", echo)
}