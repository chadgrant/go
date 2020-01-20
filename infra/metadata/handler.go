package metadata

import (
	"encoding/json"
	"net/http"
	"runtime"
	"github.com/google/martian/log"
)

type handler struct{}

// NewHandler returds Handler interface for convience of wiring up to existing systems
func NewHandler() Handler {
	return &handler{}
}

func (h *handler) Metadata(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	rsp := &Response{
		Vendor:          Vendor,
		Group:           Group,
		Service:         Service,
		Friendly:        Friendly,
		Description:     Description,
		Repo:            Repo,
		BuildNumber:     BuildNumber,
		BuiltBy:         BuiltBy,
		BuildTime:       BuildTime,
		GitHash:         GitHash,
		GitBranch:       GitBranch,
		CompilerVersion: CompilerVersion,
		Version:         1,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	enc := json.NewEncoder(w)
	if err := enc.Encode(rsp); err !=nil {
		log.Errorf("could not encode json response for /metadata %v",err)
	}
}
