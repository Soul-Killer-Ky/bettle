package data

import (
	pb "beetle/api/im/service/v1"
	"beetle/app/im/internal/biz"
	"beetle/app/im/internal/data/ent/chatmessage"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
)

type messageRepo struct {
	data *Data
	log  *log.Helper
}

// NewMessageRepo .
func NewMessageRepo(data *Data, logger log.Logger) biz.MessageRepo {
	return &messageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *messageRepo) SaveChatMessage(ctx context.Context, from int, sender int, message string) error {
	_, err := r.data.db.ChatMessage.Create().SetFrom(from).SetSender(sender).SetMessage(message).Save(ctx)
	return err
}

func (r *messageRepo) GetChatMessages(ctx context.Context, sender int, lastTime time.Time) ([]*pb.ChatMessage, error) {
	ps, err := r.data.db.ChatMessage.Query().
		Where(chatmessage.Sender(sender), chatmessage.CreatedAtGT(lastTime)).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*pb.ChatMessage, 0)
	jsonCodec := encoding.GetCodec("json")
	for _, p := range ps {
		t := &pb.ChatMessage{}
		err := jsonCodec.Unmarshal([]byte(p.Message), t)
		if err != nil {
			return nil, err
		}
		rv = append(rv, t)
	}
	return rv, nil
}

func (r *messageRepo) CacheMessage(ctx context.Context, messageKey string, message interface{}) error {
	cmd := r.data.rdb.LPush(ctx, messageKey, message)
	return cmd.Err()
}
