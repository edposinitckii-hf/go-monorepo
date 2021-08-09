package main

import (
	"context"

	"github.com/edposinitckii-hf/monorepo/services/service-one/internal"
	"github.com/edposinitckii-hf/monorepo/services/service-two/pkg"
	"github.com/edposinitckii-hf/monorepo/toolkit/log"
	"github.com/spf13/cobra"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, _ := log.NewLogger()
	svc := internal.NewService(pkg.NewSvcClient(), logger)

	root := &cobra.Command{
		Use: "service-onw",
		RunE: func(cmd *cobra.Command, args []string) error {
			return svc.Execute(ctx)
		},
	}

	_ = root.ExecuteContext(ctx)
}
