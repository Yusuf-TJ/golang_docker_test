package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World! from a golang application")
}

func main() {
    http.HandleFunc("/", helloHandler)
    if err := http.ListenAndServe(":8000", nil); err != nil {
        panic(err)
    }
}
