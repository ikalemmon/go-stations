package handler

import (
	"net/http"
	"fmt"
	"github.com/TechBowl-japan/go-stations/model"
	"encoding/json"
	"log"
	
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
// HealthHandlerに関する処理を書く。
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response = &model.HealthzResponse{}
	w.Message = `json:"OK"`
	//w.Header().Set("Content-Type", "application/json")
	// w.Message = "OK"
    json.NewEncoder(w).Encode(response)
	if (err != nil) {
		log.Println(err)
	}
}
