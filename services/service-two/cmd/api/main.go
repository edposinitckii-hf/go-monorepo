package main

import (
	"context"

	"github.com/edposinitckii-hf/monorepo/services/service-two/internal"
	"github.com/edposinitckii-hf/monorepo/toolkit/db/postgres"
	"github.com/edposinitckii-hf/monorepo/toolkit/log"
	"github.com/spf13/cobra"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := internal.NewConfig()
	logger, _ := log.NewLogger()
	db, _ := postgres.NewPostgresConnection(config.DSN)

	app := internal.NewApp(db, logger)

	root := cobra.Command{
		Use: "service-two-api",
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Execute(cmd.Context())
		},
	}

	_ = root.ExecuteContext(ctx)
}
