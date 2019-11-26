package infra

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HandleGorilla(r *mux.Router) {
	r.HandleFunc("/health/gtg", goodToGo)
	r.HandleFunc("/health/asg", serviceCanary)
	r.HandleFunc("/health", health)
	r.HandleFunc("/metadata", metadata)
	r.HandleFunc("/debug/environment", debugEnvironmentName)
	r.HandleFunc("/debug/headers", debugHeaders)
	r.HandleFunc("/debug/time", debugTime)
	r.HandleFunc("/debug/error", debugError)
	r.HandleFunc("/debug/env", debugEnvironment)
	r.HandleFunc("/debug/filedates", debugFileDates)
	r.HandleFunc("/debug/name", debugName)
}

func Handle() {
	http.HandleFunc("/health/gtg", goodToGo)
	http.HandleFunc("/health/asg", serviceCanary)
	http.HandleFunc("/health", health)
	http.HandleFunc("/metadata", metadata)
	http.HandleFunc("/debug/environment", debugEnvironmentName)
	http.HandleFunc("/debug/headers", debugHeaders)
	http.HandleFunc("/debug/time", debugTime)
	http.HandleFunc("/debug/error", debugError)
	http.HandleFunc("/debug/env", debugEnvironment)
	http.HandleFunc("/debug/filedates", debugFileDates)
	http.HandleFunc("/debug/name", debugName)
}
