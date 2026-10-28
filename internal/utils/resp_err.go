package utils

import (
	"net/http"
)

func RespErr(err error, errs []map[string]any, w http.ResponseWriter) {
	errs = append(errs, map[string]any{
		"error": err,
	})
	w.WriteHeader(http.StatusInternalServerError)
}
