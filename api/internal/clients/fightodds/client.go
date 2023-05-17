package fightodds

import (
	"context"
	"net/http"
	"time"

	"github.com/Khan/genqlient/graphql"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}

type Body struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func (c *Client) UpcomingFighterOdds(ctx context.Context, until time.Time) ([]Fighter, error) {
	untilString := until.Format("2006-01-02")
	nowString := time.Now().Format("2006-01-02")

	gqlClient := graphql.NewClient(fightOddsGraphqlURL, c.httpClient)
	resp, err := EventsQuery(
		ctx,
		gqlClient,
		"ufc",
		untilString,
		nowString,
		"",
		25,
		"date",
	)
	if err != nil {
		return nil, err
	}

	var odds []Fighter

	for _, eventNode := range resp.Promotion.Events.Edges {
		oddsResp, err := OddsQuery(ctx, gqlClient, eventNode.Node.Pk)
		if err != nil {
			return nil, err
		}

		for _, fightOfferNode := range oddsResp.EventOfferTable.FightOffers.Edges {
			f := fightOfferNode.Node
			odds = append(
				odds,
				Fighter{
					ID:        f.Fighter1.Id,
					FirstName: f.Fighter1.FirstName,
					LastName:  f.Fighter1.LastName,
					BestOdds:  f.BestOdds1,
				}, Fighter{
					ID:        f.Fighter2.Id,
					FirstName: f.Fighter2.FirstName,
					LastName:  f.Fighter2.LastName,
					BestOdds:  f.BestOdds2,
				},
			)
		}

	}

	return odds, nil
}
