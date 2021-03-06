package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *authOpsResolver) Register(ctx context.Context, obj *model.AuthOps, input model.NewUser) (*model.AuthentificationToken, error) {
	return service.UserRegister(ctx, input)
}

func (r *authOpsResolver) Login(ctx context.Context, obj *model.AuthOps, email string, password string) (*model.AuthentificationToken, error) {
	return service.UserLogin(ctx, email, password)
}

// AuthOps returns generated.AuthOpsResolver implementation.
func (r *Resolver) AuthOps() generated.AuthOpsResolver { return &authOpsResolver{r} }

type authOpsResolver struct{ *Resolver }
