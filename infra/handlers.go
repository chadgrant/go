package infra

import (
	"github.com/chadgrant/go-http-infra/infra/health"
	"github.com/chadgrant/go-http-infra/infra/metadata"
	"net/http"
)

func RegisterInfraHandlers(register func(string, func(http.ResponseWriter, *http.Request))) (hc health.HealthChecker, md metadata.Handler) {
	md = metadata.NewHandler()
	hc = health.NewHealthChecker()
	hh := health.NewHandler(hc)
	register("/live", hh.Live)
	register("/ready", hh.Ready)
	register("/health", hh.Report)
	register("/metadata", md.Metadata)
	register("/debug/environment", DebugEnvironmentName)
	register("/debug/headers", DebugHeaders)
	register("/debug/time", DebugTime)
	register("/debug/error", DebugError)
	register("/debug/name", DebugName)
	return
}
