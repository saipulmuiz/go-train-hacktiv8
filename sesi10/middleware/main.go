package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Trace struct {
	TraceId      string
	Method       string
	ResponseTime string
	Path         string
	Request      string
}

func main() {
	http.HandleFunc("/", mid1(cors(greet)))
	http.HandleFunc("/hello", mid1(hello))

	port := ":4444"

	log.Println("server running at port", port)
	http.ListenAndServe(port, nil)
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
func mid1(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// now := time.Now()
		fmt.Println("mid1 called")
		next.ServeHTTP(rw, r)

		// end := time.Since(now)
		// log.Println("response time :", end.Seconds())
	}
}

/*
	func trace() -> middleware
		log -> method, response time, path, request, traceId

	type Log struct {

	}

	-> convert log to byte
	-> display in log.Println()
*/

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		next.ServeHTTP(w, r)
	}
}
