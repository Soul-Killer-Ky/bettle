package data

import (
	"context"

	"beetle/app/im/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
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

func (r *groupRepo) Connect(ctx context.Context) error {
	return nil
}

func (r *groupRepo) CreateClient(ctx context.Context, conn *websocket.Conn, uid uint) *imClient {
	client := imClient{
		ID:     uid,
		Socket: conn,
	}

	return &client
}

func (r *groupRepo) ListByUserID(ctx context.Context, userID int) ([]*biz.Group, error) {
	ps, err := r.data.db.Group.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Group, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Group{
			ID:   p.ID,
			Name: p.Name,
			Icon: p.Icon,
			Memo: p.Memo,
		})
	}
	return rv, nil
}
