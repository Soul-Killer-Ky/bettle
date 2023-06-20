package biz

import (
	"beetle/api/beetle/errors"
	pb "beetle/api/filesystem/v1"
	"context"
	"io"
	"mime/multipart"
	"path/filepath"

	"beetle/app/filesystem/internal/data/ent"
	"beetle/app/filesystem/internal/pkg/storage"
	"beetle/internal/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

const ImageSizeLimit = 2 * 1024 * 1024
const FileSizeLimit = 50 * 1024 * 1024

type Storage struct {
	Type int32
	Path string
	Md5  string
}

type Quote struct {
	Name      string
	CreatedBy int
}

type File struct {
	Type int32
	Path string
	Md5  string
	Name string
}

type FilesystemRepo interface {
	FindStorageByMd5(context.Context, string) (*Storage, error)
	CreateStorage(context.Context, int32, string, string) (*Storage, error)
	CreateQuote(context.Context, string, string, int) (*Quote, error)
}

type FilesystemUseCase struct {
	repo FilesystemRepo

	log *log.Helper

	storage storage.Storage
}

func NewFilesystemUseCase(repo FilesystemRepo, logger log.Logger) *FilesystemUseCase {
	return &FilesystemUseCase{
		repo:    repo,
		log:     log.NewHelper(logger),
		storage: storage.New("/var/lib/filesystem"),
	}
}

func (c *FilesystemUseCase) CheckFile(ctx context.Context, header *multipart.FileHeader) error {
	c.log.Infof("header: +v", header.Header)

	contentType := header.Header.Get("Content-Type")
	uploadType := util.GetUploadType(contentType)

	if uploadType == pb.UploadType_UT_Image {
		if header.Size > ImageSizeLimit {
			return errors.ErrorUploadFileTooLarge("image size over %dMB", 2)
		}
	}

	if uploadType == pb.UploadType_UT_File {
		if header.Size > FileSizeLimit {
			return errors.ErrorUploadFileTooLarge("file size over %dMB", 50)
		}
	}

	return nil
}

func (c *FilesystemUseCase) Store(ctx context.Context, file multipart.File, header *multipart.FileHeader, userID int) (*File, error) {
	hash, err := util.Md5Reader(file)
	if err != nil {
		return nil, err
	}
	p, err := c.repo.FindStorageByMd5(ctx, hash)
	if err != nil {
		if ent.IsNotFound(err) {
			contentType := header.Header.Get("Content-Type")
			uploadType := util.GetUploadType(contentType)
			_, err := file.Seek(0, io.SeekStart)
			if err != nil {
				return nil, err
			}
			path, err := c.storage.
				SetSubDir(util.GetSubPath(uploadType)).
				Store2(file, filepath.Ext(header.Filename))
			if err != nil {
				return nil, err
			}
			p, err = c.repo.CreateStorage(ctx, int32(uploadType), path, hash)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	p2, err := c.repo.CreateQuote(ctx, hash, header.Filename, userID)

	return &File{
		Type: p.Type,
		Path: p.Path,
		Md5:  p.Md5,
		Name: p2.Name,
	}, nil
}
