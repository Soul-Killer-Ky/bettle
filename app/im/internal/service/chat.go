package service

import (
	pb "beetle/api/im/service/v1"

	"github.com/Soul-Killer-Ky/kratos/websocket"
)

func (s *ImService) OnChatMessage(session *websocket.Session, payload *pb.ChatMessage) error {
	s.log.Infof("on chat message sid: %s, bid: %v, payload %v", session.SessionID(), session.BusinessID(), payload)
	s.ws.SendMessage(session.SessionID(), websocket.MessageType(pb.ProtocolType_Chat), payload)
	s.ws.SendMessageByBID(payload.Sender, websocket.MessageType(pb.ProtocolType_Chat), payload)
	return nil
}
