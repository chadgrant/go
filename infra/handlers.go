package infra

import (
	"fmt"
	"net/http"

	"github.com/chadgrant/go-http-infra/infra/health"
	"github.com/chadgrant/go-http-infra/infra/metadata"
	"github.com/chadgrant/go-http-infra/infra/schema"
)

func RegisterInfraHandlers(register func(string, http.HandlerFunc), hc health.HealthChecker, sr schema.Registry) error {

	if sr == nil {
		sr = schema.NewRegistry()
	}
	if hc == nil {
		hc = health.NewHealthChecker()
	}
	if err := addSchemas(sr, health.Asset, metadata.Asset, schema.Asset); err != nil {
		return fmt.Errorf("registering schams %v", err)
	}
	sv, err := schema.NewValidator(sr)
	if err != nil {
		return fmt.Errorf("creating schema validator %v", err)
	}

	hh := health.NewHandler(hc)
	sh := schema.NewHandler(sr)
	register("/live", hh.Live)
	register("/ready", hh.Ready)
	register("/health", sv.Produces("http://schemas.sentex.io/service/health.json", hh.Report))
	register("/metadata", sv.Produces("http://schemas.sentex.io/service/metadata.json", metadata.NewHandler().Metadata))
	register("/schemas", sv.Produces("http://schemas.sentex.io/service/schemalist.json", sh.List))
	register("/schema", sh.Get)
	register("/debug/environment", DebugEnvironmentName)
	register("/debug/headers", DebugHeaders)
	register("/debug/time", DebugTime)
	register("/debug/error", DebugError)
	register("/debug/name", DebugName)

	return nil
}

func addSchemas(sr schema.Registry, getters ...func(string) ([]byte, error)) error {
	for _, get := range getters {
		bs, err := get("schema.json")
		if err != nil {
			return err
		}
		if err := sr.AddBytes(bs); err != nil {
			return err
		}
	}
	return nil
}
