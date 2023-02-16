package service

import (
	pb "beetle/api/im/service/v1"

	"github.com/Soul-Killer-Ky/kratos/websocket"
)

func (s *ImService) OnGroupMessage(session *websocket.Session, payload *pb.GroupMessage) error {
	s.log.Infof("on group message sid: %s, payload %v", session.SessionID(), payload)
	return nil
}
