package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you are visiting : %s\n at port 8001", r.URL.Path)
}

func main() {
    http.HandleFunc("/", helloHandler)

    fmt.Println("Starting server on port 8001")
    err := http.ListenAndServe(":8001", nil)
    if err != nil {
        fmt.Println("Error starting server: ", err)
    }
}