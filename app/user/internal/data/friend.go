package data

import (
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
	ps, err := u.QueryUsers().All(ctx)
	if err != nil {
		return nil, err
	}
	r.log.Warn(ps)
	rv := make([]*biz.Friend, 0)
	for _, p := range ps {
		fu := p.QueryFriend().OnlyX(ctx)
		r.log.Info(fu)
		rv = append(rv, &biz.Friend{
			ID:       p.ID,
			UserID:   fu.ID,
			Username: fu.Username,
			Nickname: fu.Nickname,
			Avatar:   fu.Avatar,
			Memo:     fu.Memo,
		})
	}
	return rv, nil
}
