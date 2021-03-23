package dataloader

import (
	"context"
	"myapp/graph/model"
	"myapp/service"
	"net/http"
	"time"
)

const loadersKey = "dataloaders"

type Loaders struct {
	PostCommendBatchByPostID PostCommendBatchLoaderByPostID
	PostLikeBatchByPostID    PostLikeBatchLoaderByPostID
	UserByID                 UserLoaderByID
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			PostCommendBatchByPostID: PostCommendBatchLoaderByPostID{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []int) ([][]*model.PostCommend, []error) {

					resp, err := service.PostCommendGetByArrayPostID(context.Background(), ids)

					if err != nil {
						return nil, []error{err}
					}

					postCommendById := map[int][]*model.PostCommend{}
					for _, val := range resp {
						postCommendById[val.PostID] = append(postCommendById[val.PostID], val)
					}

					postCommends := make([][]*model.PostCommend, len(ids))
					for i, id := range ids {
						postCommends[i] = postCommendById[id]
					}

					return postCommends, nil
				},
			},
			PostLikeBatchByPostID: PostLikeBatchLoaderByPostID{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []int) ([][]*model.PostLike, []error) {

					resp, err := service.PostLikeGetByArrayPostID(context.Background(), ids)

					if err != nil {
						return nil, []error{err}
					}

					postLikeById := map[int][]*model.PostLike{}
					for _, val := range resp {
						postLikeById[val.PostID] = append(postLikeById[val.PostID], val)
					}

					postLikes := make([][]*model.PostLike, len(ids))
					for i, id := range ids {
						postLikes[i] = postLikeById[id]
					}

					return postLikes, nil
				},
			},
			UserByID: UserLoaderByID{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []int) ([]*model.User, []error) {

					resp, err := service.UserGetByArrayID(context.Background(), ids)

					if err != nil {
						return nil, []error{err}
					}

					userById := map[int]*model.User{}
					for _, val := range resp {
						userById[val.ID] = val
					}

					users := make([]*model.User, len(ids))
					for i, id := range ids {
						users[i] = userById[id]
					}

					return users, nil
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

//For Get
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
