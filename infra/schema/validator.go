package schema

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

type validator struct {
	loader   *gojsonschema.SchemaLoader
	cache    map[string]*gojsonschema.Schema
	registry Registry
}

func NewValidator(reg Registry) (Validator, error) {
	sv := &validator{
		loader:   gojsonschema.NewSchemaLoader(),
		cache:    make(map[string]*gojsonschema.Schema),
		registry: reg,
	}
	for _, v := range reg.Registered() {
		if err := sv.loader.AddSchemas(gojsonschema.NewRawLoader(v)); err != nil {
			return sv, err
		}
	}
	return sv, nil
}

func (sv *validator) Validate(schemaURI string, json []byte) (error, ErrorCollection) {
	return sv.validate(schemaURI, gojsonschema.NewBytesLoader(json))
}

func (sv *validator) ValidateString(schemaURI string, json string) (error, ErrorCollection) {
	return sv.validate(schemaURI, gojsonschema.NewStringLoader(json))
}

func (sv *validator) ValidateGoType(schemaURI string, obj interface{}) (error, ErrorCollection) {
	return sv.validate(schemaURI, gojsonschema.NewGoLoader(obj))
}

func (sv *validator) Consumes(schemaURI string, next http.HandlerFunc) http.HandlerFunc {
	return Consumes(sv, schemaURI, next)
}

func (sv *validator) Produces(schemaURI string, next http.HandlerFunc) http.HandlerFunc {
	return Produces(schemaURI, next)
}

func (sv *validator) validate(schemaURI string, loader gojsonschema.JSONLoader) (error, ErrorCollection) {
	schema, err := sv.compile(schemaURI)
	if err != nil {
		return err, nil
	}
	res, err := schema.Validate(loader)
	if err != nil {
		return err, nil
	}
	if !res.Valid() {
		return nil, cleanErrors(res)
	}
	return nil, nil
}

func (sv *validator) compile(schemaURI string) (*gojsonschema.Schema, error) {
	if s, ok := sv.cache[schemaURI]; ok {
		return s, nil
	}
	s, err := sv.loader.Compile(gojsonschema.NewReferenceLoader(schemaURI))
	if err != nil {
		return nil, fmt.Errorf("loading/compiling validating schema: %v", err)
	}
	sv.cache[schemaURI] = s
	return s, nil
}

func cleanErrors(res *gojsonschema.Result) ErrorCollection {
	reserrors := res.Errors()
	errors := make([]error, len(reserrors))

	for i, e := range reserrors {
		desc := e.String()
		if strings.HasPrefix(desc, "(root): ") {
			desc = strings.Replace(desc, "(root): ", "", 1)
		}
		if strings.Contains(desc, ": Does not match") {
			if e.Value() != nil {
				desc = fmt.Sprintf("%s. value was: %v", desc, e.Value())
			}
		}
		errors[i] = fmt.Errorf("%s", strings.Trim(desc, " "))
	}
	return errors
}
