package infra

import (
	"net/http"
)

func Handle() {
	http.HandleFunc("/health/liveness", Liveness)
	http.HandleFunc("/health/readiness", Readiness)
	http.HandleFunc("/health", Health)
	http.HandleFunc("/metadata", Metadata)
	http.HandleFunc("/debug/environment", DebugEnvironmentName)
	http.HandleFunc("/debug/headers", DebugHeaders)
	http.HandleFunc("/debug/time", DebugTime)
	http.HandleFunc("/debug/error", DebugError)
	http.HandleFunc("/debug/env", DebugEnvironment)
	http.HandleFunc("/debug/filedates", DebugFileDates)
	http.HandleFunc("/debug/name", DebugName)
}
