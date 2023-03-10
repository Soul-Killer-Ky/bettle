package server

import (
	"context"

	pb "beetle/api/user/service/v1"
	"beetle/app/user/internal/conf"
	"beetle/app/user/internal/service"
	"github.com/gorilla/handlers"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList[pb.OperationUserLoginUser] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, user *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(
				jwt.Server(
					func(token *jwtv4.Token) (interface{}, error) {
						return []byte(ac.JwtKey), nil
					},
					jwt.WithSigningMethod(jwtv4.SigningMethodHS256),
				),
			).Match(NewWhiteListMatcher()).Build(),
			validate.Validator(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	pb.RegisterUserHTTPServer(srv, user)
	return srv
}
