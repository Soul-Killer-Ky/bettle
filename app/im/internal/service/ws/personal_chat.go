package ws

import (
	"context"

	pb "beetle/api/im/service/v1"

	"github.com/Soul-Killer-Ky/kratos/websocket"
)

func (s *Service) OnPersonalChatMessage(session *websocket.Session, payload *pb.PersonalChatMessage) error {
	s.log.Infof("on personal-chat message session id: %s, bid: %v, payload %v", session.SessionID(), session.BusinessID(), payload)
	err := s.mc.SaveMessage(context.Background(), pb.MessageType_PersonalChat, payload)
	if err != nil {
		s.log.Errorf("save personal-chat message error: %s", err)
		return err
	}
	// 发给自己
	s.SendMessage(int(payload.From), pb.MessageType_PersonalChat, payload)
	// 发给目标
	s.SendMessage(int(payload.Sender), pb.MessageType_PersonalChat, payload)

	return nil
}
