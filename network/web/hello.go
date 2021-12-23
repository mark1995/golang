package web

import (
	"log"
	"net/http"
	"time"
)

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)
		timeElapsed := time.Since(timeStart)
		log.Println("cost ", timeElapsed)
	})
}
