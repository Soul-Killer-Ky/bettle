package data

import (
	pb "beetle/api/im/service/v1"
	"beetle/app/im/internal/biz"
	"beetle/app/im/internal/data/ent/groupchatmessage"
	"beetle/app/im/internal/data/ent/personalchatmessage"
	"context"
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

func (r *messageRepo) SavePersonalChatMessage(ctx context.Context, from int, sender int, messageID int64, message string) error {
	_, err := r.data.db.PersonalChatMessage.
		Create().
		SetMessageID(messageID).
		SetFrom(from).
		SetSender(sender).
		SetMessage(message).
		Save(ctx)
	return err
}

func (r *messageRepo) GetPersonalChatMessages(ctx context.Context, sender int, lastMessageID int64) ([]*pb.PersonalChatMessage, error) {
	ps, err := r.data.db.PersonalChatMessage.Query().
		Where(
			personalchatmessage.MessageIDGT(lastMessageID),
			personalchatmessage.Or(
				personalchatmessage.Sender(sender),
				personalchatmessage.From(sender),
			),
		).
		All(ctx)
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

func (r *messageRepo) SaveGroupChatMessage(ctx context.Context, from int, groupID int, messageID int64, message string) error {
	_, err := r.data.db.GroupChatMessage.
		Create().
		SetMessageID(messageID).
		SetFrom(from).
		SetGroupID(groupID).
		SetMessage(message).
		Save(ctx)
	return err
}

func (r *messageRepo) GetGroupChatMessages(ctx context.Context, groupIDs []int, lastMessageID int64) ([]*pb.GroupChatMessage, error) {
	ps, err := r.data.db.GroupChatMessage.Query().
		Where(
			groupchatmessage.MessageIDGT(lastMessageID),
			groupchatmessage.GroupIDIn(groupIDs...),
		).
		All(ctx)
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
