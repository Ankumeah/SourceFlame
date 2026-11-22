package utils

import (
	"log"
	"net/http"
)

func RespErr(err error, w http.ResponseWriter) {
	log.Printf("Error: %v\n", err)
	w.WriteHeader(http.StatusInternalServerError)
}
