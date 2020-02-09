package schema

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type handler struct {
	registry Registry
}

func NewHandler(registry Registry) Handler {
	return &handler{registry}
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	type schema struct {
		URI string `json:"uri"`
		URL string `json:"url"`
	}
	schemas := make([]schema, 0)
	for k := range h.registry.Registered() {
		schemas = append(schemas, schema{k, fmt.Sprintf("http://%s/schema?key=%s", r.Host, url.PathEscape(k))})
	}

	if err := json.NewEncoder(w).Encode(schemas); err != nil {
		log.Printf("could not encode schemas json in handler %v\n", err)
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
