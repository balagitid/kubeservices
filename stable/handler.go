// stable/handler.go
package handler

import (
    "encoding/json"
    "net/http"
)

// VersionResponse represents the version response structure
type VersionResponse struct {
    Version string `json:"version"`
}

// StableVersionHandler handles requests for the stable version
func StableVersionHandler(w http.ResponseWriter, r *http.Request) {
    versionResponse := VersionResponse{
        Version: "1.0.0",
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(versionResponse)
}
