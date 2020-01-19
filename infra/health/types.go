package health

import (
	"context"
	"net/http"
	"time"
)

type (
	// Check is a health check. Any error returned is considered unhealthy. Nil is healthy
	Check func() error

	// HealthChecker is an interface for defining methods of adding and
	// reporting of health checks
	HealthChecker interface {
		// AddLiveness adds a check that indicates that this instance of the
		// application should be destroyed or restarted. A failed liveness check
		// indicates that this instance is unhealthy, not some upstream dependency.
		// Every liveness check is also included as a readiness check.
		// Executed lazily (when report requested)
		// name - name of the health check
		// interval - how long the report should be considered valid (to prevent ddosing health checks)
		// check - the function to be checked
		AddLiveness(name string, interval time.Duration, check Check)

		// AddLivenessBackground same as AddLiveness but executes on an interval
		// ctx: Context for cancellation
		// name: Name of the health check
		// duration: How often the check should be run
		// If your check takes longer than the interval to execute, the next execution will happen immediately.
		// name - name of the health check
		// interval - how long the report should be considered valid (to prevent ddosing health checks)
		// check - the function to be checked
		AddLivenessBackground(name string, interval time.Duration, check Check)

		// AddLivenessBackgroundWithContext same as AddLivenessBackground but with the ability to cancel
		// Note: if you don't need to cancel execution (because this runs forever), use AddLivenessBackground()
		// name - name of the health check
		// interval - how long the report should be considered valid (to prevent ddosing health checks)
		// check - the function to be checked
		AddLivenessBackgroundWithContext(ctx context.Context, name string, duration time.Duration, check Check)

		// AddReadiness adds a check that indicates that this instance of the
		// application is currently unable to serve requests because of an upstream
		// or some transient failure. If a readiness check fails, this instance
		// should no longer receiver requests, but should not be restarted or
		// destroyed. Executed lazily (when report requested)
		// name - name of the health check
		// interval - how long the report should be considered valid (to prevent ddosing health checks)
		// check - the function to be checked
		AddReadiness(name string, interval time.Duration, check Check)

		// AddReadinessBackground same as AddReadiness but executes on an interval
		// ctx: Context for cancellation
		// name: Name of the health check
		// duration: How often the check should be run
		// If your check takes longer than the interval to execute, the next execution will happen immediately.
		// Note: if you don't need to cancel execution (because this runs forever), use AddReadiness()
		// name - name of the health check
		// interval - how long the report should be considered valid (to prevent ddosing health checks)
		// check - the function to be checked
		AddReadinessBackground(name string, interval time.Duration, check Check)

		// AddReadinessBackgroundWithContext same as AddReadinessBackground but with the ability to cancel
		// Note: if you don't need to cancel execution (because this runs forever), use AddReadinessBackground()
		// name - name of the health check
		// interval - how long the report should be considered valid (to prevent ddosing health checks)
		// check - the function to be checked
		AddReadinessBackgroundWithContext(ctx context.Context, name string, interval time.Duration, check Check)

		// Report returns a roll up report of both liveness and readiness results
		Report() (Report, error)

		// Liveness returns the results of the liveness health check reports
		// when not defined as background runs the health checks before returning
		Liveness() ([]Result, error)

		// Readiness returns the results of the readiness health check reports
		// when not defined as background runs the health checks before returning
		Readiness() ([]Result, error)
	}

	// Handler interface for convience of wiring up to existing systems
	Handler interface {
		// Live is the HTTP handler for just the /live endpoint, which is
		// useful if you need to attach it into your own HTTP handler tree.
		Live(http.ResponseWriter, *http.Request)

		// Ready is the HTTP handler for just the /ready endpoint, which is
		// useful if you need to attach it into your own HTTP handler tree.
		Ready(http.ResponseWriter, *http.Request)

		// Report is the HTTP handler for just the /health endpoint, which is
		// useful if you need to attach it into your own HTTP handler tree.
		Report(http.ResponseWriter, *http.Request)
	}

	// Report is a structure returned from endpoints for more detailed reporting
	// typically /health
	// ReportAsOf: When the report was generated
	// Duration: How long all the reports took SUM(Results.duration)
	// Interval: How long the report is considered fresh and should not be re-run
	// Liveness: Results for liveness checks
	// Report: Results for readiness checks
	Report struct {
		ReportAsOf *time.Time `json:"report_as_of_utc"`
		Duration   uint32     `json:"duration_ms"`
		Liveness   []Result   `json:"liveness"`
		Readiness  []Result   `json:"readiness"`
	}

	// Result is the result of the health check test
	// Duration: How long the check took in milliseconds
	// Interval: For background checks, how often the check should be run
	// Name: The name of the test
	// Status: status of the health check. Anything but "OK" considered an error
	// TestedAt: When the test was run
	Result struct {
		Duration uint32     `json:"duration_ms"`
		Interval uint32     `json:"interval_ms,omitempty"`
		Name     string     `json:"name"`
		Status   string     `json:"status"`
		TestedAt *time.Time `json:"tested_at_utc"`
	}
)
