package service

import (
	"context"
	"mime/multipart"

	pb "beetle/api/filesystem/v1"
	"beetle/app/filesystem/internal/biz"
	"beetle/internal/pkg/jwt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type FilesystemService struct {
	pb.UnimplementedFilesystemServer

	fc *biz.FilesystemUseCase

	log *log.Helper
}

func NewFilesystemService(fc *biz.FilesystemUseCase, logger log.Logger) *FilesystemService {
	return &FilesystemService{fc: fc, log: log.NewHelper(logger)}
}

type UploadRequest struct {
	File   multipart.File
	Header *multipart.FileHeader
}

func (s *FilesystemService) Upload(ctx context.Context, ctx2 http.Context) (*pb.UploadReply, error) {
	req := ctx2.Request()
	f, header, err := req.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	err = s.fc.CheckFile(ctx, header)
	if err != nil {
		return nil, err
	}

	userID, err := jwt.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	file, err := s.fc.Store(ctx, f, header, userID)
	if err != nil {
		return nil, err
	}
	return &pb.UploadReply{
		Url: "http://localhost:9100/" + file.Path,
	}, nil
}
