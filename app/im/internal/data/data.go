package data

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	pbUser "beetle/api/user/service/v1"
	"beetle/app/im/internal/conf"
	"beetle/app/im/internal/data/ent"
	jwt2 "beetle/internal/pkg/jwt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	kubernetes "github.com/go-kratos/kratos/contrib/registry/kubernetes/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	kubernetesAPI "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	_ "beetle/app/im/internal/data/ent/runtime"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewK8sRegistry,
	NewRegistrar,
	NewDiscovery,
	NewUserServiceClient,
	NewMessageRepo,
	NewLoadRecordRepo,
)

// Data .
type Data struct {
	db       *ent.Client
	rdb      *redis.Client
	registry *kubernetes.Registry
	uc       pbUser.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, uc pbUser.UserClient, registry *kubernetes.Registry) (*Data, func(), error) {
	l := log.NewHelper(logger)
	drv, err := sql.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		l.WithContext(ctx).Info(i...)
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	client := ent.NewClient(ent.Driver(sqlDrv), ent.Debug())
	if err != nil {
		l.Errorf("failed opening connection to sql: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err = rdb.Ping(timeout).Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}

	d := &Data{
		db:       client,
		rdb:      rdb,
		uc:       uc,
		registry: registry,
	}
	cleanup := func() {
		l := log.NewHelper(logger)
		l.Info("closing the data resources")
		if err := d.db.Close(); err != nil {
			l.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
		d.registry.Close()
	}

	return d, cleanup, nil
}

func NewK8sRegistry() *kubernetes.Registry {
	restConfig, err := rest.InClusterConfig()
	home := homedir.HomeDir()
	if err != nil {
		kubeConfig := filepath.Join(home, ".kube", "config")
		restConfig, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			panic(err)
		}
	}
	clientSet, err := kubernetesAPI.NewForConfig(restConfig)
	if err != nil {
		panic(err)
	}
	r := kubernetes.NewRegistry(clientSet)
	return r
}

func NewDiscovery(registry *kubernetes.Registry) registry.Discovery {
	registry.Start()
	return registry
}

func NewRegistrar(registry *kubernetes.Registry) registry.Registrar {
	return registry
}

func NewUserServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider, ac *conf.Auth) pbUser.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user-svc"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
			jwt.Client(
				func(token *jwtv4.Token) (interface{}, error) {
					return []byte(ac.JwtKey), nil
				},
				jwt.WithClaims(func() jwtv4.Claims {
					return &jwt2.CustomUserClaims{}
				}),
			),
		),
	)
	if err != nil {
		panic(err)
	}
	c := pbUser.NewUserClient(conn)
	return c
}
