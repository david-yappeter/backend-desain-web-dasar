package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myapp/graph/generated"
	"myapp/graph/model"
)

func (r *postCommendOpsResolver) Create(ctx context.Context, obj *model.PostCommendOps, input model.NewPostCommend) (*model.PostCommend, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postCommendOpsResolver) Delete(ctx context.Context, obj *model.PostCommendOps, id int) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// PostCommendOps returns generated.PostCommendOpsResolver implementation.
func (r *Resolver) PostCommendOps() generated.PostCommendOpsResolver {
	return &postCommendOpsResolver{r}
}

type postCommendOpsResolver struct{ *Resolver }
