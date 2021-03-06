package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/meastblue/teckwatch/generated"
	"github.com/meastblue/teckwatch/model"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTech(ctx context.Context, input *model.NewTech) (*model.Tech, error) {
	tech := model.Tech{
		Label:       input.Label,
		Description: input.Description,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	return &tech, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Techs(ctx context.Context) ([]*model.Tech, error) {
	techs := []*model.Tech{}
	dummuyTech := model.Tech{
		Label:       "GraphQL",
		Description: "A better Rest",
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	techs = append(techs, &dummuyTech)
	return techs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
