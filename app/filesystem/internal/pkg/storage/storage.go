package storage

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"beetle/internal/pkg/util"
)

type Manager interface {
	Store(file *File) error
}

var _ Manager = &Storage{}

type Storage struct {
	baseDir string
	subDir  string
	dir     string
}

func New(dir string) Storage {
	return Storage{
		baseDir: dir,
		dir:     dir,
	}
}

func (s Storage) SetSubDir(subDirs ...string) Storage {
	s.subDir = filepath.Join(subDirs...)
	s.dir = filepath.Join(s.baseDir, s.subDir)
	fd, err := os.Stat(s.dir)
	if err != nil {
		err := os.MkdirAll(s.dir, 0o666)
		if err != nil {
			panic(err)
		}
	} else if !fd.IsDir() {
		panic("the dir is file")
	}
	return s
}

func (s Storage) Store(file *File) error {
	if err := ioutil.WriteFile(s.dir+file.name, file.buffer.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

func (s Storage) Store1(body []byte) (string, error) {
	id, err := util.MakeSnowflakeID()
	if err != nil {
		return "", err
	}
	path := strconv.FormatInt(id, 10)
	fullPath := filepath.Join(s.dir, path)
	if err := ioutil.WriteFile(fullPath, body, 0644); err != nil {
		return "", err
	}

	return path, nil
}

func (s Storage) Store2(file io.Reader, ext string) (string, error) {
	id, err := util.MakeSnowflakeID()
	if err != nil {
		return "", err
	}
	filename := strconv.FormatInt(id, 10) + ext
	fullPath := filepath.Join(s.dir, filename)
	f, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	return filepath.Join(s.subDir, filename), nil
}
