package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *userOpsResolver) EditName(ctx context.Context, obj *model.UserOps, input model.EditUserName) (string, error) {
	return service.UserUpdateName(ctx, input.Name)
}

func (r *userOpsResolver) EditPassword(ctx context.Context, obj *model.UserOps, input model.EditUserPassword) (string, error) {
	return service.UserUpdateName(ctx, input.Password)
}

func (r *userPaginationResolver) TotalData(ctx context.Context, obj *model.UserPagination) (int, error) {
	return service.UserPaginationGetTotalData(ctx)
}

func (r *userPaginationResolver) Nodes(ctx context.Context, obj *model.UserPagination) ([]*model.User, error) {
	return service.UserPaginationGetNodes(ctx, obj.Limit, obj.Page, obj.SortBy, obj.Ascending)
}

// UserOps returns generated.UserOpsResolver implementation.
func (r *Resolver) UserOps() generated.UserOpsResolver { return &userOpsResolver{r} }

// UserPagination returns generated.UserPaginationResolver implementation.
func (r *Resolver) UserPagination() generated.UserPaginationResolver {
	return &userPaginationResolver{r}
}

type userOpsResolver struct{ *Resolver }
type userPaginationResolver struct{ *Resolver }
