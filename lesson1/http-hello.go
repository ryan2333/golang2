package main

import (
	"net/http"
	"fmt"
)

func handle(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, "hello %s, Path: %s\n", r.Form.Get("user"),r.URL.Path)
}

func main(){
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}
