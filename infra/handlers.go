package infra

import (
	"github.com/chadgrant/go-http-infra/infra/health"
	"net/http"
)

func Handle(hc health.Handler) {
	http.HandleFunc("/live", hc.Live)
	http.HandleFunc("/ready", hc.Ready)
	http.HandleFunc("/health", hc.Report)
	http.HandleFunc("/metadata", Metadata)
	http.HandleFunc("/debug/environment", DebugEnvironmentName)
	http.HandleFunc("/debug/headers", DebugHeaders)
	http.HandleFunc("/debug/time", DebugTime)
	http.HandleFunc("/debug/error", DebugError)
	http.HandleFunc("/debug/name", DebugName)
}

func HandleMux(hc health.Handler, mux *http.ServeMux) {
	mux.HandleFunc("/live", hc.Live)
	mux.HandleFunc("/ready", hc.Ready)
	mux.HandleFunc("/health", hc.Report)
	mux.HandleFunc("/metadata", Metadata)
	mux.HandleFunc("/debug/environment", DebugEnvironmentName)
	mux.HandleFunc("/debug/headers", DebugHeaders)
	mux.HandleFunc("/debug/time", DebugTime)
	mux.HandleFunc("/debug/error", DebugError)
	mux.HandleFunc("/debug/name", DebugName)
}
