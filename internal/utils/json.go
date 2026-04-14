package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func ResponseWithJSON(w http.ResponseWriter, status int, data any) error {
	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	_, err := w.Write(buf.Bytes())

	return err
}

func ResponseWithError(w http.ResponseWriter, status int, msg string) {
	_ = ResponseWithJSON(w, status, map[string]string{
		"error": msg,
	})
}
