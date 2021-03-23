package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *postCommendOpsResolver) Create(ctx context.Context, obj *model.PostCommendOps, input model.NewPostCommend) (*model.PostCommend, error) {
	return service.PostCommendCreate(ctx, input)
}

func (r *postCommendOpsResolver) Delete(ctx context.Context, obj *model.PostCommendOps, id int) (string, error) {
	return service.PostCommendDelete(ctx, id)
}

// PostCommendOps returns generated.PostCommendOpsResolver implementation.
func (r *Resolver) PostCommendOps() generated.PostCommendOpsResolver {
	return &postCommendOpsResolver{r}
}

type postCommendOpsResolver struct{ *Resolver }
