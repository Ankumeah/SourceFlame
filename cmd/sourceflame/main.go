package main

import (
	"log"

	"github.com/Ankumeah/SourceFlame/internal/apis"
)

func main() {
	app := setupDI()

	if err := apis.ServeHTTP(app); err != nil {
		log.Fatalf("Error while serving http: %v\n", err)
	}
}
