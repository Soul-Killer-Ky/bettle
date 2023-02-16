package data

import (
	"context"

	"beetle/app/im/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
)

type imRepo struct {
	data *Data
	log  *log.Helper
}

// NewImRepo .
func NewImRepo(data *Data, logger log.Logger) biz.ImRepo {
	return &imRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *imRepo) Connect(ctx context.Context) error {
	return nil
}

func (r *imRepo) CreateClient(ctx context.Context, conn *websocket.Conn, uid uint) *imClient {
	client := imClient{
		ID:     uid,
		Socket: conn,
	}

	return &client
}
