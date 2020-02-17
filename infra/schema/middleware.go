package schema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Produces adds a X-Schema header for the client to know where
// to retrieve the schema that can be used to validate the response
func Produces(schemaURI string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Schema", schemaURI)
		next.ServeHTTP(w, r)
	})
}

// Consumes validates the incoming body against a schema before calling the next chain in the pipeline
func Consumes(validator Validator, schemaURI string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			serveError(w, err, "reading body to validate json: %v")
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		err, errors := validator.Validate(schemaURI, body)
		if err != nil {
			serveError(w, err, "validating json against schema: %v")
			return
		}

		if len(errors) > 0 {
			writeErrors(w, errors)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func writeErrors(w http.ResponseWriter, valErrors ErrorCollection) {
	resp := struct {
		Message string
		Errors  []string
	}{
		Message: "Validation Errors",
		Errors:  make([]string, len(valErrors)),
	}

	for i, e := range valErrors {
		resp.Errors[i] = e.Error()
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("unable to encode response for schema validation errors: %v", err)
	}
}

func serveError(w http.ResponseWriter, err error, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(fmt.Sprintf(msg, err)))
}
