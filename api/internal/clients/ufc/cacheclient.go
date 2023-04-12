package ufc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
)

// CacheClient is like a Client, but checks the `data` directory for a matching event or fight first
type CacheClient struct {
	client *Client
}

func NewCacheClient() *CacheClient {
	return &CacheClient{
		client: NewClient(),
	}
}

func (c *CacheClient) EventByID(id int) (Event, error) {
	var payload EventPayload
	var event Event
	path := fmt.Sprintf("%s/final/events/event-%d.json", cacheDataDirectory, id)
	b, err := loadFileIfExists(path)
	if err != nil {
		// This probably shouldn't happen
		log.Println("unexpected error trying to read cached event detail", err)
	}

	if b != nil {
		log.Printf("found cache entry for event %d\n", id)
		err = json.Unmarshal(b, &payload)
		event = payload.Event
	} else {
		log.Printf("cache miss for fight %d\n", id)
		event, err = c.client.EventByID(id)
		cacheErr := c.cacheResponse(path, c.client.eventURL(id))
		if cacheErr != nil {
			log.Println("could not cache event data", err)
		}
	}

	return event, err
}

func (c *CacheClient) FightByID(id int) (Fight, error) {
	var payload FightPayload
	var fight Fight
	path := fmt.Sprintf("%s/final/fights/fight-%d.json", cacheDataDirectory, id)
	b, err := loadFileIfExists(path)
	if err != nil {
		// This probably shouldn't happen
		log.Println("unexpected error trying to read cached fight detail", err)
	}

	if b != nil {
		log.Printf("found cache entry for fight %d\n", id)
		err = json.Unmarshal(b, &payload)
		fight = payload.Fight
	} else {
		log.Printf("cache miss for fight %d\n", id)
		fight, err = c.client.FightByID(id)
		cacheErr := c.cacheResponse(path, c.client.fightURL(id))
		if cacheErr != nil {
			log.Println("could not cache fight data", err)
		}
	}

	return fight, err
}

func (c *CacheClient) cacheResponse(path, url string) error {
	resp, err := c.client.httpClient.Get(url)
	if err != nil {
		return errors.Wrap(err, "could not fetch url for caching")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "could not read response body for caching")
	}

	return errors.Wrap(os.WriteFile(path, b, 0666), "could not write response body for caching")
}

// best effort attempt to load the file from disk
func loadFileIfExists(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return b, nil
}
