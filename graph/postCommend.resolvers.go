package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/dataloader"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *postCommendResolver) User(ctx context.Context, obj *model.PostCommend) (*model.User, error) {
	return dataloader.For(ctx).UserByID.Load(obj.UserID)
}

func (r *postCommendOpsResolver) Create(ctx context.Context, obj *model.PostCommendOps, input model.NewPostCommend) (*model.PostCommend, error) {
	return service.PostCommendCreate(ctx, input)
}

func (r *postCommendOpsResolver) Delete(ctx context.Context, obj *model.PostCommendOps, id int) (string, error) {
	return service.PostCommendDelete(ctx, id)
}

// PostCommend returns generated.PostCommendResolver implementation.
func (r *Resolver) PostCommend() generated.PostCommendResolver { return &postCommendResolver{r} }

// PostCommendOps returns generated.PostCommendOpsResolver implementation.
func (r *Resolver) PostCommendOps() generated.PostCommendOpsResolver {
	return &postCommendOpsResolver{r}
}

type postCommendResolver struct{ *Resolver }
type postCommendOpsResolver struct{ *Resolver }
