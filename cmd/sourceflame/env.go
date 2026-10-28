package main

import (
	"log"
	"os"

	"github.com/Ankumeah/SourceFlame/internal/app"
)

// Temprory map to store type unsafe env vars
var envVars = map[string]string{
	"API_VERSION": "",
	"HOST":        "",

	"DB_URL":      "",
	"REPOS_PATH":  "",
	"HASH_CONFIG": "",
}

// Loads all needed env vars into passsed config.
// Exits program on first unset env var
func loadEnv(app *app.App) {
	log.Println("Loading env")

	for env := range envVars {
		_env, ok := os.LookupEnv(env)
		if !ok {
			log.Fatalf("Unset env var: %v\n", env)
			os.Exit(1)
		}
		envVars[env] = _env
	}
	setSettings(app)

	log.Println("Loaded env")
}

// Loads the type unsafe map into the type safe config
func setSettings(app *app.App) {
	app.Config.ApiVersion = envVars["API_VERSION"]
	app.Config.Host = envVars["HOST"]

	app.Config.DBURL = envVars["DB_URL"]
	app.Config.ReposPath = envVars["REPOS_PATH"]
	app.Config.HashConfig = envVars["HASH_CONFIG"]
}
