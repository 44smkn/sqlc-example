package usecase

import (
	"context"
	"strconv"

	"github.com/44smkn/sqlc-sample/pkg/config"
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
	condition, err := generateSeacrhCondtion(param)
	if err != nil {
		return nil, err
	}
	chairs, err := chairRepository.ListWithCondtion(ctx, condition)
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

func generateSeacrhCondtion(param param.SearchChairParam) (*domain.ChairSearchCondition, error) {
	maxPrice, minPrice, err := getRange(param.PriceRangeID, chairSearchCondition.Price)
	if err != nil {
		return nil, xerrors.Errorf("getting price range is failed: %w", err)
	}
	maxWidth, minWidth, err := getRange(param.WidthRangeID, chairSearchCondition.Width)
	if err != nil {
		return nil, xerrors.Errorf("getting width range is failed: %w", err)
	}
	maxDepth, minDepth, err := getRange(param.DepthRangeID, chairSearchCondition.Depth)
	if err != nil {
		return nil, xerrors.Errorf("getting depth range is failed: %w", err)
	}
	maxHeight, minHeight, err := getRange(param.HeightRangeID, chairSearchCondition.Height)
	if err != nil {
		return nil, xerrors.Errorf("getting depth range is failed: %w", err)
	}
	condition := domain.ChairSearchCondition{
		MaxPrice:  maxPrice,
		MinPrice:  minPrice,
		MaxWidth:  maxWidth,
		MinWidth:  minWidth,
		MaxDepth:  maxDepth,
		MinDepth:  minDepth,
		MaxHeight: maxHeight,
		MinHeight: minHeight,
	}
	return &condition, nil
}

func getRange(id string, cond config.RangeCondition) (max, min *int32, err error) {
	if id == "" {
		return
	}
	rangeIdx, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	if rangeIdx < 0 || len(cond.Ranges) <= rangeIdx {
		err = xerrors.New("Unexpected Range ID")
		return
	}

	r := cond.Ranges[rangeIdx]
	if r.Min != -1 {
		min = &r.Min
	}
	if r.Max != -1 {
		max = &r.Max
	}
	return
}
