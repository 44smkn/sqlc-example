package usecase

import (
	"context"

	"github.com/44smkn/sqlc-sample/pkg/domain"
	"github.com/44smkn/sqlc-sample/pkg/presentation/param"
	"github.com/44smkn/sqlc-sample/pkg/usecase/dto"
	"golang.org/x/xerrors"
)

func GetChairDetail(ctx context.Context, chairID string) (*dto.ChairDetailDto, error) {
	id, err := domain.NewChairID(chairID)
	if err != nil {
		return nil, xerrors.Errorf("Failed to GetChairDetail: %w", err)
	}
	detail, err := chairRepository.Find(ctx, id)
	if err != nil {
		return nil, xerrors.Errorf("Failed to GetChairDetail: %w", err)
	}
	return &dto.ChairDetailDto{
		ID:          int(detail.ID.Value()),
		Name:        detail.Name.Value(),
		Description: detail.Description,
		Thumbnail:   detail.Thumbnail.Value(),
		Height:      detail.Height,
		Width:       detail.Width,
		Depth:       detail.Depth,
		Color:       detail.Color,
		Features:    detail.Features,
		Kind:        detail.Kind,
		Popularity:  detail.Popularity,
		Stock:       detail.Stock,
	}, nil
}

func PostChair(ctx context.Context, param param.PostChairParam) error {
	name, err := domain.NewChairName(param.Name)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	thumbnail, err := domain.NewChairThumbnail(param.Thumbnail)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	entity := &domain.Chair{
		Name:        *name,
		Description: param.Description,
		Thumbnail:   *thumbnail,
		Height:      int(param.Height),
		Width:       int(param.Width),
		Depth:       int(param.Depth),
		Color:       param.Color,
		Features:    param.Features,
		Kind:        param.Kind,
		Popularity:  int(param.Popularity),
		Stock:       int(param.Stock),
	}
	err = chairRepository.Create(ctx, entity)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	return nil
}
