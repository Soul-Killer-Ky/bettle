package ws

import (
	"context"
	"time"

	pbErrors "beetle/api/beetle/errors"
	pb "beetle/api/im/service/v1"
	"beetle/app/im/internal/biz"
	"beetle/internal/pkg/util"

	"github.com/Soul-Killer-Ky/kratos/websocket"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	pb.UnimplementedImServer

	mc  *biz.MessageUseCase
	src *biz.SynchronizeRecordUseCase
	uc  *biz.UserUseCase
	log *log.Helper

	ws *websocket.Server

	lastMessageID int64
}

func NewService(mc *biz.MessageUseCase, src *biz.SynchronizeRecordUseCase, uc *biz.UserUseCase, logger log.Logger) *Service {
	return &Service{mc: mc, src: src, uc: uc, log: log.NewHelper(logger)}
}

func (s *Service) SetWebsocketServer(ws *websocket.Server) {
	s.ws = ws

	s.ws.RegisterMessageHandler(
		websocket.MessageType(pb.MessageType_PersonalChat),
		s.OnMessage,
		func() websocket.Any {
			return &pb.PersonalChatMessage{}
		},
	)
	s.ws.RegisterMessageHandler(
		websocket.MessageType(pb.MessageType_GroupChat),
		s.OnMessage,
		func() websocket.Any {
			return &pb.GroupChatMessage{}
		},
	)
}

func (s *Service) OnMessage(session *websocket.Session, payload websocket.MessagePayload) error {
	msgID, err := util.MakeSnowflakeID()
	if err != nil {
		return pbErrors.ErrorMessageIdGenerateFailure("generate fail")
	}
	// 记录最后分配的消息ID
	s.lastMessageID = msgID
	switch t := payload.(type) {
	case *pb.PersonalChatMessage:
		t.MessageId = msgID
		t.Timestamp = timestamppb.New(time.Now())
		return s.OnPersonalChatMessage(session, t)
	case *pb.GroupChatMessage:
		t.MessageId = msgID
		t.Timestamp = timestamppb.New(time.Now())
		return s.OnGroupChatMessage(session, t)
	default:
		return pbErrors.ErrorInvalidPayloadType("the payload type is invalid")
	}
}

func (s *Service) OnWebsocketConnect(sessionID websocket.SessionID, connect bool) {
	s.log.Infof("on websocket conn, session id: %s, conn: %v", sessionID, connect)
	ctx := context.Background()
	session, ok := s.ws.GetSession(sessionID)
	if !ok {
		s.log.Errorf("not found session, session id: %s", sessionID)
		return
	}
	deviceIDStr := session.Request().FormValue("d")
	deviceID, err := uuid.Parse(deviceIDStr)
	if err != nil {
		s.log.Errorf("parse device id error: %v, device id: %s", err, deviceIDStr)
		return
	}
	uid := session.BusinessID().(int)
	if connect {
		//session, ok := s.ws.GetSession(sessionID)
		//if !ok {
		//	s.log.Errorf("not found session, session id: %s", sessionID)
		//	return
		//}
		//uid := session.BusinessID().(int)
		//deviceIDStr := session.Request().FormValue("d")
		//deviceID, err := uuid.Parse(deviceIDStr)
		//if err != nil {
		//	s.log.Errorf("parse device id error: %v, device id: %s", err, deviceIDStr)
		//	return
		//}
		s.HandleOfflineMessage(ctx, uid, deviceID, session)
	} else {
		// 离线时记录最后分配的消息ID
		if s.lastMessageID > 0 {
			_, err = s.src.SaveSynchronizeRecord(ctx, uid, deviceID, s.lastMessageID)
			if err != nil {
				s.log.Errorf("save synchronize record is error on user offline: %s", err)
			}
		}
	}
}

func (s *Service) HandleOfflineMessage(ctx context.Context, uid int, deviceID uuid.UUID, session *websocket.Session) {
	s.log.Infof("load offline message by user id: %d, deviceID: %s", uid, deviceID.String())
	messageID, err := s.src.GetLastMessageID(ctx, uid, deviceID)
	if err != nil {
		s.log.Errorf("get synchronize last message id error: %v", err)
		return
	}

	var lastMessageID int64
	pms, err := s.mc.GetPersonalChatMessages(ctx, uid, messageID)
	if err != nil {
		s.log.Errorf("get personal-chat message error: %v", err)
		return
	}
	for _, pm := range pms {
		if pm.MessageId > lastMessageID {
			lastMessageID = pm.MessageId
		}
		s.log.Infof("send message to user id: %d, payload: %+v", uid, pm)
		s.ws.SendMessage(session.SessionID(), websocket.MessageType(pb.MessageType_PersonalChat), pm)
	}

	groups, err := s.uc.ListGroup(ctx)
	if err != nil {
		s.log.Errorf("list group error: %v", err)
		return
	}
	var groupIDs []int
	for _, group := range groups {
		groupIDs = append(groupIDs, group.ID)
	}

	gms, err := s.mc.GetGroupChatMessages(ctx, groupIDs, messageID)
	if err == nil {
		for _, gm := range gms {
			if gm.MessageId > lastMessageID {
				lastMessageID = gm.MessageId
			}
			s.log.Infof("send message to user id: %d, payload: %+v", uid, gm)
			s.ws.SendMessage(session.SessionID(), websocket.MessageType(pb.MessageType_GroupChat), gm)
		}
	}

	if lastMessageID > 0 {
		_, err = s.src.SaveSynchronizeRecord(ctx, uid, deviceID, lastMessageID)
		if err != nil {
			s.log.Errorf("save synchronize record error: %s", err)
		}
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
