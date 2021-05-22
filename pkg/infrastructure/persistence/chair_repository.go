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
	features := strings.Join(entity.Features, ",")
	param := mysql.CreateChairParams{
		ID:          entity.ID.Value(),
		Name:        entity.Name.Value(),
		Description: entity.Description,
		Thumbnail:   entity.Thumbnail.Value(),
		Price:       convertInt32(&entity.Price),
		Height:      convertInt32(&entity.Height),
		Width:       convertInt32(&entity.Width),
		Depth:       convertInt32(&entity.Depth),
		Color:       convertString(&entity.Color),
		Features:    convertString(&features),
		Kind:        convertString(&entity.Kind),
		Popularity:  int32(entity.Popularity),
		Stock:       int32(entity.Stock),
	}
	if err := c.queries.CreateChair(ctx, param); err != nil {
		return xerrors.Errorf("failed to insert the row: %w", err)
	}
	return nil
}

func (c *ChairRepository) ListWithCondtion(ctx context.Context, cond *domain.ChairSearchCondition) ([]domain.Chair, error) {
	param := mysql.ListChairWithCondtionParams{
		MaxPrice:  convertInt32(cond.MaxPrice),
		MinPrice:  convertInt32(cond.MinPrice),
		MaxWidth:  convertInt32(cond.MaxWidth),
		MinWidth:  convertInt32(cond.MinWidth),
		MaxHeight: convertInt32(cond.MaxHeight),
		MinHeight: convertInt32(cond.MinHeight),
		MaxDepth:  convertInt32(cond.MaxDepth),
		MinDepth:  convertInt32(cond.MinDepth),
	}
	rows, err := c.queries.ListChairWithCondtion(ctx, param)
	if err != nil {
		return nil, xerrors.Errorf("ListWithCondtion is failed: %w", err)
	}
	chairs := make([]domain.Chair, 0, len(rows))
	for _, r := range rows {
		chairs = append(chairs, *domain.NewChairFromRecord(r))
	}
	return chairs, nil

}

func convertInt32(input *int32) sql.NullInt32 {
	if input == nil {
		return sql.NullInt32{Int32: 0, Valid: false}
	}
	return sql.NullInt32{Int32: *input, Valid: true}
}

func convertString(input *string) sql.NullString {
	if input == nil {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: *input, Valid: true}
}
