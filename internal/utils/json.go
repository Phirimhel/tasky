package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func ResponseWithJSON(w http.ResponseWriter, data any, status int) error {
	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	_, err := w.Write(buf.Bytes())

	return err
}

type ErrorDTO struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func ResponseWithError(w http.ResponseWriter, message string, status int) {
	_ = ResponseWithJSON(
		w,
		ErrorDTO{
			Message: message,
			Time:    time.Now(),
		},
		status,
	)
}
