package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var listenAddr string
	flag.StringVar(&listenAddr, "listen-addr", ":8080", "server listen address")
	flag.Parse()

	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Println("================")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, World!")
}
