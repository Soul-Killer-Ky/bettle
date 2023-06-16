package data

import (
	"beetle/app/filesystem/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type filesystemRepo struct {
	data *Data
	log  *log.Helper
}

func NewFilesystemRepo(data *Data, logger log.Logger) biz.FilesystemRepo {
	return &filesystemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *filesystemRepo) GetGreeter(context.Context, int) (*biz.Greeter, error) {
	return nil, nil
}
