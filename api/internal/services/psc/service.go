package psc

import (
	"context"

	"github.com/bufbuild/connect-go"

	"github.com/nherson/psc/api/ent"
	apiv1 "github.com/nherson/psc/api/proto/api/v1"
)

type PSCServer struct {
	DB *ent.Client
}

func (s *PSCServer) ListEvents(ctx context.Context, req *connect.Request[apiv1.ListEventsRequest]) (*connect.Response[apiv1.ListEventsResponse], error) {
	events, err := s.DB.Event.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var apiEvents []*apiv1.Event
	for _, e := range events {
		apiEvents = append(apiEvents, dbEventToApi(e))
	}

	res := connect.NewResponse(&apiv1.ListEventsResponse{
		Events: apiEvents,
	})

	return res, nil
}
