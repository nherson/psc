package ufcstats

import (
	"context"

	v1 "github.com/nherson/psc/api/internal/proto/api/v1"
)

type Service struct{}

func (s *Service) Hello(ctx context.Context, req *v1.HelloReq) (*v1.HelloResp, error) {
	return &v1.HelloResp{
		Text: "Hello, world!",
	}, nil
}
