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
	http.HandleFunc("/debug/name", DebugName)
}

func HandleMux(mux *http.ServeMux) {
	mux.HandleFunc("/health/liveness", Liveness)
	mux.HandleFunc("/health/readiness", Readiness)
	mux.HandleFunc("/health", Health)
	mux.HandleFunc("/metadata", Metadata)
	mux.HandleFunc("/debug/environment", DebugEnvironmentName)
	mux.HandleFunc("/debug/headers", DebugHeaders)
	mux.HandleFunc("/debug/time", DebugTime)
	mux.HandleFunc("/debug/error", DebugError)
	mux.HandleFunc("/debug/name", DebugName)
}