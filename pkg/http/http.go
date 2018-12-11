package http

import (
	"encoding/json"
	"net/http"

	"github.com/tbaud0n/sample-api-architecture/pkg/logger"
)

// Error output the error to the HTTP
func Error(w http.ResponseWriter, err error) {
	http.Error(w, "Internal error", http.StatusInternalServerError)
}

// NotFound handler
func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotFound)
}

// SetJSONResponse writes the data as JSON as response of the http request
func SetJSONResponse(w http.ResponseWriter, data interface{}) (err error) {
	var b []byte
	if b, err = json.Marshal(data); err != nil {
		Error(w, logger.LogError(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(b); err != nil {
		logger.LogError(err)
	}

	return
}
