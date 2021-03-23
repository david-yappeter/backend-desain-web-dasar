package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *postLikeOpsResolver) Create(ctx context.Context, obj *model.PostLikeOps, input model.NewPostLike) (*model.PostLike, error) {
	return service.PostLikeCreate(ctx, input)
}

func (r *postLikeOpsResolver) Delete(ctx context.Context, obj *model.PostLikeOps, id int) (string, error) {
	return service.PostLikeDelete(ctx, id)
}

// PostLikeOps returns generated.PostLikeOpsResolver implementation.
func (r *Resolver) PostLikeOps() generated.PostLikeOpsResolver { return &postLikeOpsResolver{r} }

type postLikeOpsResolver struct{ *Resolver }
