package ws

import (
	pb "beetle/api/im/service/v1"
	"beetle/app/im/internal/biz"
	"context"
	"github.com/google/uuid"

	"github.com/Soul-Killer-Ky/kratos/websocket"
	"github.com/go-kratos/kratos/v2/log"
)

type Service struct {
	pb.UnimplementedImServer

	mc  *biz.MessageUseCase
	lrc *biz.LoadRecordUseCase
	log *log.Helper

	ws *websocket.Server
}

func NewService(mc *biz.MessageUseCase, lrc *biz.LoadRecordUseCase, logger log.Logger) *Service {
	return &Service{mc: mc, lrc: lrc, log: log.NewHelper(logger)}
}

func (s *Service) SetWebsocketServer(ws *websocket.Server) {
	s.ws = ws
}

func (s *Service) OnWebsocketConnect(sessionID websocket.SessionID, connect bool) {
	s.log.Infof("on websocket conn, session id: %s, conn: %v", sessionID, connect)
	if connect {
		session, ok := s.ws.GetSession(sessionID)
		if !ok {
			s.log.Errorf("not found session, session id: %s", sessionID)
			return
		}
		uid := session.BusinessID().(int)
		s.lrc.Test(context.Background(), uid)
		deviceIDStr := session.Request().FormValue("d")
		deviceID, err := uuid.Parse(deviceIDStr)
		if err != nil {
			s.log.Errorf("parse device id error: %v, device id: %s", err, deviceIDStr)
			return
		}
		s.LoadMessage(uid, deviceID)
	}
}

func (s *Service) LoadMessage(uid int, deviceID uuid.UUID) {
	ctx := context.Background()
	s.log.Infof("load message user id: %d, deviceID: %s", uid, deviceID.String())
	lastTime, err := s.lrc.GetLastTime(ctx, uid, deviceID)
	if err != nil {
		s.log.Errorf("get load last time error: %v", err)
		return
	}
	messages, err := s.mc.LoadChatMessage(ctx, uid, *lastTime)
	if err == nil {
		for _, message := range messages {
			s.log.Infof("send message user id: %d, body: %s", message.Sender, message.Content.Body)
			s.ws.SendMessageByBID(int(message.Sender), websocket.MessageType(pb.MessageType_PersonalChat), message)
		}
	}
	_ = s.lrc.SaveLoadRecord(ctx, uid, deviceID)
}
