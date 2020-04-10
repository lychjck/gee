package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/nice", niceHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "url.path=%q\n", req.URL.Path)
}

func niceHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
}
