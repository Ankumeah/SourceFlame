package main

import (
	"log"
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/apis"
	"github.com/Ankumeah/SourceFlame/internal/frontend"
	"github.com/Ankumeah/SourceFlame/internal/middlewares"
)

func main() {
	app := setupDI()

	mux := http.NewServeMux()

	frontend.Handle(app, mux)
	apis.Handle(app, mux)

	// lOoK aT ThE LiSP CoDer hERe,
	// sybau bro, we have 2 middlewares my guy,
	// we cross that bridge when we need to
	if err := http.ListenAndServe(
		app.Config.Host,
		middlewares.Log(
			middlewares.Cors(
				mux,
			),
		),
	); err != nil {
		log.Fatalf("Error while serving http: %v\n", err)
	}
}
