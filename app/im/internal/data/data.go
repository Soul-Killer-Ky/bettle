package data

import (
	"context"
	"fmt"

	"beetle/app/im/internal/conf"
	"beetle/app/im/internal/data/ent"
	"entgo.io/ent/dialect"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	_ "beetle/app/im/internal/data/ent/runtime"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewImRepo)

// Data .
type Data struct {
	db *ent.Client
	// TODO wrapped database client
	rdb *redis.Client
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

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})

	d := &Data{
		db:  client,
		rdb: rdb,
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
	}

	return d, cleanup, nil
}
