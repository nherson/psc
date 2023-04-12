package psc

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"

	apiv1 "github.com/nherson/psc/api/proto/api/v1"
)

type PSCServer struct{}

func (s *PSCServer) Hello(ctx context.Context, req *connect.Request[apiv1.HelloRequest]) (*connect.Response[apiv1.HelloResponse], error) {
	res := connect.NewResponse(&apiv1.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})

	return res, nil
}
