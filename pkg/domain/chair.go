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
	Price       int
	Height      int
	Width       int
	Depth       int
	Color       string
	Features    []string
	Kind        string
	Popularity  int
	Stock       int
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
		Height:      int(record.Height),
		Width:       int(record.Width),
		Depth:       int(record.Depth),
		Color:       record.Color,
		Features:    []string{record.Features},
		Kind:        record.Kind,
		Popularity:  int(record.Popularity),
		Stock:       int(record.Stock),
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
