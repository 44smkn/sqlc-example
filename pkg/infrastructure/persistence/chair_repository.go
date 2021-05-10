package persistence

import (
	"context"
	"database/sql"
	"strings"

	"github.com/44smkn/sqlc-sample/pkg/domain"
	"github.com/44smkn/sqlc-sample/pkg/infrastructure/persistence/mysql"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

type ChairRepository struct {
	queries *mysql.Queries
	logger  *zap.Logger
}

func NewChairRepository(conn *sql.DB, logger *zap.Logger) *ChairRepository {
	queries := mysql.New(conn)
	return &ChairRepository{
		queries: queries,
		logger:  logger,
	}
}

func (c *ChairRepository) Find(ctx context.Context, id *domain.ChairID) (*domain.Chair, error) {
	chair, err := c.queries.GetChair(ctx, id.Value())
	if err != nil {
		return nil, xerrors.Errorf("failed to retrive rows: %w", err)
	}
	return domain.NewChairFromRecord(chair), nil
}

func (c *ChairRepository) Create(ctx context.Context, entity *domain.Chair) error {
	param := mysql.CreateChairParams{
		ID:          entity.ID.Value(),
		Name:        entity.Name.Value(),
		Description: entity.Description,
		Thumbnail:   entity.Thumbnail.Value(),
		Price:       int32(entity.Price),
		Height:      int32(entity.Height),
		Width:       int32(entity.Width),
		Depth:       int32(entity.Depth),
		Color:       entity.Color,
		Features:    strings.Join(entity.Features, ","),
		Kind:        entity.Kind,
		Popularity:  int32(entity.Popularity),
		Stock:       int32(entity.Stock),
	}
	if err := c.queries.CreateChair(ctx, param); err != nil {
		return xerrors.Errorf("failed to insert the row: %w", err)
	}
	return nil
}
