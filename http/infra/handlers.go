package infra

import "net/http"

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