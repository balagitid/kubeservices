// latest/main.go
package main

import (
    "log"
    "net/http"
    "github.com/balagitid/kubeservices/latest/handler"
)

func main() {
    http.HandleFunc("/api/version", handler.LatestVersionHandler)
    log.Fatal(http.ListenAndServe(":8081", nil))
}
