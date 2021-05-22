package domain

import (
	"strconv"
	"strings"

	"github.com/44smkn/sqlc-sample/pkg/infrastructure/persistence/mysql"
	"golang.org/x/xerrors"
)

type Chair struct {
	ID          ChairID
	Name        ChairName
	Description string
	Thumbnail   ChairThumbnail
	Price       int32
	Height      int32
	Width       int32
	Depth       int32
	Color       string
	Features    []string
	Kind        string
	Popularity  int32
	Stock       int32
}

func NewChairFromRecord(record mysql.Chair) *Chair {
	id := ChairID{val: record.ID}
	name := ChairName{val: record.Name}
	thumbnail := ChairThumbnail{val: record.Thumbnail}

	return &Chair{
		ID:          id,
		Name:        name,
		Description: record.Description,
		Thumbnail:   thumbnail,
		Height:      record.Height.Int32,
		Width:       record.Width.Int32,
		Depth:       record.Depth.Int32,
		Color:       record.Color.String,
		Features:    []string{record.Features.String},
		Kind:        record.Kind.String,
		Popularity:  record.Popularity,
		Stock:       record.Stock,
	}
}

type ChairID struct {
	val int64
}

func NewChairID(chairID string) (*ChairID, error) {
	id, err := strconv.ParseInt(chairID, 10, 64)
	if err != nil {
		return nil, xerrors.Errorf("chairID must be number: %w", err)
	}
	return &ChairID{id}, nil
}

func (c *ChairID) Value() int64 {
	return c.val
}

type ChairName struct {
	val string
}

func NewChairName(name string) (*ChairName, error) {
	if len(name) > 256 {
		return nil, xerrors.Errorf("name '%v' is too long", name)
	}
	return &ChairName{name}, nil
}

func (c *ChairName) Value() string {
	return c.val
}

type ChairThumbnail struct {
	val string
}

func NewChairThumbnail(url string) (*ChairThumbnail, error) {
	if !strings.HasPrefix(url, "https://") {
		return nil, xerrors.Errorf("thumbnail url must start 'https://': %v", url)
	}
	return &ChairThumbnail{url}, nil
}

func (c *ChairThumbnail) Value() string {
	return c.val
}
