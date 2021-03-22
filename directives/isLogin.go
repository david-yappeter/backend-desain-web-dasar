package directives

import (
	"context"
	"myapp/service"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

//IsLogin is Logged In Directives
func IsLogin(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	if service.ForContext(ctx) == nil {
		return nil, &gqlerror.Error{
			Message: "Access Denied: Not Logged In",
			Extensions: map[string]interface{}{
				"code": "UNAUTHENTICATED",
			},
		}
	}
	return next(ctx)
}
