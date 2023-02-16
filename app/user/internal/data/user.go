package data

import (
	"context"

	"beetle/app/user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

func (r *userRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

func (r *userRepo) FindByID(context.Context, int64) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) FindByName(ctx context.Context, name string) (*biz.User, error) {
	u := biz.User{}
	err := r.data.DB(ctx).Where("username = ?", name).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) ListByHello(context.Context, string) ([]*biz.User, error) {
	return nil, nil
}

func (r *userRepo) ListAll(context.Context) ([]*biz.User, error) {
	return nil, nil
}

type consumeRepo struct {
	data *Data
	log  *log.Helper
}

type Consume struct {
	ID         int64
	UserID     int64
	OrderID    string
	OrderPrice int64
}

// NewConsumeRepo  .
func NewConsumeRepo(data *Data, logger log.Logger) biz.ConsumeRepo {
	return &consumeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *consumeRepo) CreateConsume(ctx context.Context, cs *biz.Consume) (int64, error) {
	result := c.data.DB(ctx).Create(&cs)
	return cs.ID, result.Error
}
