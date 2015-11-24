package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from version %s\n", os.Getenv("APPVERSION"))
}

func main() {  
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}