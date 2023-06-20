package server

import (
	jwt2 "beetle/internal/pkg/jwt"
	"context"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	http2 "net/http"

	pb "beetle/api/filesystem/v1"
	"beetle/app/filesystem/internal/conf"
	"beetle/app/filesystem/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, filesystem *service.FilesystemService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(
				jwt.Server(
					func(token *jwtv4.Token) (interface{}, error) {
						return []byte(ac.JwtKey), nil
					},
					jwt.WithSigningMethod(jwtv4.SigningMethodHS256),
					jwt.WithClaims(func() jwtv4.Claims {
						return &jwt2.CustomUserClaims{}
					}),
				),
				jwt2.Auth(),
			).Match(NewWhiteListMatcher()).Build(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Upgrade", "Authorization"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
			handlers.AllowCredentials(),
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
	route := srv.Route("/")
	route.POST("/v1/upload", filesystemUploadHttpHandler(filesystem))

	srv.HandlePrefix("/", http2.FileServer(http2.Dir("/var/lib/filesystem/")))

	//pb.RegisterFilesystemHTTPServer(srv, filesystem)
	return srv
}

func filesystemUploadHttpHandler(srv *service.FilesystemService) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		http.SetOperation(ctx, pb.OperationFilesystemUpload)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Upload(ctx, req.(http.Context))
		})
		out, err := h(ctx, ctx)
		if err != nil {
			return err
		}
		reply := out.(*pb.UploadReply)
		return ctx.Result(200, reply)
	}
}
