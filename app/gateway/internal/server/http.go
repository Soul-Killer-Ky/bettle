package server

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	http2 "net/http"
	"strings"

	pbIm "beetle/api/im/service/v1"
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
	imServerEndpoint   = flag.String("im-server-endpoint", "im-svc:9000", "gRPC server endpoint")
)

type wrapper struct {
	mux http2.Handler
	log *log.Helper
}

func (wr *wrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.ToLower(strings.Split(r.Header.Get("Content-Type"), ";")[0]) == "application/x-www-form-urlencoded" {
		wr.log.Info("Rewriting form data as json")
		if err := r.ParseForm(); err != nil {
			http2.Error(w, err.Error(), http2.StatusBadRequest)
			wr.log.Errorf("Bad form request: %s", err.Error())
			return
		}
		jsonMap := make(map[string]interface{}, len(r.Form))
		for k, v := range r.Form {
			if len(v) > 0 {
				jsonMap[k] = v[0]
			}
		}
		jsonBody, err := json.Marshal(jsonMap)
		if err != nil {
			http2.Error(w, err.Error(), http2.StatusBadRequest)
		}

		r.Body = ioutil.NopCloser(bytes.NewReader(jsonBody))
		r.ContentLength = int64(len(jsonBody))
		r.Header.Set("Content-Type", "application/json")
	}
	wr.mux.ServeHTTP(w, r)
}

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
	err = pbIm.RegisterImHandlerFromEndpoint(context.Background(), mux, *imServerEndpoint, gopts)
	if err != nil {
		panic(err)
	}

	l := log.NewHelper(logger)
	srv := http.NewServer(opts...)
	srv.HandleFunc("/chat", func(w http2.ResponseWriter, r *http2.Request) {
		l.Info("websocket")
		websocketProxy(w, r)
	})
	formWrapper := wrapper{mux: mux, log: l}
	srv.HandlePrefix("/", &formWrapper)

	//r := srv.Route("/")
	//r.GET("/", func(ctx http.Context) error {
	//	mux.ServeHTTP(ctx.Response(), ctx.Request())
	//	return nil
	//})

	return srv
}
