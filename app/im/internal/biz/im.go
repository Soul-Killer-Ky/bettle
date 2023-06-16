package biz

import (
	"io"

	pb "beetle/api/im/service/v1"
	"beetle/app/im/internal/pkg/storage"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ImUseCase struct {
	log *log.Helper

	storage storage.Storage
}

func NewImUseCase(logger log.Logger) *ImUseCase {
	return &ImUseCase{log: log.NewHelper(logger)}
}

func (c *ImUseCase) Upload(stream pb.Im_UploadServer) error {
	name := "some-unique-name.png"
	file := storage.NewFile(name)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if err := c.storage.Store(file); err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			return stream.SendAndClose(&pb.UploadReply{Url: name})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if err := file.Write(req.GetChunk()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
}
