package data

import (
	pb "beetle/api/user/service/v1"
	"beetle/app/user/internal/biz"
	"beetle/app/user/internal/data/ent"
	"beetle/app/user/internal/data/ent/group"
	"beetle/app/user/internal/data/ent/groupmember"
	"beetle/app/user/internal/data/ent/user"
	"context"

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

func (r *groupRepo) ListUserByID(ctx context.Context, id int) ([]*biz.User, error) {
	members, err := r.data.db.GroupMember.Query().Where(groupmember.GroupID(id)).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.User, 0)
	for _, p := range members {
		u := p.QueryUser().OnlyX(ctx)
		rv = append(rv, &biz.User{
			ID:       u.ID,
			Phone:    u.Phone,
			Username: u.Username,
			Password: u.Password,
			Nickname: u.Nickname,
			Avatar:   u.Avatar,
			Memo:     u.Memo,
		})
	}
	return rv, nil
}

func (r *groupRepo) FindByID(ctx context.Context, id int) (*biz.Group, error) {
	p, err := r.data.db.Group.Query().Where(group.ID(id)).First(ctx)
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

func (r *groupRepo) IsJoined(ctx context.Context, userID, groupID int) (exist bool, err error) {
	return r.data.db.GroupMember.Query().
		Where(groupmember.UserID(userID), groupmember.GroupID(groupID)).
		Exist(ctx)
}

func (r *groupRepo) Join(ctx context.Context, userID int, groupID int) error {
	_, err := r.data.db.GroupMember.Create().
		SetGroupID(groupID).
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
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
			SetRole(int32(pb.GroupUserRole_Administrator)).
			Save(ctx)
		return err
	})
	return err
}
