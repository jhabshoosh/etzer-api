package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jhabshoosh/etzer-api/internal/graph/generated"
	"github.com/jhabshoosh/etzer-api/internal/graph/model"
)

// CreatePerson is the resolver for the createPerson field.
func (r *mutationResolver) CreatePerson(ctx context.Context, input model.CreatePersonInput) (*model.Person, error) {
	return r.PersonService.CreatePerson(ctx, input)
}

// UpdateParent is the resolver for the updateParent field.
func (r *mutationResolver) UpdateParent(ctx context.Context, input model.UpdateParentInput) (string, error) {
	return r.PersonService.UpdateParent(ctx, input)
}

// Father is the resolver for the father field.
func (r *personResolver) Father(ctx context.Context, obj *model.Person) (*model.Person, error) {
	return r.PersonService.GetParent(ctx, obj, "father")
}

// Mother is the resolver for the mother field.
func (r *personResolver) Mother(ctx context.Context, obj *model.Person) (*model.Person, error) {
	return r.PersonService.GetParent(ctx, obj, "mother")
}

// Persons is the resolver for the persons field.
func (r *queryResolver) Persons(ctx context.Context) ([]*model.Person, error) {
	return r.PersonService.FetchPersons(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Person returns generated.PersonResolver implementation.
func (r *Resolver) Person() generated.PersonResolver { return &personResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type personResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
