module github.com/edposinitckii-hf/monorepo/services

go 1.16

require (
	github.com/edposinitckii-hf/monorepo/toolkit v1.0.0
	github.com/jmoiron/sqlx v1.3.0
	github.com/spf13/cobra v1.2.1
	go.uber.org/zap v1.18.0
)

replace github.com/edposinitckii-hf/monorepo/toolkit v1.0.0 => ../toolkit
