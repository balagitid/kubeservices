// latest/main.go
package main

import (
	"log"
	"net/http"
	"latest/handler"
)

func main() {
	http.HandleFunc("/api/version", handler.LatestVersionHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

