package data

import (
	"beetle/app/user/internal/conf"
	"beetle/app/user/internal/data/ent"
	"context"
	"fmt"
	"path/filepath"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	kubernetes "github.com/go-kratos/kratos/contrib/registry/kubernetes/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	kubernetesAPI "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	_ "beetle/app/user/internal/data/ent/runtime"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRegistrar,
	NewDiscovery,
	NewUserRepo,
	NewFriendRepo,
	NewGroupRepo,
)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
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

	cleanup := func() {
		l.Info("closing the data resources")
		if err := client.Close(); err != nil {
			l.Error(err)
		}
	}

	return &Data{db: client}, cleanup, nil
}

func NewDiscovery() registry.Discovery {
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

func NewRegistrar() registry.Registrar {
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

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
