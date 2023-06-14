package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID       int
	Phone    string
	Username *string
	Nickname string
	Avatar   string
	Memo     string
}

type Group struct {
	ID   int
	Type int32
	Name string
	Icon string
	Memo string
}

type UserRepo interface {
	GetUser(context.Context, int) (*User, error)
	ListUserByGroup(context.Context, int) ([]*User, error)
	ListGroup(context.Context) ([]*Group, error)
}

type UserUseCase struct {
	log *log.Helper

	repo UserRepo
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c *UserUseCase) ListUserByGroupID(ctx context.Context, groupID int) ([]*User, error) {
	users, err := c.repo.ListUserByGroup(ctx, groupID)
	if err != nil {
		c.log.Errorf("list user by group error: %s", err)
		return nil, err
	}

	return users, nil
}

func (c *UserUseCase) ListGroup(ctx context.Context) ([]*Group, error) {
	groups, err := c.repo.ListGroup(ctx)
	if err != nil {
		c.log.Errorf("list group error: %s", err)
		return nil, err
	}

	return groups, nil
}
