package main

import (
	"context"
	"fmt"
	"time"

	"github.com/edposinitckii-hf/monorepo/svc/echo/pkg/http"
	"github.com/edposinitckii-hf/monorepo/svc/monitor/internal"
	"github.com/edposinitckii-hf/monorepo/toolkit/log"
	"github.com/spf13/cobra"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, _ := log.NewLogger()
	svc := internal.NewService(http.NewEchoClient("localhost"), logger)

	root := &cobra.Command{
		Use: "service-onw",
		RunE: func(cmd *cobra.Command, args []string) error {
			echo, err := svc.Execute(ctx, args[0])
			if err != nil {
				return err
			}

			fmt.Printf("Message: %s\nTime: %s\n", echo.Message, echo.Timestamp.Format(time.RFC850))
			return nil
		},
	}

	_ = root.ExecuteContext(ctx)
}
