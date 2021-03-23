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

func (r *postResolver) Commends(ctx context.Context, obj *model.Post) ([]*model.PostCommend, error) {
	return dataloader.For(ctx).PostCommendBatchByPostID.Load(obj.ID)
}

func (r *postResolver) Likes(ctx context.Context, obj *model.Post) ([]*model.PostLike, error) {
	return dataloader.For(ctx).PostLikeBatchByPostID.Load(obj.ID)
}

func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	return dataloader.For(ctx).UserByID.Load(obj.UserID)
}

func (r *postOpsResolver) Create(ctx context.Context, obj *model.PostOps, input model.NewPost) (*model.Post, error) {
	return service.PostCreate(ctx, input)
}

func (r *postOpsResolver) Delete(ctx context.Context, obj *model.PostOps, id int) (string, error) {
	return service.PostDelete(ctx, id)
}

func (r *postPaginationResolver) TotalData(ctx context.Context, obj *model.PostPagination) (int, error) {
	return service.PostPaginationGetTotalData(ctx)
}

func (r *postPaginationResolver) Nodes(ctx context.Context, obj *model.PostPagination) ([]*model.Post, error) {
	return service.PostPaginationGetNodes(ctx, obj.Limit, obj.Page, obj.SortBy, obj.Ascending)
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// PostOps returns generated.PostOpsResolver implementation.
func (r *Resolver) PostOps() generated.PostOpsResolver { return &postOpsResolver{r} }

// PostPagination returns generated.PostPaginationResolver implementation.
func (r *Resolver) PostPagination() generated.PostPaginationResolver {
	return &postPaginationResolver{r}
}

type postResolver struct{ *Resolver }
type postOpsResolver struct{ *Resolver }
type postPaginationResolver struct{ *Resolver }
