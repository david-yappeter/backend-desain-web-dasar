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

func (r *postLikeResolver) User(ctx context.Context, obj *model.PostLike) (*model.User, error) {
	return dataloader.For(ctx).UserByID.Load(obj.UserID)
}

func (r *postLikeOpsResolver) Create(ctx context.Context, obj *model.PostLikeOps, input model.NewPostLike) (*model.PostLike, error) {
	return service.PostLikeCreate(ctx, input)
}

func (r *postLikeOpsResolver) Delete(ctx context.Context, obj *model.PostLikeOps, id int) (string, error) {
	return service.PostLikeDelete(ctx, id)
}

// PostLike returns generated.PostLikeResolver implementation.
func (r *Resolver) PostLike() generated.PostLikeResolver { return &postLikeResolver{r} }

// PostLikeOps returns generated.PostLikeOpsResolver implementation.
func (r *Resolver) PostLikeOps() generated.PostLikeOpsResolver { return &postLikeOpsResolver{r} }

type postLikeResolver struct{ *Resolver }
type postLikeOpsResolver struct{ *Resolver }
