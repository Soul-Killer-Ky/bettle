package data

import (
	"beetle/app/user/internal/biz"
	"beetle/app/user/internal/conf"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewTransaction, NewConsumeRepo, NewUserRepo)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

// 用来承载事务的上下文
type contextTxKey struct{}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

// ExecTx gorm Transaction
func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// DB 根据此方法来判断当前的 db 是不是使用 事务的 DB
func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

// NewData .
func NewData(c *conf.Data, db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "transaction/data"))
	cleanup := func() {
		l.Info("closing the data resources")
	}

	return &Data{db: db, log: l}, cleanup, nil
}

// NewDB gorm Connecting to a Database
func NewDB(c *conf.Data, logger log.Logger) *gorm.DB {
	l := log.NewHelper(log.With(logger, "module", "user-service/data/gorm"))
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		l.Fatalf("failed opening connection to mysql: %v", err)
	}
	setupModel(db, logger)

	return db
}

func setupModel(db *gorm.DB, logger log.Logger) {
	l := log.NewHelper(log.With(logger, "module", "user-service.data.setupModel"))
	if err := db.AutoMigrate(&biz.User{}); err != nil {
		l.Fatal(err)
	}
}
