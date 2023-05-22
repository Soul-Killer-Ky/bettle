package biz

import (
	pb "beetle/api/im/service/v1"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"strconv"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
)

type MessageRepo interface {
	SaveChatMessage(context.Context, int, int, string) error
	GetChatMessages(context.Context, int) ([]*pb.ChatMessage, error)
	CacheMessage(context.Context, string, interface{}) error
}

type MessageUseCase struct {
	log *log.Helper

	repo MessageRepo
}

func NewMessageUseCase(repo MessageRepo, logger log.Logger) *MessageUseCase {
	return &MessageUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (mc *MessageUseCase) SaveMessage(ctx context.Context, messageType pb.MessageType, payload interface{}) (err error) {
	jsonCodec := encoding.GetCodec("json")
	switch messageType {
	case pb.MessageType_Chat:
		t := payload.(*pb.ChatMessage)
		messageKey := messageType.String() + "-" + strconv.FormatUint(t.Sender, 10)
		buf, err := jsonCodec.Marshal(t)
		if err != nil {
			return err
		}
		err = mc.repo.SaveChatMessage(ctx, int(t.From), int(t.Sender), string(buf))
		if err != nil {
			return err
		}
		err = mc.repo.CacheMessage(ctx, messageKey, buf)
	case pb.MessageType_Group:
		t := payload.(*pb.GroupMessage)
		messageKey := messageType.String() + "-" + t.GroupId
		buf, err := jsonCodec.Marshal(t)
		if err == nil {
			err = mc.repo.CacheMessage(ctx, messageKey, buf)
		}
	default:
		err = errors.BadRequest("MESSAGE_ERROR", "message type not active")
	}
	return err
}

func (mc *MessageUseCase) LoadChatMessage(ctx context.Context, uid int) ([]*pb.ChatMessage, error) {
	messages, err := mc.repo.GetChatMessages(ctx, uid)
	if err != nil {
		mc.log.WithContext(ctx).Errorf("load chat message error: %s", err)
		return nil, err
	}

	return messages, nil
}
