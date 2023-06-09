package data

import (
	pb "beetle/api/im/service/v1"
	"beetle/app/im/internal/biz"
	"beetle/app/im/internal/data/ent/groupchatmessage"
	"beetle/app/im/internal/data/ent/personalchatmessage"
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

func (r *messageRepo) SavePersonalChatMessage(ctx context.Context, from int, sender int, message string) error {
	_, err := r.data.db.PersonalChatMessage.Create().SetFrom(from).SetSender(sender).SetMessage(message).Save(ctx)
	return err
}

func (r *messageRepo) GetPersonalChatMessages(ctx context.Context, sender int, lastTime time.Time) ([]*pb.PersonalChatMessage, error) {
	ps, err := r.data.db.PersonalChatMessage.Query().
		Where(personalchatmessage.Sender(sender), personalchatmessage.CreatedAtGT(lastTime)).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*pb.PersonalChatMessage, 0)
	jsonCodec := encoding.GetCodec("json")
	for _, p := range ps {
		t := &pb.PersonalChatMessage{}
		err := jsonCodec.Unmarshal([]byte(p.Message), t)
		if err != nil {
			return nil, err
		}
		rv = append(rv, t)
	}
	return rv, nil
}

func (r *messageRepo) SaveGroupChatMessage(ctx context.Context, from int, groupID int, message string) error {
	_, err := r.data.db.GroupChatMessage.Create().SetFrom(from).SetGroupID(groupID).SetMessage(message).Save(ctx)
	return err
}

func (r *messageRepo) GetGroupChatMessages(ctx context.Context, groupID int, lastTime time.Time) ([]*pb.GroupChatMessage, error) {
	ps, err := r.data.db.GroupChatMessage.Query().
		Where(groupchatmessage.GroupID(groupID), groupchatmessage.CreatedAtGT(lastTime)).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*pb.GroupChatMessage, 0)
	jsonCodec := encoding.GetCodec("json")
	for _, p := range ps {
		t := &pb.GroupChatMessage{}
		err := jsonCodec.Unmarshal([]byte(p.Message), t)
		if err != nil {
			return nil, err
		}
		rv = append(rv, t)
	}
	return rv, nil
}

func (r *messageRepo) SetCache(ctx context.Context, messageKey string, message interface{}) error {
	cmd := r.data.rdb.LPush(ctx, messageKey, message)
	return cmd.Err()
}
