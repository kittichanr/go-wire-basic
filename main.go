package main

import (
	"context"
	"log"

	"github.com/kittichanr/go-wire/di"
)

func main() {
	ctx := context.Background()
	_, cleanUpFn, err := di.InitializeApplication(ctx)
	defer cleanUpFn()

	if err != nil {
		log.Fatal("initial app failed")
		panic(err)
	}
}
