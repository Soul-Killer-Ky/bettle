package server

import (
	pb "beetle/api/im/service/v1"
	"beetle/app/im/internal/conf"
	"beetle/app/im/internal/service"
	jwt2 "beetle/internal/pkg/jwt"
	"context"
	"errors"

	"github.com/Soul-Killer-Ky/kratos/websocket"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

// NewWebsocketServer create a websocket server.
func NewWebsocketServer(c *conf.Server, ac *conf.Auth, logger log.Logger, svc *service.ImService) *websocket.Server {
	srv := websocket.NewServer(
		websocket.WithAddress(c.Websocket.Addr),
		websocket.WithPath(c.Websocket.Path),
		websocket.WithConnectHandle(svc.OnWebsocketConnect),
		websocket.WithCodec("json"),
		websocket.WithMiddleware(http.Middleware(
			recovery.Recovery(),
			selector.Server(
				jwt.Server(
					func(token *jwtv4.Token) (interface{}, error) {
						return []byte(ac.JwtKey), nil
					},
					jwt.WithSigningMethod(jwtv4.SigningMethodHS256),
					jwt.WithClaims(func() jwtv4.Claims {
						return &jwt2.CustomUserClaims{}
					}),
				),
			).Match(NewWhiteListMatcher()).Build(),
		)),
		websocket.WithBusinessIDFunc(func(ctx context.Context) interface{} {
			token, ok := jwt.FromContext(ctx)
			if !ok {
				return nil
			}
			claims := token.(*jwt2.CustomUserClaims)
			return uint64(claims.ID)
		}),
	)

	svc.SetWebsocketServer(srv)

	srv.RegisterMessageHandler(websocket.MessageType(pb.ProtocolType_Chat),
		func(session *websocket.Session, payload websocket.MessagePayload) error {
			switch t := payload.(type) {
			case *pb.ChatMessage:
				t.From = session.BusinessID().(uint64)
				return svc.OnChatMessage(session, t)
			default:
				return errors.New("invalid payload type")
			}
		},
		func() websocket.Any { return &pb.ChatMessage{} },
	)
	srv.RegisterMessageHandler(websocket.MessageType(pb.ProtocolType_Group),
		func(session *websocket.Session, payload websocket.MessagePayload) error {
			switch t := payload.(type) {
			case *pb.GroupMessage:
				return svc.OnGroupMessage(session, t)
			default:
				return errors.New("invalid payload type")
			}
		},
		func() websocket.Any { return &pb.GroupMessage{} },
	)

	return srv
}
