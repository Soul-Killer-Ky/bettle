package biz

import (
	"context"

	"beetle/app/filesystem/internal/pkg/storage"

	"github.com/go-kratos/kratos/v2/log"
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

type FilesystemRepo interface {
	GetGreeter(context.Context, int) (*Greeter, error)
}

type FilesystemUseCase struct {
	repo FilesystemRepo

	log *log.Helper

	storage storage.Storage
}

func NewFilesystemUseCase(repo FilesystemRepo, logger log.Logger) *FilesystemUseCase {
	return &FilesystemUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c *FilesystemUseCase) Upload() error {
	return nil
	//name := "some-unique-name.png"
	//file := storage.NewFile(name)

	//for {
	//	req, err := stream.Recv()
	//	if err == io.EOF {
	//		if err := c.storage.Store(file); err != nil {
	//			return status.Error(codes.Internal, err.Error())
	//		}
	//
	//		return stream.SendAndClose(&pb.UploadReply{Url: name})
	//	}
	//	if err != nil {
	//		return status.Error(codes.Internal, err.Error())
	//	}
	//
	//	if err := file.Write(req.GetChunk()); err != nil {
	//		return status.Error(codes.Internal, err.Error())
	//	}
	//}
}
