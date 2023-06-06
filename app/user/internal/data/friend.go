package data

import (
	"beetle/app/user/internal/data/ent"
	"beetle/app/user/internal/data/ent/user"
	"context"

	"beetle/app/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type friendRepo struct {
	data *Data
	log  *log.Helper
}

func NewFriendRepo(data *Data, logger log.Logger) biz.FriendRepo {
	return &friendRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *friendRepo) ListByUserID(ctx context.Context, userID int) ([]*biz.Friend, error) {
	u, err := r.data.db.User.Query().Where(user.ID(userID)).First(ctx)
	if err != nil {
		return nil, err
	}
	ps, err := u.QueryFriends().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Friend, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Friend{
			UserID:   p.ID,
			Phone:    p.Phone,
			Username: p.Username,
			Nickname: p.Nickname,
			Avatar:   p.Avatar,
			Memo:     p.Memo,
		})
	}
	return rv, nil
}

func (r *friendRepo) Get(ctx context.Context, userID, friendUserID int) (*biz.Friend, error) {
	u, err := r.data.db.User.Query().Where(user.ID(userID)).First(ctx)
	if err != nil {
		return nil, err
	}
	u2, err := u.QueryFriends().Where(user.ID(friendUserID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.Friend{
		UserID:   u2.ID,
		Phone:    u2.Phone,
		Username: u2.Username,
		Nickname: u2.Nickname,
		Avatar:   u2.Avatar,
		Memo:     u2.Memo,
	}, nil
}

func (r *friendRepo) Add(ctx context.Context, userID, friendUserID int) error {
	err := WithTx(ctx, r.data.db, func(tx *ent.Tx) error {
		u, err := tx.User.Query().Where(user.ID(userID)).First(ctx)
		if err != nil {
			return err
		}
		u2, err := tx.User.Query().Where(user.ID(friendUserID)).First(ctx)
		if err != nil {
			return err
		}
		_, err = u.Update().AddFriends(u2).Save(ctx)
		if err != nil {
			return err
		}
		_, err = u2.Update().AddFriends(u).Save(ctx)
		if err != nil {
			return err
		}
		return err
	})
	return err
}
