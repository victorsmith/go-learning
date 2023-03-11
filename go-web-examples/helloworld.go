package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// W is written to => This is the message that will be sent to the client
	fmt.Fprintf(w, "Response: %v\nHost: %s\nURL: %s\n", r, r.Host, r.URL)
}

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	// })

	http.HandleFunc("/", handler)

	fmt.Println("Listening on port 80")
	http.ListenAndServe(":80", nil)
}
