package ws

import (
	pbErrors "beetle/api/beetle/errors"
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
	src *biz.SynchronizeRecordUseCase
	uc  *biz.UserUseCase
	log *log.Helper

	ws *websocket.Server
}

func NewService(mc *biz.MessageUseCase, src *biz.SynchronizeRecordUseCase, uc *biz.UserUseCase, logger log.Logger) *Service {
	return &Service{mc: mc, src: src, uc: uc, log: log.NewHelper(logger)}
}

func (s *Service) SetWebsocketServer(ws *websocket.Server) {
	s.ws = ws

	s.ws.RegisterMessageHandler(websocket.MessageType(pb.MessageType_PersonalChat),
		func(session *websocket.Session, payload websocket.MessagePayload) error {
			switch t := payload.(type) {
			case *pb.PersonalChatMessage:
				return s.OnPersonalChatMessage(session, t)
			default:
				return pbErrors.ErrorInvalidPayloadType("the payload type is not personal-chat")
			}
		},
		func() websocket.Any { return &pb.PersonalChatMessage{} },
	)
	s.ws.RegisterMessageHandler(websocket.MessageType(pb.MessageType_GroupChat),
		func(session *websocket.Session, payload websocket.MessagePayload) error {
			switch t := payload.(type) {
			case *pb.GroupChatMessage:
				return s.OnGroupChatMessage(session, t)
			default:
				return pbErrors.ErrorInvalidPayloadType("the payload type is not group-chat")
			}
		},
		func() websocket.Any { return &pb.GroupChatMessage{} },
	)
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
		deviceIDStr := session.Request().FormValue("d")
		deviceID, err := uuid.Parse(deviceIDStr)
		if err != nil {
			s.log.Errorf("parse device id error: %v, device id: %s", err, deviceIDStr)
			return
		}
		s.HandleUnSynchronizedMessage(uid, deviceID)
	}
}

func (s *Service) HandleUnSynchronizedMessage(uid int, deviceID uuid.UUID) {
	ctx := context.Background()
	s.log.Infof("load message user id: %d, deviceID: %s", uid, deviceID.String())
	lastTime, err := s.src.GetLastTime(ctx, uid, deviceID)
	if err != nil {
		s.log.Errorf("get synchronize last time error: %v", err)
		return
	}
	messages, err := s.mc.GetChatMessages(ctx, uid, *lastTime)
	if err == nil {
		for _, message := range messages {
			s.log.Infof("send message user id: %d, payload: %+v", uid, message.Payload)
			s.ws.SendMessageByBID(uid, websocket.MessageType(message.MessageType), message.Payload)
		}
	}
	_, err = s.src.SaveSynchronizeRecord(ctx, uid, deviceID)
	if err != nil {
		s.log.Errorf("save synchronize record error: %s", err)
	}
}

func (s *Service) SendMessage(uid int, messageType pb.MessageType, payload interface{}) {
	sessions, ok := s.ws.GetSessionsByBusinessID(uid)
	if !ok {
		s.log.Warnf("no active sessions found, uid: %d", uid)
		return
	}

	for _, session := range sessions {
		s.ws.SendMessage(session.SessionID(), websocket.MessageType(messageType), payload)
	}
}
