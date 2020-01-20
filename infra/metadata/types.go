package metadata

import (
	"net/http"
)

// will be filled in by ldflags at build time
var (
	Vendor          string
	Group           string
	Service         string
	Friendly        string
	Description     string
	Repo            string
	BuildNumber     string
	BuiltBy         string
	BuildTime       string
	GitHash         string
	GitBranch       string
	CompilerVersion string
	Version         string
)

type (
	// Response is the metadata response payload
	// Vendor: Vendor of the service i.e. Google
	// Group: Group responsible for the service. i.e. "Growth"
	// Service: Short name of the service i.e. "member_api"
	// Friendly: Human friendly name for the service i.e. "Member API"
	// Description: Optional description of the service
	// Repo: Optional Repo for the servicve
	// BuildNumber: build number of the service i.e. 1.0.0
	// BuiltBy: user that triggered the build
	// BuildTime: date/time of build
	// Githash: git hash of the build
	// GitBranch: branch the build was built out of
	// CompilerVersion: what version of the go compiler used in the build
	// Version: version of this schema
	Response struct {
		Vendor          string `json:"vendor,omitempty"`
		Group           string `json:"group,omitempty"`
		Service         string `json:"service,omitempty"`
		Friendly        string `json:"friendly,omitempty"`
		Description     string `json:"description,omitempty"`
		Repo            string `json:"build_repo,omitempty"`
		BuildNumber     string `json:"build_number,omitempty"`
		BuiltBy         string `json:"built_by,omitempty"`
		BuildTime       string `json:"build_time,omitempty"`
		GitHash         string `json:"git_hash,omitempty"`
		GitBranch       string `json:"git_branch,omitempty"`
		CompilerVersion string `json:"compiler_version,omitempty"`
		Version         int    `json:"version,omitempty"`
	}

	// Handler interface for convience of wiring up to existing systems
	Handler interface {
		// Metadata is the HTTP handler for the /metadata endpoint, which is
		// useful if you need to attach it into your own HTTP handlers.
		Metadata(http.ResponseWriter, *http.Request)
	}
)
