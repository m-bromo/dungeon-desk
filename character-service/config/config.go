package config

import (
	"log"
	"log/slog"

	"github.com/Netflix/go-env"
)

var Env Environment

func SetupConfig() {
	_, err := env.UnmarshalFromEnviron(&Env)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Environment loaded")
}
