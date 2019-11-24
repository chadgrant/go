package infra

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type errorResponse struct {
	Errors []apiError `json:"errors"`
}

type apiError struct {
	Source  string `json:"source"`
	Message string `json:"message"`
}

func (e apiError) Error() string {
	return fmt.Sprintf("Source:%s\n Message:%s\n", e.Source, e.Message)
}

func Error(w http.ResponseWriter, r *http.Request, status int, errors ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errs := make([]apiError, 0)
	for _, v := range errors {
		switch e := v.(type) {
		case *apiError:
			errs = append(errs, *e)
			break
		case apiError:
			errs = append(errs, e)
			break
		case string:
			errs = append(errs, apiError{Source: "service", Message: e})
			break
		default:
			errs = append(errs, apiError{Source: "httpError", Message: fmt.Sprintf("Could not add error type %T", v)})
		}
	}

	for _, e := range errs {
		log.Println(e)
	}

	b, err := json.Marshal(&errorResponse{errs})
	if err == nil {
		w.Write(b)
	}
}