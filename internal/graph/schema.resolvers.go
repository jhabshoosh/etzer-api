package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/jhabshoosh/etzer/graph/generated"
	"github.com/jhabshoosh/etzer/graph/model"
)

// CreatePerson is the resolver for the createPerson field.
func (r *mutationResolver) CreatePerson(ctx context.Context, input model.NewPerson) (*model.Person, error) {
	person := model.Person{
		Name: input.Name,
		ID:   uuid.NewString(),
	}

	return &person, nil
}

// Persons is the resolver for the persons field.
func (r *queryResolver) Persons(ctx context.Context) ([]*model.Person, error) {
	var persons []*model.Person
	dummyPerson := model.Person{
		Name: "Me You",
		ID:   uuid.NewString(),
	}
	persons = append(persons, &dummyPerson)
	return persons, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
