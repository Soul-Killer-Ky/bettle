package ws

import (
	"context"
	"time"

	pb "beetle/api/im/service/v1"

	"github.com/Soul-Killer-Ky/kratos/websocket"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) OnChatMessage(session *websocket.Session, payload *pb.PersonalChatMessage) error {
	payload.Timestamp = timestamppb.New(time.Now())
	s.log.Infof("on chat message session id: %s, bid: %v, payload %v", session.SessionID(), session.BusinessID(), payload)
	s.ws.SendMessage(session.SessionID(), websocket.MessageType(pb.MessageType_PersonalChat), payload)
	if !s.ws.SendMessageByBID(payload.Sender, websocket.MessageType(pb.MessageType_PersonalChat), payload) {
		err := s.mc.SaveMessage(context.Background(), pb.MessageType_PersonalChat, payload)
		s.log.Infof("OnChatMessage err: %v", err)
	}
	return nil
}
