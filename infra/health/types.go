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
		// interval: How often the check should be run
		// If your check takes longer than the interval to execute, the next execution will happen immediately.
		// check - the function to be checked
		AddLivenessBackground(name string, interval time.Duration, check Check)

		// AddLivenessBackgroundWithContext same as AddLivenessBackground but with the ability to cancel
		// Note: if you don't need to cancel execution (because this runs forever), use AddLivenessBackground()
		// ctx: Context for cancellation
		// name - name of the health check
		// interval - How often the check should be run
		// check - the function to be checked
		AddLivenessBackgroundWithContext(ctx context.Context, name string, interval time.Duration, check Check)

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
		// interval: How often the check should be run
		// check - the function to be checked
		// If your check takes longer than the interval to execute, the next execution will happen immediately.
		// Note: if you don't need to cancel execution (because this runs forever), use AddReadiness()
		AddReadinessBackground(name string, interval time.Duration, check Check)

		// AddReadinessBackgroundWithContext same as AddReadinessBackground but with the ability to cancel
		// Note: if you don't need to cancel execution (because this runs forever), use AddReadinessBackground()
		// ctx: Context for cancellation
		// name - name of the health check
		// interval - how often the check should be run
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
	// Liveness: Results for liveness checks
	// Readiness: Results for readiness checks
	// HostName: Name of the current machine serving this request
	// UpSince: timestamp of when the service started
	// OsArch: architecture of the current host
	// OsName: os name of the current host
	// OsNumProcessors: num of procs on the host
	// GoVersion: version of the go runtime
	// Version: version of this schema
	Report struct {
		ReportAsOf      *time.Time `json:"report_as_of_utc"`
		Duration        uint32     `json:"duration_ms"`
		Liveness        []Result   `json:"liveness"`
		Readiness       []Result   `json:"readiness"`
		HostName        string     `json:"host_name,omitempty"`
		UpSince         time.Time  `json:"up_since,omitempty"`
		OsArch          string     `json:"os_arch,omitempty"`
		OsName          string     `json:"os_name,omitempty"`
		OsNumProcessors int        `json:"os_num_processors"`
		OsGoRoutines    int        `json:"os_num_goroutines"`
		GoVersion       string     `json:"go_version,omitempty"`
		Version         int        `json:"version,omitempty"`
		Memory          Memory     `json:"memory"`
	}

	// Memory details of the memory profile
	// https://golang.org/pkg/runtime/#MemStats
	Memory struct {
		Alloc         uint64  `json:"alloc,omitempty"`
		TotalAlloc    uint64  `json:"total_alloc,omitempty"`
		Sys           uint64  `json:"sys,omitempty"`
		Lookups       uint64  `json:"lookups,omitempty"`
		Mallocs       uint64  `json:"mallocs,omitempty"`
		Frees         uint64  `json:"frees,omitempty"`
		HeapAlloc     uint64  `json:"heap_allocs,omitempty"`
		HeapSys       uint64  `json:"heap_sys,omitempty"`
		HeapIdle      uint64  `json:"heap_idle,omitempty"`
		HeapInuse     uint64  `json:"heap_in_use,omitempty"`
		HeapReleased  uint64  `json:"heap_released,omitempty"`
		HeapObjects   uint64  `json:"heap_objects,omitempty"`
		StackInuse    uint64  `json:"stack_in_use,omitempty"`
		StackSys      uint64  `json:"stack_sys,omitempty"`
		MSpanInuse    uint64  `json:"mspan_inuse,omitempty"`
		MSpanSys      uint64  `json:"mspan_sys,omitempty"`
		MCacheInuse   uint64  `json:"mcache_inuse,omitempty"`
		MCacheSys     uint64  `json:"mcache_sys,omitempty"`
		BuckHashSys   uint64  `json:"buck_hash_sys,omitempty"`
		GCSys         uint64  `json:"gc_sys,omitempty"`
		OtherSys      uint64  `json:"other_sys,omitempty"`
		NextGC        uint64  `json:"next_gc,omitempty"`
		LastGC        uint64  `json:"last_gc,omitempty"`
		PauseTotalNs  uint64  `json:"pause_total_n,omitempty"`
		NumGC         uint32  `json:"num_gc,omitempty"`
		NumForcedGC   uint32  `json:"num_forced_gc,omitempty"`
		GCCPUFraction float64 `json:"gcc_cpu_fraction,omitempty"`
		EnableGC      bool    `json:"enable_gc"`
		DebugGC       bool    `json:"debug_gc"`
	}

	// Result is the result of the health check test
	// Name: The name of the test
	// Interval: For background checks, how often the check should be run
	//			 For lazy checks, how long the check is considered fresh
	// Duration: How long the check took in milliseconds
	// Status: status of the health check. Anything but "OK" considered unhealthy
	// TestedAt: When the test was run
	Result struct {
		Name     string     `json:"name"`
		Interval uint32     `json:"interval_ms,omitempty"`
		Duration uint32     `json:"duration_ms"`
		Status   string     `json:"status"`
		TestedAt *time.Time `json:"tested_at_utc"`
	}
)
