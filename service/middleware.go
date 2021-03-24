package service

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

//CorsMiddleware CORS Middleware
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow cross domain AJAX requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS");
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization")
		next.ServeHTTP(w, r)
	})
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")

		if bearerToken == "" {
			next.ServeHTTP(w, r)
			return
		}

		token, err := splitBearer(bearerToken)

		if err != nil {
			fmt.Println(err)
			http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
			return
		}

		tokenBeforeClaim, err := TokenValidate(context.Background(), token)

		claims, ok := tokenBeforeClaim.Claims.(*UserClaim)
		if !ok && !tokenBeforeClaim.Valid {
			fmt.Println(err)
			http.Error(w, "Invalid Token", http.StatusForbidden)
			return
		}

		if err != nil {
			fmt.Println(err)
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		if time.Now().UTC().UnixNano()/int64(time.Millisecond) > claims.ExpiresAt {
			http.Error(w, fmt.Sprint("Token Expired"), http.StatusBadRequest)
			return
		}

		if _, err = UserGetByID(context.Background(), claims.ID); err != nil {
			fmt.Println(err)
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		// put it in context
		ctx := context.WithValue(r.Context(), userCtxKey, claims)

		// and call the next with our new context
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *UserClaim {
	raw, _ := ctx.Value(userCtxKey).(*UserClaim)
	return raw
}

func splitBearer(authorization string) (string, error) {
	splitted := strings.Split(strings.Trim(authorization, " "), " ")

	if len(splitted) != 2 {
		return "Failed", &gqlerror.Error{
			Message: "Invalid Authorization",
			Extensions: map[string]interface{}{
				"code": "INVALID_AUTHORIZATION",
			},
		}
	}

	return splitted[1], nil
}
