package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jhabshoosh/etzer-api/internal/graph/generated"
	"github.com/jhabshoosh/etzer-api/internal/graph/model"
	"github.com/jhabshoosh/etzer-api/internal/models"
)

// CreatePerson is the resolver for the createPerson field.
func (r *mutationResolver) CreatePerson(ctx context.Context, input model.CreatePersonInput) (*models.Person, error) {
	return r.PersonService.CreatePerson(ctx, input)
}

// UpdateParents is the resolver for the updateParents field.
func (r *mutationResolver) UpdateParents(ctx context.Context, input model.UpdateParentsInput) (string, error) {
	return r.PersonService.UpdateParents(ctx, input)
}

// Parents is the resolver for the parents field.
func (r *personResolver) Parents(ctx context.Context, obj *models.Person) ([]*models.Person, error) {
	return r.PersonService.Parents(ctx, obj)
}

// Children is the resolver for the children field.
func (r *personResolver) Children(ctx context.Context, obj *models.Person) ([]*models.Person, error) {
	return r.PersonService.Children(ctx, obj)
}

// GetPerson is the resolver for the getPerson field.
func (r *queryResolver) GetPerson(ctx context.Context, input model.GetPersonInput) (*models.Person, error) {
	return r.PersonService.GetPerson(ctx, input)
}

// GetRootAncestor is the resolver for the getRootAncestor field.
func (r *queryResolver) GetRootAncestor(ctx context.Context) (*models.Person, error) {
	return r.PersonService.GetRootAncestor(ctx)
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
