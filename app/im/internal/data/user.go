package data

import (
	pbUser "beetle/api/user/service/v1"
	"context"

	"beetle/app/im/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) GetUser(ctx context.Context, uid int) (*biz.User, error) {
	reply, err := r.data.uc.GetUser(ctx, &pbUser.GetUserRequest{Id: int64(uid)})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:       int(reply.UserId),
		Phone:    reply.Phone,
		Username: &reply.Username,
		Nickname: reply.Nickname,
		Avatar:   reply.Avatar,
		Memo:     reply.Memo,
	}, err
}

func (r *userRepo) ListUserByGroup(ctx context.Context, groupID int) ([]*biz.User, error) {
	reply, err := r.data.uc.ListUserByGroup(ctx, &pbUser.ListUserByGroupRequest{GroupId: int64(groupID)})
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.User, 0)
	for _, u := range reply.Users {
		rv = append(rv, &biz.User{
			ID:       int(u.Id),
			Phone:    u.Phone,
			Username: &u.Username,
			Nickname: u.Nickname,
			Avatar:   u.Avatar,
			Memo:     u.Memo,
		})
	}
	return rv, nil
}

func (r *userRepo) ListGroup(ctx context.Context) ([]*biz.Group, error) {
	reply, err := r.data.uc.ListGroup(ctx, &pbUser.ListGroupRequest{})
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Group, 0)
	for _, g := range reply.Groups {
		rv = append(rv, &biz.Group{
			ID:   int(g.Id),
			Type: g.Type,
			Name: g.Name,
			Icon: g.Icon,
			Memo: g.Memo,
		})
	}
	return rv, nil
}
