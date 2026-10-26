package main

import (
	"log"
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/apis"
	"github.com/Ankumeah/SourceFlame/internal/app"
)

func main() {
	a := &app.App{
		Config: app.Config{},
	}

	if err := http.ListenAndServe(
		"0.0.0.0:8080",
		apis.Apis(a),
	); err != nil {
		log.Fatalf("main.http.ListenAndServe: %v\n", err)
	}
}
