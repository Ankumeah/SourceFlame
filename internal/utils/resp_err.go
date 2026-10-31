package utils

import (
	"log"
	"net/http"
)

func RespErr(err error, w http.ResponseWriter) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
}
