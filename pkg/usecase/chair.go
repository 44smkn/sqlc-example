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
		Height:      param.Height,
		Width:       param.Width,
		Depth:       param.Depth,
		Color:       param.Color,
		Features:    param.Features,
		Kind:        param.Kind,
		Popularity:  param.Popularity,
		Stock:       param.Stock,
	}
	err = chairRepository.Create(ctx, entity)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	return nil
}

func SearchChair(ctx context.Context, param param.SearchChairParam) (dto.ChairSearchDto, error) {
	condition := domain.ChairSearchCondition{
		MaxPrice:  param.MaxPrice,
		MinPrice:  param.MinPrice,
		MaxWidth:  param.MaxWidth,
		MinWidth:  param.MinWidth,
		MaxDepth:  param.MaxDepth,
		MinDepth:  param.MinDepth,
		MaxHeight: param.MaxHeight,
		MinHeight: param.MinHeight,
	}
	chairs, err := chairRepository.ListWithCondtion(ctx, &condition)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	searchDto := make(dto.ChairSearchDto, 0, len(chairs))
	for _, c := range chairs {
		e := dto.ChairSearchItemDto{
			Name:        c.Name.Value(),
			Description: c.Description,
			Thumbnail:   c.Thumbnail.Value(),
			Height:      c.Height,
			Width:       c.Width,
			Depth:       c.Depth,
			Color:       c.Color,
			Features:    c.Features,
			Kind:        c.Kind,
			Popularity:  c.Popularity,
			Stock:       c.Stock,
		}
		searchDto = append(searchDto, e)
	}
	return searchDto, nil
}
