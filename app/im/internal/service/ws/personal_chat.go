package ws

import (
	"context"
	"time"

	pb "beetle/api/im/service/v1"

	"github.com/Soul-Killer-Ky/kratos/websocket"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) OnPersonalChatMessage(session *websocket.Session, payload *pb.PersonalChatMessage) error {
	payload.Timestamp = timestamppb.New(time.Now())
	s.log.Infof("on personal-chat message session id: %s, bid: %v, payload %v", session.SessionID(), session.BusinessID(), payload)
	//s.ws.SendMessage(session.SessionID(), websocket.MessageType(pb.MessageType_PersonalChat), payload)
	err := s.mc.SaveMessage(context.Background(), pb.MessageType_PersonalChat, payload)
	if err != nil {
		s.log.Errorf("save personal-chat message error: %s", err)
		return err
	}
	s.SendMessage(int(payload.Sender), pb.MessageType_PersonalChat, payload)

	return nil
}
