package utils

import (
	"encoding/json"
	"net/http"
)

func RespMsg(message string, status int, w http.ResponseWriter) {
	resp, err := json.Marshal(map[string]any{
		"error": message,
	})
	if err != nil {
		RespErr(err, w)
		return
	}

	w.WriteHeader(status)
	w.Write(resp)
}
