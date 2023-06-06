package biz

import (
	"context"

	pb "beetle/api/user/service/v1"
	"beetle/app/user/internal/conf"
	"beetle/app/user/internal/data/ent"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Friend struct {
	UserID   int
	Phone    string
	Username *string
	Nickname string
	Avatar   string
	Memo     string
}

type FriendRepo interface {
	ListByUserID(context.Context, int) ([]*Friend, error)
	Get(context.Context, int, int) (*Friend, error)
	Add(context.Context, int, int) error
}

type FriendUseCase struct {
	authConf   *conf.Auth
	friendRepo FriendRepo
	log        *log.Helper
}

func NewFriendUseCase(authConf *conf.Auth, friendRepo FriendRepo, logger log.Logger) *FriendUseCase {
	return &FriendUseCase{authConf: authConf, friendRepo: friendRepo, log: log.NewHelper(logger)}
}

func (uc *FriendUseCase) ListFriend(ctx context.Context, userID int) ([]*Friend, error) {
	friends, err := uc.friendRepo.ListByUserID(ctx, userID)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("list friends error: %s", err)
		return nil, errors.BadRequest("QUERY_ERROR", "list friends error")
	}

	return friends, nil
}

func (uc *FriendUseCase) AddFriend(ctx context.Context, userID, friendUserID int) (*Friend, error) {
	_, err := uc.friendRepo.Get(ctx, userID, friendUserID)
	if !ent.IsNotFound(err) {
		return nil, pb.ErrorItemExists("the user already exists")
	}
	err = uc.friendRepo.Add(ctx, userID, friendUserID)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("add friend error: %s", err)
		return nil, errors.BadRequest("ADD_ERROR", "add friend error")
	}
	friend, _ := uc.friendRepo.Get(ctx, userID, friendUserID)

	return friend, nil
}
