package main

// Ref. 1
import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    // Ref. 2
    resp, err := http.Get("http://localhost:8080")
    if err != nil {
        fmt.Println("Failed to make request:", err)
        return
    }
    defer resp.Body.Close()

    // Ref. 3
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Failed to read response body:", err)
        return
    }

    // Ref. 4
    fmt.Println(string(body))
}
