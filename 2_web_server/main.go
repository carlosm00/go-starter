package main

//Ref. 1
import (
    "fmt"
    "net/http"
)

// Ref. 2
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Your Golang-based Web Server is working properly!")
}

// Ref. 3
func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil)) 
}
