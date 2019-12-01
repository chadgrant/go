package infra

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

const (
	Healthy   = "HEALTHY"
	Unhealthy = "UNHEALTHY"
	Degraded  = "DEGRADED"
)

type (
	HealthCheck func(context.Context) *HealthCheckTestResult

	HealthCheckReport struct {
		ReportAsOf time.Time                `json:"reportAsOf"`
		Duration   int64                    `json:"duration"`
		Interval   int                      `json:"interval"`
		Tests      []*HealthCheckTestResult `json:"tests"`
	}

	HealthCheckTestResult struct {
		DurationMilliseconds int64     `json:"durationMilliseconds"`
		Name                 string    `json:"name"`
		Result               string    `json:"result"`
		TestedAt             time.Time `json:"testedAt"`
	}
)

var tests = make([]HealthCheck, 0)

func RegisterHealthCheck(healthcheck HealthCheck) {
	tests = append(tests, healthcheck)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rpt := &HealthCheckReport{
		ReportAsOf: time.Now().UTC(),
		Duration:   0,
		Interval:   10,
		Tests:      make([]*HealthCheckTestResult, 0),
	}
	json.NewEncoder(w).Encode(rpt)
}

func Readiness(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	w.Header().Set("Content-Type", "text/plain")
}

func Liveness(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}
