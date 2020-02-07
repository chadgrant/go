package schema

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

type (
	registry struct {
		registered map[string]map[string]interface{}
	}
)

func NewRegistry() Registry {
	return &registry{
		registered: make(map[string]map[string]interface{}),
	}
}

func (sr *registry) AddDirectory(dir string) error {
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f.Mode().IsRegular() && strings.HasSuffix(f.Name(), ".json") {
			abs, err := filepath.Abs(path)
			if err != nil {
				return fmt.Errorf("switching path to absolute %s %v", abs, err)
			}
			abs = "file://" + abs
			if err := sr.Add(abs); err != nil {
				return fmt.Errorf("could not add schema %s %v", abs, err)
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("reading schema dir : %v", err)
	}
	return nil
}

func (sr *registry) AddBytes(schema []byte) error {
	json, err := gojsonschema.NewBytesLoader(schema).LoadJSON()
	if err != nil {
		return err
	}
	return sr.addJSON(json, "")
}

func (sr *registry) Add(schemaURI string) error {
	json, err := gojsonschema.NewReferenceLoader(schemaURI).LoadJSON()
	if err != nil {
		return err
	}
	return sr.addJSON(json, schemaURI)
}

func (sr *registry) addJSON(json interface{}, uri string) error {
	if m, ok := json.(map[string]interface{}); ok {
		var id string
		if tid, ok := m["$id"]; ok {
			if sid, sok := tid.(string); sok {
				id = sid
			}
		}
		if len(id) == 0 {
			if tid, ok := m["id"]; ok {
				if sid, sok := tid.(string); sok {
					id = sid
				}
			}
		}
		if len(id) == 0 {
			return fmt.Errorf("did not find an $id for schema %s", uri)
		}
		if len(uri) == 0 {
			uri = id
		}
		sr.registered[id] = m

		return nil
	}
	return fmt.Errorf("type passed was not json / map[string]interface{} was : %T", json)
}

func (sr *registry) Registered() map[string]map[string]interface{} {
	return sr.registered
}
