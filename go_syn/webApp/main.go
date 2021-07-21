package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello go", r.URL.Path)
}

func main() {

	http.HandleFunc("/", handler)

	//
	http.ListenAndServe(":7890", nil)
}
