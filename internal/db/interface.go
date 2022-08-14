package db

import (
	"context"

	"github.com/jhabshoosh/etzer-api/internal/graph/model"
)

type Db interface {
	FetchPersons(ctx context.Context) ([]*model.Person, error)
	FindPersonById(ctx context.Context, uuid string) (*model.Person, error)
}