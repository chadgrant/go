package infra

import (
	"net/http"
	"os"
	"runtime"
	"time"
)

//will be filled in by ldflags at build time
var (
	Application         string
	ApplicationFriendly string
	BuildNumber         string
	BuiltBy             string
	BuiltWhen           string
	GitSha1             string
	GitBranch           string
	CompilerVersion     string
	GroupID             string
	GoVersion           string

	startTime = time.Now().UTC()
	md        = make(map[string]string, 0)
)

//Cannot get osVersion, osAvgLoad or osArch from GO
//arch is an "educated" guess/lie
type MetadataResponse struct {
	Application         string    `json:"application,omitempty"`
	ApplicationFriendly string    `json:"applicationFriendly,omitempty"`
	BuildNumber         string    `json:"buildNumber,omitempty"`
	BuiltBy             string    `json:"builtBy,omitempty"`
	BuiltWhen           string    `json:"builtWhen,omitempty"`
	GitSha1             string    `json:"gitSha1,omitempty"`
	GitBranch           string    `json:"gitBranch,omitempty"`
	CompilerVersion     string    `json:"compilerVersion,omitempty"`
	MachineName         string    `json:"machineName,omitempty"`
	UpSince             time.Time `json:"upSince,omitempty"`
	CurrentTime         time.Time `json:"currentTime,omitempty"`
	GroupID             string    `json:"groupId,omitempty"`
	OsArch              string    `json:"osArch,omitempty"`
	OsName              string    `json:"osName,omitempty"`
	OsNumProcessors     int       `json:"osNumProcessors,omitempty"`
	GoVersion           string    `json:"goVersion,omitempty"`
	Version             int       `json:"version,omitempty"`
}

func parseDate(s string) string {
	t, err := time.Parse(time.UnixDate, s)
	if err != nil {
		return err.Error()
	}
	return t.Format(time.RFC3339)
}

func Metadata(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	m := &MetadataResponse{
		Application:         Application,
		ApplicationFriendly: ApplicationFriendly,
		BuildNumber:         BuildNumber,
		BuiltBy:             BuiltBy,
		BuiltWhen:           BuiltWhen,
		GitSha1:             GitSha1,
		GitBranch:           GitBranch,
		CompilerVersion:     CompilerVersion,
		GroupID:             GroupID,
		MachineName:         hostname,
		UpSince:             startTime,
		CurrentTime:         time.Now().UTC(),
		OsArch:              runtime.GOARCH,
		OsName:              runtime.GOOS,
		OsNumProcessors:     runtime.NumCPU(),
		GoVersion:           runtime.Version(),
		Version:             1,
	}
	jsonResponse(w, r, m)
}
