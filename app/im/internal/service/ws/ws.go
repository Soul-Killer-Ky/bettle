package ws

import (
	pb "beetle/api/im/service/v1"
	"beetle/app/im/internal/biz"
	"context"

	"github.com/Soul-Killer-Ky/kratos/websocket"
	"github.com/go-kratos/kratos/v2/log"
)

type Service struct {
	pb.UnimplementedImServer

	mc  *biz.MessageUseCase
	log *log.Helper

	ws *websocket.Server
}

func NewService(mc *biz.MessageUseCase, logger log.Logger) *Service {
	return &Service{mc: mc, log: log.NewHelper(logger)}
}

func (s *Service) SetWebsocketServer(ws *websocket.Server) {
	s.ws = ws
}

func (s *Service) OnWebsocketConnect(sessionID websocket.SessionID, connect bool) {
	s.log.Infof("on websocket conn, session id: %s, conn: %v", sessionID, connect)
	if connect {
		session, ok := s.ws.GetSession(sessionID)
		if !ok {
			return
		}
		uid := session.BusinessID().(int)
		s.LoadMessage(uid)
	}
}

func (s *Service) LoadMessage(uid int) {
	s.log.Infof("load message user id: %d", uid)
	messages, err := s.mc.LoadChatMessage(context.Background(), uid)
	if err == nil {
		for _, message := range messages {
			s.log.Infof("send message user id: %d, body: %s", message.Sender, message.Content.Body)
			s.ws.SendMessageByBID(int(message.Sender), websocket.MessageType(pb.MessageType_Chat), message)
		}
	}
}
