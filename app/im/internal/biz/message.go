package biz

import (
	"context"
	"strconv"
	"time"

	pb "beetle/api/im/service/v1"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Message struct {
	MessageType pb.MessageType
	Payload     interface{}
}

type MessageRepo interface {
	SavePersonalChatMessage(context.Context, int, int, string) error
	GetPersonalChatMessages(context.Context, int, time.Time) ([]*pb.PersonalChatMessage, error)
	SaveGroupChatMessage(context.Context, int, int, string) error
	GetGroupChatMessages(context.Context, int, time.Time) ([]*pb.GroupChatMessage, error)
	SetCache(context.Context, string, interface{}) error
}

type MessageUseCase struct {
	log *log.Helper

	repo MessageRepo
}

func NewMessageUseCase(repo MessageRepo, logger log.Logger) *MessageUseCase {
	return &MessageUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c *MessageUseCase) SaveMessage(ctx context.Context, messageType pb.MessageType, payload interface{}) (err error) {
	jsonCodec := encoding.GetCodec("json")
	switch messageType {
	case pb.MessageType_PersonalChat:
		t := payload.(*pb.PersonalChatMessage)
		messageKey := messageType.String() + "-" + strconv.FormatInt(t.Sender, 10)
		buf, err := jsonCodec.Marshal(t)
		if err != nil {
			return err
		}
		err = c.repo.SavePersonalChatMessage(ctx, int(t.From), int(t.Sender), string(buf))
		if err != nil {
			return err
		}
		err = c.repo.SetCache(ctx, messageKey, buf)
	case pb.MessageType_GroupChat:
		t := payload.(*pb.GroupChatMessage)
		messageKey := messageType.String() + "-" + strconv.FormatInt(t.GroupId, 10)
		buf, err := jsonCodec.Marshal(t)
		if err != nil {
			return err
		}
		err = c.repo.SaveGroupChatMessage(ctx, int(t.From), int(t.GroupId), string(buf))
		if err == nil {
			err = c.repo.SetCache(ctx, messageKey, buf)
		}
	default:
		err = errors.BadRequest("MESSAGE_ERROR", "message type not active")
	}
	return err
}

func (c *MessageUseCase) GetChatMessages(ctx context.Context, uid int, lastTime time.Time) ([]*Message, error) {
	pms, err := c.repo.GetPersonalChatMessages(ctx, uid, lastTime)
	if err != nil {
		c.log.WithContext(ctx).Errorf("get personal-chat message error: %s", err)
		return nil, err
	}
	rv := make([]*Message, 0)
	for _, pm := range pms {
		rv = append(rv, &Message{
			MessageType: pb.MessageType_PersonalChat,
			Payload:     pm,
		})
	}
	gms, err := c.repo.GetPersonalChatMessages(ctx, uid, lastTime)
	if err != nil {
		c.log.WithContext(ctx).Errorf("get group-chat message error: %s", err)
		return nil, err
	}
	for _, gm := range gms {
		rv = append(rv, &Message{
			MessageType: pb.MessageType_GroupChat,
			Payload:     gm,
		})
	}

	return rv, nil
}
