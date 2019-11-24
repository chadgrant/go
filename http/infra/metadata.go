package infra

import (
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

//Cannot get osVersion, osAvgLoad or osArch from GO
//arch is an "educated" guess/lie
type metadataPayload struct {
	BuildNumber     string    `json:"buildNumber,omitempty"`
	BuiltBy         string    `json:"builtBy,omitempty"`
	BuiltWhen       time.Time `json:"builtWhen,omitempty"`
	GitSha1         string    `json:"gitSha1,omitempty"`
	CompilerVersion string    `json:"compilerVersion,omitempty"`
	MachineName     string    `json:"machineName,omitempty"`
	UpSince         time.Time `json:"upSince,omitempty"`
	CurrentTime     time.Time `json:"currentTime,omitempty"`
	GroupID         string    `json:"groupId,omitempty"`
	OsArch          string    `json:"osArch,omitempty"`
	OsName          string    `json:"osName,omitempty"`
	OsNumProcessors int       `json:"osNumProcessors,omitempty"`
	GoVersion       string    `json:"goVersion,omitempty"`
	Version         int       `json:"version,omitempty"`
}

var (
	startTime = time.Now().UTC()
	md        = make(map[string]string, 0)
)

func init() {
	readMetadata()
}

func readMetadata() {
	bytes, err := ioutil.ReadFile("metadata.txt")
	if err == nil {
		lines := strings.Split(string(bytes), "\n")
		for _, line := range lines {
			sp := strings.SplitN(line, "=", 2)
			if len(sp) != 2 {
				continue
			}
			md[strings.ToLower(sp[0])] = sp[1]
		}
	}
}

func getMetadata(key string) string {
	if val, ok := md[key]; ok {
		return val
	}
	return "unknown"
}

func parseDate(s string) time.Time {
	t, err := time.Parse(time.UnixDate, s)
	if err != nil {
		return t
	}
	return time.Now().UTC()
}

func metadata(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	m := &metadataPayload{
		BuildNumber:     getMetadata("build_number"),
		BuiltBy:         getMetadata("build_user"),
		BuiltWhen:       parseDate(getMetadata("build_date")),
		GitSha1:         getMetadata("build_hash"),
		CompilerVersion: getMetadata("build_compiler"),
		GroupID:         getMetadata("build_group"),
		MachineName:     hostname,
		UpSince:         startTime,
		CurrentTime:     time.Now().UTC(),
		OsArch:          runtime.GOARCH,
		OsName:          runtime.GOOS,
		OsNumProcessors: runtime.NumCPU(),
		GoVersion:       runtime.Version(),
		Version:         1,
	}
	jsonResponse(w, r, m)
}