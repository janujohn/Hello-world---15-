package main

import (
    "net/http" 
	    "fmt"
		)

func handler(w http.ResponseWriter, r *http.Request) { 
	    fmt.Println("Inside handler")
		    fmt.Fprintf(w, "Hello world from my Go program!")
}

func main() {
	    http.HandleFunc("/", handler) 
		    http.ListenAndServe("localhost:9000", nil) 
}
