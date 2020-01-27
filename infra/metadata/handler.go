package metadata

import (
	"encoding/json"
	"log"
	"net/http"
)

type handler struct{}

// NewHandler returns Handler interface for convience of wiring up to existing systems
func NewHandler() Handler {
	return &handler{}
}

func (h *handler) Metadata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	if err := enc.Encode(Get()); err != nil {
		log.Printf("could not encode json response for /metadata %v", err)
	}
}
