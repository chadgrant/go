package infra

import (
	"net/http"

	"github.com/chadgrant/go-http-infra/infra/health"
	"github.com/chadgrant/go-http-infra/infra/metadata"
	"github.com/chadgrant/go-http-infra/infra/schema"
)

func RegisterInfraHandlers(register func(string, func(http.ResponseWriter, *http.Request))) (hc health.HealthChecker, md metadata.Handler, sr schema.Registry, err error) {
	md = metadata.NewHandler()
	hc = health.NewHealthChecker()
	hh := health.NewHandler(hc)
	sr, err = addSchemas(health.Asset, metadata.Asset)
	sh := schema.NewHandler(sr)
	register("/live", hh.Live)
	register("/ready", hh.Ready)
	register("/health", hh.Report)
	register("/metadata", md.Metadata)
	register("/schemas", sh.List)
	register("/schema", sh.Get)
	register("/debug/environment", DebugEnvironmentName)
	register("/debug/headers", DebugHeaders)
	register("/debug/time", DebugTime)
	register("/debug/error", DebugError)
	register("/debug/name", DebugName)
	return
}

func addSchemas(getters ...func(string) ([]byte, error)) (schema.Registry, error) {
	sr := schema.NewRegistry()
	for _, get := range getters {
		bs, err := get("schema.json")
		if err != nil {
			return sr, err
		}
		if err := sr.AddBytes(bs); err != nil {
			return sr, err
		}
	}
	return sr, nil
}
