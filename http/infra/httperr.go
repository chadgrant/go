package infra

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Errors []APIError `json:"errors"`
}

type APIError struct {
	Source  string `json:"source"`
	Message string `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("Source:%s\n Message:%s\n", e.Source, e.Message)
}

func Error(w http.ResponseWriter, r *http.Request, status int, errors ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errs := make([]APIError, 0)
	for _, v := range errors {
		switch e := v.(type) {
		case *APIError:
			errs = append(errs, *e)
			break
		case APIError:
			errs = append(errs, e)
			break
		case string:
			errs = append(errs, APIError{Source: "service", Message: e})
			break
		default:
			errs = append(errs, APIError{Source: "httpError", Message: fmt.Sprintf("Could not add error type %T", v)})
		}
	}

	for _, e := range errs {
		log.Println(e)
	}

	b, err := json.Marshal(&ErrorResponse{errs})
	if err == nil {
		w.Write(b)
	}
}
