package main

import (
	"net/http"
	"fmt"
)

func handle(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, "hello %s, Path: %s\n", r.URL,r.URL.Path)
}

func main(){
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}
