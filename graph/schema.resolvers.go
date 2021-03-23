package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *mutationResolver) User(ctx context.Context) (*model.UserOps, error) {
	return &model.UserOps{}, nil
}

func (r *mutationResolver) Auth(ctx context.Context) (*model.AuthOps, error) {
	return &model.AuthOps{}, nil
}

func (r *mutationResolver) Post(ctx context.Context) (*model.PostOps, error) {
	return &model.PostOps{}, nil
}

func (r *mutationResolver) PostLike(ctx context.Context) (*model.PostLikeOps, error) {
	return &model.PostLikeOps{}, nil
}

func (r *mutationResolver) PostCommend(ctx context.Context) (*model.PostCommendOps, error) {
	return &model.PostCommendOps{}, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return service.UserGetByToken(ctx)
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	return service.UserGetByID(ctx, id)
}

func (r *queryResolver) Users(ctx context.Context, limit *int, page *int, sortBy *string, ascending *bool) (*model.UserPagination, error) {
	return &model.UserPagination{Limit: limit, Page: page, SortBy: sortBy, Ascending: ascending}, nil
}

func (r *queryResolver) Post(ctx context.Context, id int) (*model.Post, error) {
	return service.PostGetByID(ctx, id)
}

func (r *queryResolver) Posts(ctx context.Context, limit *int, page *int, sortBy *string, ascending bool) (*model.PostPagination, error) {
	return &model.PostPagination{Limit: limit, Page: page, SortBy: sortBy, Ascending: &ascending}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
