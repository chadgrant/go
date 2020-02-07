package schema

import "net/http"

type (
	ErrorCollection []error

	Registry interface {
		Add(schemaURI string) error
		AddDirectory(dir string) error
		AddBytes(schema []byte) error
		Registered() map[string]map[string]interface{}
	}

	Validator interface {
		Validate(schemaURI string, json []byte) (error, ErrorCollection)
		ValidateString(schemaURI string, json string) (error, ErrorCollection)
		ValidateGoType(schemaURI string, obj interface{}) (error, ErrorCollection)

		Consumes(schemaURI string, next http.HandlerFunc) http.HandlerFunc
		Produces(schemaURI string, next http.HandlerFunc) http.HandlerFunc
	}

	Handler interface {
		List(http.ResponseWriter, *http.Request)
		Get(http.ResponseWriter, *http.Request)
	}
)
