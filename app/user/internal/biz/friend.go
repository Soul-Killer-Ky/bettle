package biz

import (
	"beetle/app/user/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Friend struct {
	ID       int
	UserID   int
	Username string
	Nickname string
	Avatar   string
	Memo     string
}

type FriendRepo interface {
	ListByUserID(context.Context, int) ([]*Friend, error)
}

type FriendUseCase struct {
	authConf   *conf.Auth
	friendRepo FriendRepo
	log        *log.Helper
}

func NewFriendUseCase(authConf *conf.Auth, friendRepo FriendRepo, logger log.Logger) *FriendUseCase {
	return &FriendUseCase{authConf: authConf, friendRepo: friendRepo, log: log.NewHelper(logger)}
}

func (uc *FriendUseCase) GetFriends(ctx context.Context, userID int) ([]*Friend, error) {
	friends, err := uc.friendRepo.ListByUserID(ctx, userID)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("list friends error: %s", err)
		return nil, errors.BadRequest("QUERY_ERROR", "list friends error")
	}

	return friends, nil
}
