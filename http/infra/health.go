package infra

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	Healthy   = "OK"
	Unhealthy = "UNHEALTHY"
	Degraded  = "DEGRADED"
)

var tests = make([]func() HealthCheckTestResult, 0)

type HealthCheckReport struct {
	ReportAsOf time.Time `json:"reportAsOf"`
	Duration   int64     `json:"duration"`
	Interval   int       `json:"interval"`
	Tests      []HealthCheckTestResult
}

type HealthCheckTestResult struct {
	DurationMilliseconds int64     `json:"durationMilliseconds"`
	Name                 string    `json:"name"`
	Result               string    `json:"result"`
	TestedAt             time.Time `json:"testedAt"`
}

func RegisterHealthCheck(f func() HealthCheckTestResult) {
	tests = append(tests, f)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dur := int64(0)
	results := make([]HealthCheckTestResult, 0)
	for _, t := range tests {
		res := t()
		results = append(results, res)
		dur += res.DurationMilliseconds
	}
	rpt := &HealthCheckReport{
		ReportAsOf: time.Now().UTC(),
		Duration:   dur,
		Interval:   10,
		Tests:      results,
	}
	json.NewEncoder(w).Encode(rpt)
}

func Readiness(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	w.Header().Set("Content-Type", "text/plain")
}

func Liveness(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	for _, t := range tests {
		if t().Result != Healthy {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("FAILED"))
		}
	}
	w.Write([]byte("OK"))
}
