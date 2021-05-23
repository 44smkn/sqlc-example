package usecase

import (
	"database/sql"

	"github.com/44smkn/sqlc-sample/pkg/config"
	"github.com/44smkn/sqlc-sample/pkg/infrastructure/persistence"

	"github.com/44smkn/sqlc-sample/pkg/domain"
	"go.uber.org/zap"
)

var chairRepository domain.ChairRepository

var chairSearchCondition config.ChairSearchCondition

func InitializeRepository(conn *sql.DB, logger *zap.Logger) {
	chairRepository = persistence.NewChairRepository(conn, logger)
}

func IniticalizeSearchCondition(csc config.ChairSearchCondition) {
	chairSearchCondition = csc
}
