package main

import (
	"fmt"
	"net/http"
)

func hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello")
}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}
