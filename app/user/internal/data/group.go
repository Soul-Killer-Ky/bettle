package data

import (
	"context"
	"fmt"

	"beetle/app/user/internal/biz"
	"beetle/app/user/internal/data/ent"
	"beetle/app/user/internal/data/ent/group"
	"beetle/app/user/internal/data/ent/user"

	"github.com/go-kratos/kratos/v2/log"
)

type groupRepo struct {
	data *Data
	log  *log.Helper
}

// NewGroupRepo .
func NewGroupRepo(data *Data, logger log.Logger) biz.GroupRepo {
	return &groupRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *groupRepo) ListByUserID(ctx context.Context, userID int) ([]*biz.Group, error) {
	u, err := r.data.db.User.Query().Where(user.ID(userID)).First(ctx)
	if err != nil {
		return nil, err
	}
	ps, err := u.QueryJoinedGroups().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Group, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Group{
			ID:   p.ID,
			Type: p.Type,
			Name: p.Name,
			Icon: p.Icon,
			Memo: p.Memo,
		})
	}
	return rv, nil
}

func (r *groupRepo) FindByName(ctx context.Context, name string) (*biz.Group, error) {
	p, err := r.data.db.Group.Query().Where(group.Name(name)).First(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.Group{
		ID:   p.ID,
		Name: p.Name,
		Icon: p.Icon,
		Memo: p.Memo,
	}, nil
}

func (r *groupRepo) Join(ctx context.Context, userID int, name string) error {
	u, err := r.data.db.User.Query().Where(user.ID(userID)).First(ctx)
	if err != nil {
		return err
	}
	gp, err := u.QueryJoinedGroups().Where(group.Name(name)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("already in the group: %s", name)
		} else {
			return err
		}
	}
	_, err = u.Update().AddJoinedGroups(gp).Save(ctx)
	return err
}

func (r *groupRepo) Create(ctx context.Context, userID int, gp *biz.Group) error {
	err := WithTx(ctx, r.data.db, func(tx *ent.Tx) error {
		p, err := tx.Group.Create().
			SetType(gp.Type).
			SetName(gp.Name).
			SetIcon(gp.Icon).
			SetMemo(gp.Memo).
			SetCreatedByID(userID).
			Save(ctx)
		if err != nil {
			return err
		}
		gp.ID = p.ID
		_, err = tx.GroupMember.Create().
			SetGroup(p).
			SetUserID(userID).
			SetRole(1).
			Save(ctx)
		return err
	})
	return err
}
