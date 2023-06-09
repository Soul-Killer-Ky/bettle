package biz

import (
	"context"

	pb "beetle/api/user/service/v1"
	"beetle/app/user/internal/data/ent"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Group struct {
	ID   int
	Type int32
	Name string
	Icon string
	Memo string
}

// GroupRepo is a Greater repo.
type GroupRepo interface {
	ListByUserID(context.Context, int) ([]*Group, error)
	ListUserByID(context.Context, int) ([]*User, error)
	FindByID(context.Context, int) (*Group, error)
	FindByName(context.Context, string) (*Group, error)
	IsJoined(context.Context, int, int) (bool, error)
	Join(context.Context, int, int) error
	Create(context.Context, int, *Group) error
}

// GroupUseCase is a group usecase.
type GroupUseCase struct {
	log *log.Helper

	repo GroupRepo
}

// NewGroupUseCase new a im usecase.
func NewGroupUseCase(repo GroupRepo, logger log.Logger) *GroupUseCase {
	return &GroupUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c *GroupUseCase) ListGroup(ctx context.Context, userID int) ([]*Group, error) {
	groups, err := c.repo.ListByUserID(ctx, userID)
	if err != nil {
		c.log.WithContext(ctx).Errorf("list groups error: %s", err)
		return nil, errors.BadRequest("QUERY_ERROR", "list groups error")
	}

	return groups, nil
}

func (c *GroupUseCase) JoinGroup(ctx context.Context, userID int, groupName string) (*Group, error) {
	group, err := c.repo.FindByName(ctx, groupName)
	if err != nil {
		c.log.WithContext(ctx).Errorf("get group error: %s", err)
		return nil, errors.BadRequest("QUERY_ERROR", "get group error")
	}
	isJoined, err := c.repo.IsJoined(ctx, userID, group.ID)
	if err != nil {
		c.log.Errorf("request is join error: %s", err)
		return nil, errors.BadRequest("QUERY_ERROR", "request is join error")
	}
	if isJoined {
		return nil, errors.BadRequest("JOIN_ERROR", "the group already join")
	}
	err = c.repo.Join(ctx, userID, group.ID)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (c *GroupUseCase) CreateGroup(ctx context.Context, userID int, req *pb.CreateGroupRequest) (*Group, error) {
	_, err := c.repo.FindByName(ctx, req.Name)
	if !ent.IsNotFound(err) {
		return nil, pb.ErrorItemExists("the group already exists")
	}
	gp := &Group{
		Type: int32(req.Type.Number()),
		Name: req.Name,
		Icon: req.Icon,
		Memo: req.Memo,
	}
	err = c.repo.Create(ctx, userID, gp)
	if err != nil {
		return nil, err
	}

	return gp, nil
}

func (c *GroupUseCase) ListUserByID(ctx context.Context, groupID int) ([]*User, error) {
	_, err := c.repo.FindByID(ctx, groupID)
	if err != nil {
		return nil, errors.BadRequest("QUERY_ERROR", "the group dont exist")
	}
	users, err := c.repo.ListUserByID(ctx, groupID)
	if err != nil {
		return nil, errors.BadRequest("QUERY_ERROR", "list users error")
	}

	return users, nil
}
