package domain

import "context"

type ChairRepository interface {
	Find(ctx context.Context, id *ChairID) (*Chair, error)
	Create(ctx context.Context, entity *Chair) error
	ListWithCondtion(ctx context.Context, condition *ChairSearchCondition) ([]Chair, error)
}

type ChairSearchCondition struct {
	MaxPrice  *int32
	MinPrice  *int32
	MaxHeight *int32
	MinHeight *int32
	MaxWidth  *int32
	MinWidth  *int32
	MaxDepth  *int32
	MinDepth  *int32
	Kind      *string
	Color     *string
	Features  []string
}
