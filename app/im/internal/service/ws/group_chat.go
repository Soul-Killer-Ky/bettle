package ws

import (
	"context"

	pb "beetle/api/im/service/v1"

	"github.com/Soul-Killer-Ky/kratos/websocket"
)

func (s *Service) OnGroupChatMessage(session *websocket.Session, payload *pb.GroupChatMessage) error {
	s.log.Infof("on group-chat message sid: %s, payload %v", session.SessionID(), payload)
	err := s.mc.SaveMessage(context.Background(), pb.MessageType_GroupChat, payload)
	if err != nil {
		s.log.Errorf("save group-chat message error: %s", err)
		return err
	}

	userID := int(payload.From)
	// 发给自己
	s.SendMessage(userID, pb.MessageType_GroupChat, payload)
	// 发给目标
	go func(retryCount int) {
		for i := 0; i < retryCount; i++ {
			users, err := s.uc.ListUserByGroupID(context.Background(), int(payload.GroupId))
			if err != nil {
				s.log.Errorf("list group users error: %s", err)
				continue
			}
			for _, u := range users {
				if u.ID == userID {
					continue
				}
				s.log.Infof("send group-chat to user: %s", *u.Username)
				s.SendMessage(u.ID, pb.MessageType_GroupChat, payload)
			}
			break
		}
	}(3)

	return nil
}
