package data

import (
	"context"

	"beetle/app/user/internal/biz"
	"beetle/app/user/internal/data/ent/user"

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

func (r *userRepo) FindByID(ctx context.Context, id int) (*biz.User, error) {
	u, err := r.data.db.User.Query().Where(user.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Memo:     u.Memo,
	}, nil
}

func (r *userRepo) FindByName(ctx context.Context, name string) (*biz.User, error) {
	u, err := r.data.db.User.Query().Where(user.Username(name)).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Memo:     u.Memo,
	}, nil
}

func (r *userRepo) ListByHello(context.Context, string) ([]*biz.User, error) {
	return nil, nil
}

func (r *userRepo) ListAll(context.Context) ([]*biz.User, error) {
	return nil, nil
}
