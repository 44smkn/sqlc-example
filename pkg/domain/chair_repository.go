package domain

import "context"

type ChairRepository interface {
	Find(ctx context.Context, id *ChairID) (*Chair, error)
	Create(ctx context.Context, entity *Chair) error
}
