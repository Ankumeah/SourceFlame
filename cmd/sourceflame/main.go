package main

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/Ankumeah/SourceFlame/internal/apis"
	"github.com/Ankumeah/SourceFlame/internal/app"
)

var ctx = context.Background()
var a = &app.App{}

func main() {
	loadEnv(a)

	var wg sync.WaitGroup
	wg.Go(func() { getDB(ctx, a) })
	wg.Go(func() { getHasher(ctx, a) })
	wg.Wait()

	if err := http.ListenAndServe(
		a.Config.Host,
		apis.Apis(a),
	); err != nil {
		log.Fatalf("main.http.ListenAndServe: %v\n", err)
	}
}
