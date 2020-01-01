package infra

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

func DebugEnvironmentName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(os.Getenv("ENVIRONMENT")))
}

func DebugHeaders(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, r, r.Header)
}

func DebugTime(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, r, time.Now())
}

func DebugName(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	w.Write([]byte(hostname))
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

	w.Write(b)
}
