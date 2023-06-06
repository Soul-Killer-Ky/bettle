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
	FindByName(context.Context, string) (*Group, error)
	Join(context.Context, int, string) error
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
	err = c.repo.Join(ctx, userID, groupName)
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
		Type: req.Type,
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
