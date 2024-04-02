package main

import (
	"github.com/rs/zerolog/log"
	"helloWorld/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal().Err(err)
	}

	if err := a.Run(); err != nil {
		log.Fatal().Err(err)
	}
}
