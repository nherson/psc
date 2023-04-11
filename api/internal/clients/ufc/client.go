package ufc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) EventByID(id int) (Event, error) {
	var payload EventPayload

	resp, err := c.httpClient.Get(c.eventURL(id))
	if err != nil {
		return Event{}, errors.Wrap(err, "could not fetch event details from ufc api")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Event{}, errors.Wrap(err, "could not read response body")
	}

	err = json.Unmarshal(b, &payload)
	if err != nil {
		return Event{}, errors.Wrap(err, "could not unmarshal ufc event data")
	}

	return payload.Event, nil
}

func (c *Client) FightByID(id int) (Fight, error) {
	var payload FightPayload

	resp, err := c.httpClient.Get(c.fightURL(id))
	if err != nil {
		return Fight{}, errors.Wrap(err, "could not fetch fight details from ufc api")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Fight{}, errors.Wrap(err, "could not read response body")
	}

	err = json.Unmarshal(b, &payload)
	if err != nil {
		return Fight{}, errors.Wrap(err, "could not unmarshal ufc fight data")
	}

	return payload.Fight, nil
}

func (c *Client) eventURL(id int) string {
	return fmt.Sprintf("%s/%d.json", eventBaseURL, id)
}

func (c *Client) fightURL(id int) string {
	return fmt.Sprintf("%s/%d.json", fightBaseURL, id)
}
