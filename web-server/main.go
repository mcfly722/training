package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Welcome!")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}