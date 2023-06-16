package service

import (
	"context"

	pb "beetle/api/filesystem/v1"
	"beetle/app/filesystem/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type FilesystemService struct {
	pb.UnimplementedFilesystemServer

	fc *biz.FilesystemUseCase

	log *log.Helper
}

func NewFilesystemService(fc *biz.FilesystemUseCase, logger log.Logger) *FilesystemService {
	return &FilesystemService{fc: fc, log: log.NewHelper(logger)}
}

func (s *FilesystemService) Upload(ctx context.Context, req *pb.UploadRequest) (*pb.UploadReply, error) {
	s.log.Infof("name: %v", req.Name)
	s.log.Infof("file: %v", req.File)
	return &pb.UploadReply{}, nil
}
