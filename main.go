package main

import (
	"fmt"
	"os"

	"post-service/pkg/config"
	"post-service/pkg/http/rest"
	"post-service/pkg/util"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg := config.NewConfig()

	err := cfg.LoadConfig()
	if err != nil {
		return util.WrapErrorf(err, util.ErrorCodeInternal, "Environment configuration failed")
	}

	server := rest.NewServer(
		cfg.Version,
		cfg.Environment,
		cfg.HTTP,
	)
	server.Init()

	// Runs the new server instance.
	server.Run(cfg.Name)

	return nil
}
