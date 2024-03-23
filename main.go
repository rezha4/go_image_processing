package main

import (
    "fmt"
    "log"
    "net/http"
)

func convertHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement PNG to JPEG conversion
}

func resizeHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement image resizing
}

func compressHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement image compression
}

func main() {
    http.HandleFunc("/convert", convertHandler)
    http.HandleFunc("/resize", resizeHandler)
    http.HandleFunc("/compress", compressHandler)

    fmt.Println("Server listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
