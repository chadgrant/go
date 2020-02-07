package schema

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type handler struct {
	registry Registry
}

func NewHandler(registry Registry) Handler {
	return &handler{registry}
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	ids := make([]string, 0)
	for k := range h.registry.Registered() {
		ids = append(ids, k)
	}

	if err := json.NewEncoder(w).Encode(ids); err != nil {
		log.Printf("could not encode schema keys json in handler %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	for k, v := range h.registry.Registered() {
		if strings.EqualFold(k, key) {
			if err := json.NewEncoder(w).Encode(v); err != nil {
				log.Printf("could not encode schema json in handler %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
