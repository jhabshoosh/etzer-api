package db

import (
	"context"

	"github.com/jhabshoosh/etzer/internal/graph/model"
)


type Db interface {
	FindPersons(ctx context.Context) ([]*model.Person, error)
	FindPersonById(ctx context.Context, uuid string) (*model.Person, error)
}