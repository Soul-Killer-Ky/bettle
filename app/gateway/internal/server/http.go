package server

import (
	"context"
	"flag"

	pbUser "beetle/api/user/service/v1"
	"beetle/app/gateway/internal/conf"
	"beetle/app/gateway/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// command-line options:
	// gRPC server endpoint
	userServerEndpoint = flag.String("user-server-endpoint", "user-svc:9000", "gRPC server endpoint")
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
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
	srv.HandlePrefix("/", mux)

	return srv
}
