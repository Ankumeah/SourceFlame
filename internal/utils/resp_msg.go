package utils

import (
	"encoding/json"
	"net/http"
)

func RespMsg(message string, status int, errs []map[string]any, w http.ResponseWriter) {
	resp, err := json.Marshal(map[string]any{
		"error": message,
	})
	if err != nil {
		RespErr(err, errs, w)
		return
	}

	w.WriteHeader(status)
	w.Write(resp)
}
