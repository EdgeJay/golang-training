package main

import (
	"fmt"
	"net/http"
)

// Handler can be anything that implments
// ServeHTTP(w http.ResponseWriter, r *http.Request)
type anyHTTP int

func (aHttp anyHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	var handler anyHTTP
	http.ListenAndServe(":8181", handler)
}
