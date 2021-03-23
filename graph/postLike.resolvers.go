package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myapp/graph/generated"
	"myapp/graph/model"
)

func (r *postLikeOpsResolver) Create(ctx context.Context, obj *model.PostLikeOps, input model.NewPostLike) (*model.PostLike, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postLikeOpsResolver) Delete(ctx context.Context, obj *model.PostLikeOps, id int) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// PostLikeOps returns generated.PostLikeOpsResolver implementation.
func (r *Resolver) PostLikeOps() generated.PostLikeOpsResolver { return &postLikeOpsResolver{r} }

type postLikeOpsResolver struct{ *Resolver }
