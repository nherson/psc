package psc

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"

	"github.com/nherson/psc/api/ent"
	"github.com/nherson/psc/api/ent/event"
	"github.com/nherson/psc/api/ent/fighter"
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

func (s *PSCServer) ListResultsForEvent(
	ctx context.Context,
	req *connect.Request[apiv1.ListResultsForEventRequest],
) (*connect.Response[apiv1.ListResultsForEventResponse], error) {

	var resp apiv1.ListResultsForEventResponse
	e, err := s.DB.Event.Query().
		Where(
			event.ID(int(req.Msg.GetEventId())),
		).
		WithFights(func(q *ent.FightQuery) {
			q.WithFighterResults(
				func(q *ent.FighterResultsQuery) {
					q.WithFighter()
				},
			)
		}).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, connect.NewError(connect.CodeNotFound, err)
	} else if err != nil {
		return nil, err
	}
	resp.Event = dbEventToApi(e)

	for _, f := range e.Edges.Fights {
		if len(f.Edges.FighterResults) != 2 {
			return nil, fmt.Errorf("bad number of fight results for fight: %d", len(f.Edges.FighterResults))
		}

		resp.FightResults = append(resp.FightResults, dbFightResultsToApi(f, e))
	}

	return connect.NewResponse(&resp), nil
}

func (s *PSCServer) ListResultsForFighter(
	ctx context.Context,
	req *connect.Request[apiv1.ListResultsForFighterRequest],
) (*connect.Response[apiv1.ListResultsForFighterResponse], error) {

	f, err := s.DB.Fighter.Query().
		Where(fighter.ID(int(req.Msg.GetFighterId()))).
		WithFights(func(q *ent.FightQuery) {
			q.WithEvent()
			q.WithFighterResults(
				func(q *ent.FighterResultsQuery) {
					q.WithFighter()
				},
			)
		}).
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, connect.NewError(connect.CodeNotFound, err)
	} else if err != nil {
		return nil, err
	}

	var resp apiv1.ListResultsForFighterResponse
	for _, fight := range f.Edges.Fights {
		resp.FightResults = append(resp.FightResults, dbFightResultsToApi(fight, fight.Edges.Event))
	}

	return connect.NewResponse(&resp), nil
}
