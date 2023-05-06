package server

import (
	"context"
	"flag"
	http2 "net/http"

	pbUser "beetle/api/user/service/v1"
	"beetle/app/gateway/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// command-line options:
	// gRPC server endpoint
	userServerEndpoint = flag.String("user-server-endpoint", "user-svc:9000", "gRPC server endpoint")
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, websocketProxy http2.HandlerFunc, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			logging.Server(logger),
			recovery.Recovery(),
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

	mux := runtime.NewServeMux()
	gopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pbUser.RegisterUserHandlerFromEndpoint(context.Background(), mux, *userServerEndpoint, gopts)
	if err != nil {
		panic(err)
	}

	srv := http.NewServer(opts...)
	srv.HandleFunc("/chat", func(a http2.ResponseWriter, b *http2.Request) {
		websocketProxy(a, b)
	})
	srv.HandlePrefix("/", mux)

	//r := srv.Route("/")
	//r.GET("/", func(ctx http.Context) error {
	//	mux.ServeHTTP(ctx.Response(), ctx.Request())
	//	return nil
	//})

	return srv
}
