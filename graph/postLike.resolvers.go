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

func (r *postLikeOpsResolver) LikeOrUnlike(ctx context.Context, obj *model.PostLikeOps, postID int) (*model.PostLike, error) {
	return service.PostLikeLikeOrUnlike(ctx, postID)
}

// PostLike returns generated.PostLikeResolver implementation.
func (r *Resolver) PostLike() generated.PostLikeResolver { return &postLikeResolver{r} }

// PostLikeOps returns generated.PostLikeOpsResolver implementation.
func (r *Resolver) PostLikeOps() generated.PostLikeOpsResolver { return &postLikeOpsResolver{r} }

type postLikeResolver struct{ *Resolver }
type postLikeOpsResolver struct{ *Resolver }
