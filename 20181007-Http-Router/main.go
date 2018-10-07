package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type fooData struct {
	Bleh string `json:"blah"`
	Bar  string
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Index")
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s\n", p.ByName("name"))
}

func foo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var data fooData
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", data)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"success":true}`)
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/hello/:name", hello)
	router.POST("/foo", foo)
	http.ListenAndServe(":8080", router)
}
