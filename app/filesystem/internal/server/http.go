package server

import (
	pb "beetle/api/filesystem/v1"
	"beetle/app/filesystem/internal/conf"
	"beetle/app/filesystem/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"strconv"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, filesystem *service.FilesystemService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	route := srv.Route("/")
	route.POST("/v1/upload", filesystemUploadHttpHandler(filesystem))

	//pb.RegisterFilesystemHTTPServer(srv, filesystem)
	return srv
}

func filesystemUploadHttpHandler(srv pb.FilesystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in pb.UploadRequest
		req := ctx.Request()
		file, handler, err := req.FormFile("file")
		if err != nil {
			return err
		}
		defer file.Close()
		file.Read(in.File)
		in.Name = handler.Filename
		fileType, err := strconv.ParseInt(req.FormValue("type"), 10, 32)
		if err != nil {
			return err
		}
		in.Type = pb.UploadType(int32(fileType))

		//f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0o666)
		//if err != nil {
		//	return err
		//}
		//defer f.Close()
		//_, _ = io.Copy(f, file)
		//
		//if err := bind(ctx, &in); err != nil {
		//	return err
		//}
		http.SetOperation(ctx, pb.OperationFilesystemUpload)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Upload(ctx, req.(*pb.UploadRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*pb.UploadReply)
		return ctx.Result(200, reply)
	}
}
