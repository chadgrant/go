package metadata

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	Vendor = "test-vendor"
	Group = "test-group"
	Service = "test-service"
	Friendly = "test-friendly"
	Description = "test-description"
	Repo = "test-repo"
	BuildNumber = "1.0.1"
	BuildTime = "2020-01-22T22:57:49Z"
	GitHash = "blahledah"
	GitBranch = "master"
	CompilerVersion = "go/linux v1.13"

	h := NewHandler()

	r, _ := http.NewRequest(http.MethodGet, "metadata", nil)
	w := httptest.NewRecorder()

	h.Metadata(w, r)

	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code %d", w.Result().StatusCode)
	}
	js, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("handler response failed: %v", err)
	}

	resp := &Response{}
	if err := json.Unmarshal(js, resp); err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	property(t, "Vendor", Vendor, resp.Vendor)
	property(t, "Group", Group, resp.Group)
	property(t, "Service", Service, resp.Service)
	property(t, "Friendly", Friendly, resp.Friendly)
	property(t, "Description", Description, resp.Description)
	property(t, "Repo", Repo, resp.Repo)
	property(t, "BuildNumber", BuildNumber, resp.BuildNumber)
	property(t, "BuiltBy", BuiltBy, resp.BuiltBy)
	property(t, "BuildTime", BuildTime, resp.BuildTime)
	property(t, "GitHash", GitHash, resp.GitHash)
	property(t, "GitBranch", GitBranch, resp.GitBranch)
	property(t, "CompilerVersion", CompilerVersion, resp.CompilerVersion)
}

func property(t *testing.T, prop, expect, got string) {
	if expect != got {
		t.Errorf("%s: expected %s got: %s", prop, expect, got)
	}
}
