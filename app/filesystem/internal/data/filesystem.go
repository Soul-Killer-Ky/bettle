package data

import (
	"beetle/app/filesystem/internal/biz"
	"beetle/app/filesystem/internal/data/ent/storage"
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

func (r *filesystemRepo) FindStorageByMd5(ctx context.Context, hash string) (*biz.Storage, error) {
	p, err := r.data.db.Storage.Query().Where(storage.Md5(hash)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Storage{
		Type: p.Type,
		Path: p.Path,
		Md5:  p.Md5,
	}, err
}

func (r *filesystemRepo) CreateStorage(ctx context.Context, typ int32, path, hash string) (*biz.Storage, error) {
	p, err := r.data.db.Storage.Create().
		SetType(typ).
		SetPath(path).
		SetMd5(hash).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Storage{
		Type: p.Type,
		Path: p.Path,
		Md5:  p.Md5,
	}, err
}

func (r *filesystemRepo) CreateQuote(ctx context.Context, hash, name string, createdBy int) (*biz.Quote, error) {
	p, err := r.data.db.Storage.Query().Where(storage.Md5(hash)).Only(ctx)
	if err != nil {
		return nil, err
	}
	p2, err := r.data.db.Quote.Create().
		SetStorage(p).
		SetName(name).
		SetCreatedBy(createdBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Quote{
		Name:      p2.Name,
		CreatedBy: p2.CreatedBy,
	}, err
}
