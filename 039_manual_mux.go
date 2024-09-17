package main

import (
	"fmt"
	"net/http"
)

type my_mux struct{}

func (p *my_mux) serve_http(w http.ResponseWriter, r *http.Request) {
	fmt.Println("init serve_http")
	if r.URL.Path == "/" {
		say_hello(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func say_hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	mux := &my_mux{}
	http.ListenAndServe(":12138", http.HandlerFunc(mux.serve_http))
}
