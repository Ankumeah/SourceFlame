package main

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/Ankumeah/SourceFlame/internal/apis"
	"github.com/Ankumeah/SourceFlame/internal/app"
	"github.com/Ankumeah/SourceFlame/internal/frontend"
	"github.com/Ankumeah/SourceFlame/internal/middlewares"
	"github.com/Ankumeah/SourceFlame/internal/smtp"
)

func main() {
	app, ctx := setupDI()

	wg := sync.WaitGroup{}
	wg.Go(func() { startHTTP(app) })
	wg.Go(func() { startSMTP(ctx, app) })
	wg.Wait()
}

func startHTTP(app *app.App) {
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

func startSMTP(ctx context.Context, app *app.App) {
	if err := smtp.Listen(ctx, app); err != nil {
		log.Fatalf("Error while serving smtp: %v\n", err)
	}
}
