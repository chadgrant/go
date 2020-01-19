package infra

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func DebugEnvironmentName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte(os.Getenv("ENVIRONMENT")))
	if err != nil {
		log.Println("[httperr] unable to write error to response writer")
	}
}

func DebugHeaders(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, r, r.Header)
}

func DebugTime(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, r, time.Now())
}

func DebugName(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	_, err := w.Write([]byte(hostname))
	if err != nil {
		log.Println("[httperr] unable to write error to response writer")
	}
}

func DebugError(w http.ResponseWriter, r *http.Request) {
	panic("testing error system")
}

func jsonResponse(w http.ResponseWriter, r *http.Request, o interface{}) {
	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(o)
	if err != nil {
		Error(w, r, http.StatusInternalServerError, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		log.Println("[debug-handler] unable to write error to response writer")
	}
}
