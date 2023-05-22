package service

import (
	"context"

	pb "beetle/api/im/service/v1"
	"beetle/app/im/internal/biz"
	"beetle/internal/pkg/jwt"

	"github.com/Soul-Killer-Ky/kratos/websocket"
	"github.com/go-kratos/kratos/v2/log"
)

type ImService struct {
	pb.UnimplementedImServer

	gc  *biz.GroupUseCase
	log *log.Helper

	ws *websocket.Server
}

func NewImService(gc *biz.GroupUseCase, logger log.Logger) *ImService {
	return &ImService{gc: gc, log: log.NewHelper(logger)}
}

func (s *ImService) OnWebsocketConnect(sessionID websocket.SessionID, conn bool) {
	s.log.Infof("on websocket conn, session id: %s, conn: %v", sessionID, conn)
}

func (s *ImService) SetWebsocketServer(ws *websocket.Server) {
	s.ws = ws
}

func (s *ImService) CreateIm(ctx context.Context, req *pb.CreateImRequest) (*pb.CreateImReply, error) {
	return &pb.CreateImReply{}, nil
}
func (s *ImService) UpdateIm(ctx context.Context, req *pb.UpdateImRequest) (*pb.UpdateImReply, error) {
	return &pb.UpdateImReply{}, nil
}
func (s *ImService) DeleteIm(ctx context.Context, req *pb.DeleteImRequest) (*pb.DeleteImReply, error) {
	return &pb.DeleteImReply{}, nil
}
func (s *ImService) GetIm(ctx context.Context, req *pb.GetImRequest) (*pb.GetImReply, error) {
	return &pb.GetImReply{}, nil
}
func (s *ImService) ListIm(ctx context.Context, req *pb.ListImRequest) (*pb.ListImReply, error) {
	return &pb.ListImReply{}, nil
}
func (s *ImService) ConnectIm(ctx context.Context, req *pb.ConnectImRequest) (*pb.ConnectImReply, error) {
	//tr, _ := transport.FromServerContext(ctx)
	//htr := tr.(http.Transporter)
	//request := htr.Request()
	//http.RequestFromServerContext()
	//response := request.Response
	//s.log.Infof("ctx: %v %v", request, response)
	//s.ic.Connect(ctx, request, response.(*http.ResponseWriter))

	return &pb.ConnectImReply{}, nil
}

func (s *ImService) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.GetGroupReply, error) {
	userID, err := jwt.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	ps, err := s.gc.GetGroups(ctx, userID)
	groups := make([]*pb.GetGroupReply_Group, 0)
	for _, p := range ps {
		groups = append(groups, &pb.GetGroupReply_Group{
			Id:   int64(p.ID),
			Name: p.Name,
			Icon: p.Icon,
			Memo: p.Memo,
		})
	}
	return &pb.GetGroupReply{
		Groups: groups,
	}, nil
}
