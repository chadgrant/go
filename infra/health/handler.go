package health

import (
	"encoding/json"
	"log"
	"net/http"
)

type handler struct {
	healthChecker HealthChecker
}

// NewHandler returds Handler interface for convience of wiring up to existing systems
func NewHandler(healthChecker HealthChecker) Handler {
	return &handler{healthChecker}
}

func (h *handler) Report(w http.ResponseWriter, r *http.Request) {
	rpt, rerr := h.healthChecker.Report()

	b, err := json.Marshal(rpt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	h.handle(w, r, rerr)

	_, err = w.Write(b)
	if err != nil {
		log.Println("healthcheck reporter could not write to the output writer")
	}
}

// Ready does not write any body, only status codes
func (h *handler) Ready(w http.ResponseWriter, r *http.Request) {
	_, err := h.healthChecker.Readiness()
	h.handle(w, r, err)
}

// Live does not write any body, only status codes
func (h *handler) Live(w http.ResponseWriter, r *http.Request) {
	_, err := h.healthChecker.Liveness()
	h.handle(w, r, err)
}

func (h *handler) handle(w http.ResponseWriter, r *http.Request, err error) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	status := http.StatusOK
	if err != nil {
		status = http.StatusServiceUnavailable
	}

	w.WriteHeader(status)
}
