package infra

import (
	"fmt"
	"net/http"
	"time"
)

type healthCheck struct {
	ReportAsOf time.Time     `json:"reportAsOf"`
	Duration   time.Duration `json:"duration"`
	Interval   int           `json:"interval"`
	Tests      []healthCheckTest
}

type healthCheckTest struct {
	DurationMilliseconds int64     `json:"durationMilliseconds"`
	Name                 string    `json:"name"`
	Result               string    `json:"result"`
	TestedAt             time.Time `json:"testedAt"`
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Health")
}

func goodToGo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	w.Header().Set("Content-Type", "text/plain")
}

func serviceCanary(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	//TODO: run health tests
	w.Header().Set("Content-Type", "text/plain")
}