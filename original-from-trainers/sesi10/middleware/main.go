package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := route()
	port := ":4444"

	log.Println("server running at port", port)
	http.ListenAndServe(port, mux)
}

func route() *http.ServeMux {
	endpoint := http.HandlerFunc(greet)
	mux := http.NewServeMux()
	mux.Handle("/", mid1(endpoint))
	mux.Handle("/hello", mid1(http.HandlerFunc(hello)))

	return mux
}

// application handler
func greet(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(map[string]string{
		"success": "true",
		"method":  r.Method,
	})
}

func hello(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(map[string]string{
		"success": "true",
		"method":  r.Method,
		"path":    r.URL.Path,
	})
}

// middleware
func mid1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		now := time.Now()

		next.ServeHTTP(rw, r)

		end := time.Since(now)
		log.Println("response time :", end.Seconds())
	})
}
