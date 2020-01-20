package gorilla

import (
	"github.com/chadgrant/go-http-infra/infra"
	"github.com/chadgrant/go-http-infra/infra/health"
	"github.com/chadgrant/go-http-infra/infra/metadata"
	"github.com/gorilla/mux"
)

func Handle(r *mux.Router, hc health.Handler) {
	md := metadata.NewHandler()
	r.HandleFunc("/live", hc.Live)
	r.HandleFunc("/ready", hc.Ready)
	r.HandleFunc("/health", hc.Report)
	r.HandleFunc("/metadata", md.Metadata)
	r.HandleFunc("/debug/environment", infra.DebugEnvironmentName)
	r.HandleFunc("/debug/headers", infra.DebugHeaders)
	r.HandleFunc("/debug/time", infra.DebugTime)
	r.HandleFunc("/debug/error", infra.DebugError)
	r.HandleFunc("/debug/name", infra.DebugName)
}
