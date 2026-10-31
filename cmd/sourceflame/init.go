package main

import (
	"context"
	"log"

	"github.com/Ankumeah/SourceFlame/internal/app"
	"github.com/Ankumeah/SourceFlame/internal/db"
	"github.com/Ankumeah/SourceFlame/internal/db/migrations"
	"github.com/Ankumeah/SourceFlame/internal/git"
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

func migrateDB(ctx context.Context, app *app.App) {
	log.Println("Starting DB migrations")

	before, after, err := migrations.ApplyMigrations(ctx, app.DB.DB())
	if err != nil {
		log.Fatalf("Error while migrating db: %v\n", err)
	}

	log.Printf("DB migrations complete (%v -> %v)\n", before, after)
}

func getHasher(app *app.App) {
	log.Println("Getting hasher")

	var err error
	app.Hasher, err = hash.GetHasher(app.Config.HashConfig)
	if err != nil {
		log.Fatalf("Error while getting hasher: %v\n", err)
	}

	log.Println("Got hasher")
}

func getGitController(app *app.App) {
	log.Println("Getting git controller")

	app.GitController = git.GetGitController(app.Config.ReposPath)

	log.Println("Got git controller")
}
