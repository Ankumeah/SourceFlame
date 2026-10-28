package main

import (
	"context"
	"log"

	"github.com/Ankumeah/SourceFlame/internal/app"
	"github.com/Ankumeah/SourceFlame/internal/db"
	"github.com/Ankumeah/SourceFlame/internal/hash"
)

func getDB(ctx context.Context, app *app.App) {
	log.Println("Getting DB conn")

	var err error
	app.DB, err = db.GetDBController(app.Config.DBURL)
	if err != nil {
		log.Fatalf("Error while getting db conn: %v\n", err)
	}

	log.Println("Got DB conn")
}

func getHasher(ctx context.Context, app *app.App) {
	log.Println("Getting hasher")

	var err error
	app.Hasher, err = hash.GetHasher(app.Config.HashConfig)
	if err != nil {
		log.Fatalf("Error while getting hasher: %v\n", err)
	}

	log.Println("Got hasher")
}
