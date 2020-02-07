package schema

import (
	"testing"
)

func TestLoadFromDisk(t *testing.T) {
	root := "./test-schemas"
	reg := NewRegistry()
	err := reg.AddDirectory(root)
	if err != nil {
		t.Errorf("loading schemas from disk : %s %v", root, err)
	}
}

func TestLoadFromURI(t *testing.T) {
	reg := NewRegistry()
	err := reg.Add("http://schemas.vevo.com/models/content/3.0/artist.json")
	if err != nil {
		t.Error(err)
	}
}

func TestAllSchemasLoaded(t *testing.T) {
	root := "./test-schemas"
	reg := NewRegistry()
	err := reg.AddDirectory(root)
	if err != nil {
		t.Errorf("loading schemas from disk : %s %v", root, err)
	}
	m := reg.Registered()
	if len(m) != 3 {
		t.Errorf("got %d schemas, expected %d", len(m), 3)
	}
}
