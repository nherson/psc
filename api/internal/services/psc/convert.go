package psc

import (
	"strconv"

	"github.com/nherson/psc/api/ent"
	apiv1 "github.com/nherson/psc/api/proto/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func dbEventToApi(event *ent.Event) *apiv1.Event {
	return &apiv1.Event{
		Id:   strconv.Itoa(event.ID),
		Name: event.Name,
		Date: timestamppb.New(event.Date),
	}
}
