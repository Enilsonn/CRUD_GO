package tools

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func SendJson(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error(
			"Failed to marshal json data",
			"error", err,
		)
		SendJson(
			w,
			Response{Error: "Failed to marshal json data"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error(
			"Error to write response to client",
			"error", err,
		)
		return
	}
}
