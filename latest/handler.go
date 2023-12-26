// latest/handler.go
package handler

import (
	"encoding/json"
	"net/http"
)

// VersionResponse represents the version response structure
type VersionResponse struct {
	Version string `json:"version"`
}

// LatestVersionHandler handles requests for the latest version
func LatestVersionHandler(w http.ResponseWriter, r *http.Request) {
	versionResponse := VersionResponse{
		Version: "2.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(versionResponse)
}

