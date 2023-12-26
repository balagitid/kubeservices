// stable/main.go
package main

import (
    "log"
    "net/http"
    "github.com/balagitid/kubeservices/stable/handler"
)

func main() {
    http.HandleFunc("/api/version", handler.StableVersionHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
