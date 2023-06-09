package ws

import (
	"beetle/app/im/internal/biz"
	"context"
	"time"

	pb "beetle/api/im/service/v1"

	"github.com/Soul-Killer-Ky/kratos/websocket"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) OnGroupChatMessage(session *websocket.Session, payload *pb.GroupChatMessage) error {
	payload.Timestamp = timestamppb.New(time.Now())
	s.log.Infof("on group-chat message sid: %s, payload %v", session.SessionID(), payload)
	//s.ws.SendMessage(session.SessionID(), websocket.MessageType(pb.MessageType_GroupChat), payload)

	err := s.mc.SaveMessage(context.Background(), pb.MessageType_GroupChat, payload)
	if err != nil {
		s.log.Errorf("save group-chat message error: %s", err)
		return err
	}

	users, err := s.uc.ListUserByGroupID(context.Background(), int(payload.GroupId))
	if err != nil {
		s.log.Errorf("list group users error: %s", err)
		return err
	}
	s.sendMessageToGroupUsers(users, payload)

	return nil
}

func (s *Service) sendMessageToGroupUsers(users []*biz.User, payload *pb.GroupChatMessage) {
	for _, u := range users {
		s.log.Infof("send group-chat to user: %s", *u.Username)
		go s.SendMessage(u.ID, pb.MessageType_GroupChat, payload)
	}
}
