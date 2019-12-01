package gorilla

import (
	"github.com/chadgrant/go/http/infra"
	"github.com/gorilla/mux"
)

func HandleGorilla(r *mux.Router) {
	r.HandleFunc("/health/liveness", infra.Liveness)
	r.HandleFunc("/health/readiness", infra.Readiness)
	r.HandleFunc("/health", infra.Health)
	r.HandleFunc("/metadata", infra.Metadata)
	r.HandleFunc("/debug/environment", infra.DebugEnvironmentName)
	r.HandleFunc("/debug/headers", infra.DebugHeaders)
	r.HandleFunc("/debug/time", infra.DebugTime)
	r.HandleFunc("/debug/error", infra.DebugError)
	r.HandleFunc("/debug/env", infra.DebugEnvironment)
	r.HandleFunc("/debug/filedates", infra.DebugFileDates)
	r.HandleFunc("/debug/name", infra.DebugName)
}
