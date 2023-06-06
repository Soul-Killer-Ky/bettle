package service

import (
	"context"

	pb "beetle/api/im/service/v1"
	"github.com/Soul-Killer-Ky/kratos/websocket"
	"github.com/go-kratos/kratos/v2/log"
)

type ImService struct {
	pb.UnimplementedImServer

	log *log.Helper

	ws *websocket.Server
}

func NewImService(logger log.Logger) *ImService {
	return &ImService{log: log.NewHelper(logger)}
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
