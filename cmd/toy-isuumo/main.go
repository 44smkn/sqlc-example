package main

import (
	"fmt"
	"os"

	"github.com/44smkn/sqlc-sample/pkg/config"
	"github.com/44smkn/sqlc-sample/pkg/log"
	"github.com/44smkn/sqlc-sample/pkg/usecase"
	"go.uber.org/zap"
)

const (
	ExitCodeOK int = 0

	// Errors start at 10
	ExitCodeError = 10 + iota
	ExitCodeParseEnvironmentVariables
	ExitCodeCreateLoggerError
	ExitDBConnectError
	ExitCodeRunServerError
)

func main() {
	os.Exit(run())
}

func run() int {
	cfg, err := config.ReadFromEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to read environment variables: %s\n", err.Error())
		return ExitCodeParseEnvironmentVariables
	}

	logger, err := log.NewLogger(cfg.LogLevel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to create logger: %s\n", err.Error())
		return ExitCodeCreateLoggerError
	}

	conn, err := cfg.DBConfig.GetDB()
	if err != nil {
		logger.Error("Failed to connect DB", zap.Error(err))
		return ExitDBConnectError
	}
	usecase.InitializeRepository(conn, logger)

	srv := NewServer(cfg.Port, logger)
	err = srv.Run()
	if err != nil {
		logger.Error("Failed to run server", zap.Error(err))
		return ExitCodeRunServerError
	}

	return ExitCodeOK
}
