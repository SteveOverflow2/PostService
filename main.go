package main

import (
	"fmt"
	"os"

	"post-service/pkg/config"
	"post-service/pkg/http/rest"
	"post-service/pkg/rabbitmq"
	"post-service/pkg/storage/mysql"
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

	sql, err := mysql.NewMySQLConnection(cfg.MySQL)
	if err != nil {
		return err
	}
	go rabbitmq.StartServer()

	server := rest.NewServer(
		cfg.Version,
		cfg.Environment,
		cfg.HTTP,
		sql,
	)
	server.Init()

	// Runs the new server instance.
	server.Run(cfg.Name)

	return nil
}
