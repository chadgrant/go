package infra

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

func DebugEnvironment(w http.ResponseWriter, r *http.Request) {
	keys := os.Environ()
	arr := make([]string, len(keys))
	for i, v := range keys {
		arr[i] = strings.Split(v, "=")[0]
	}
	jsonResponse(w, r, arr)
}

func DebugError(w http.ResponseWriter, r *http.Request) {
	panic("testing error system")
}

func DebugFileDates(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		Error(w, r, http.StatusInternalServerError, err)
		return
	}

	type fd struct {
		Name     string    `json:"name"`
		Modified time.Time `json:"modified"`
		Size     int64     `json:"size"`
	}

	arr := make([]fd, len(files))
	for i, f := range files {
		arr[i] = fd{f.Name(), f.ModTime(), f.Size()}
	}

	jsonResponse(w, r, arr)
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
