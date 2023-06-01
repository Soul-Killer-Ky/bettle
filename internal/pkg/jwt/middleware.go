package jwt

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
)

type userIDKey struct{}

func Auth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			token, ok := jwt.FromContext(ctx)
			if !ok {
				return nil, errors.BadRequest("AUTH_FAILURE", "token parse error")
			}
			claims := token.(*CustomUserClaims)
			ctx = context.WithValue(ctx, userIDKey{}, claims.ID)
			return handler(ctx, req)
		}
	}
}

func GetUserID(ctx context.Context) (userID int, err error) {
	userID, ok := ctx.Value(userIDKey{}).(int)
	if !ok {
		err = errors.BadRequest("AUTH_FAILURE", "get user id error")
	}
	return
}
