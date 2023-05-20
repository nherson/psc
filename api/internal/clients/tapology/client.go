package tapology

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/corpix/uarand"
	colly "github.com/gocolly/colly/v2"
	"github.com/pkg/errors"
)

var dateRegexp = regexp.MustCompile("^.* ([0-9][0-9].[0-9][0-9].[0-9][0-9][0-9][0-9]) .*$")

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) UpcomingEvents(ctx context.Context, until time.Time) ([]Event, error) {
	eventListScraper := colly.NewCollector(colly.UserAgent(uarand.GetRandom()))
	eventScraper := colly.NewCollector(colly.UserAgent(uarand.GetRandom()))

	eventListScraper.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	eventListScraper.OnHTML("section.fcListing", func(e *colly.HTMLElement) {
		url := e.ChildAttr("span.name a", "href")
		eventScraper.Visit("https://www.tapology.com" + url)
	})

	var events []Event
	eventScraper.OnHTML("div#content", func(e *colly.HTMLElement) {
		// Timestamp, name, and ID of event
		t, err := c.parseEventTime(e)
		if err != nil {
			fmt.Println("Error parsing time:", err.Error())
			return
		}
		if !t.Before(until) {
			return
		}
		fmt.Println("Got event date time:", t)

		urlStringParts := strings.Split(e.Request.URL.String(), "/")
		id := urlStringParts[len(urlStringParts)-1]
		name := e.ChildText("div.eventPageHeaderTitles h1")

		// Parse the fights of the event
		var fights []Fight
		e.ForEach("div.fightCardBout", func(i int, e *colly.HTMLElement) {
			// Filter out the weird "predictions" section at the bottom of the page
			if !isFightElement(e) {
				return
			}

			f, err := c.parseFightCardBout(i, e)
			if err != nil {
				fmt.Printf("Error parsing fight information at index %d: %s\n", i, err.Error())
			} else {
				fights = append(fights, f)
			}
		})

		events = append(events, Event{
			ID:     id,
			Name:   name,
			Date:   t,
			Fights: fights,
		})
	})

	eventListScraper.OnRequest(func(r *colly.Request) {
		fmt.Println("Getting event", r.URL.String())
	})

	err := eventListScraper.Visit("https://www.tapology.com/fightcenter?group=ufc")
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (c *Client) parseFightCardBout(i int, e *colly.HTMLElement) (Fight, error) {
	var fighters []Fighter
	e.ForEach("div.fightCardFighterName", func(_ int, e *colly.HTMLElement) {
		fighters = append(fighters, fighterFromFightCardBoutFighter(e))
	})
	if len(fighters) != 2 {
		return Fight{}, errors.New("did not find 2 fighters in .fightCardBout")
	}

	return Fight{
		Fighter1:   fighters[0],
		Fighter2:   fighters[1],
		BoutNumber: i + 1, // 1-index for consistency with UFC canonical data
	}, nil
}

func (c *Client) parseEventTime(e *colly.HTMLElement) (time.Time, error) {
	dateString := e.ChildText("div.details div.right li.header")

	matches := dateRegexp.FindStringSubmatch(dateString)
	if len(matches) != 2 {
		return time.Time{}, fmt.Errorf("bad regex parse of event date string %q", dateString)
	}

	t, err := time.Parse("01.02.2006", matches[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("bad time parse of %q", dateString)
	}

	return t, nil
}

func isFightElement(e *colly.HTMLElement) bool {
	classAttrs := strings.Split(e.Attr("class"), " ")
	for _, attr := range classAttrs {
		if attr == "picks" {
			return false
		}
	}
	return true
}

func fighterFromFightCardBoutFighter(e *colly.HTMLElement) Fighter {
	var firstName, lastName string
	fullName := strings.TrimSpace(e.Text)
	parts := strings.Split(fullName, " ")
	if len(parts) == 1 {
		firstName, lastName = "", parts[0]
	} else {
		firstName, lastName = parts[0], strings.Join(parts[1:], " ")
	}

	href := e.ChildAttr("div.fightCardFighterName a", "href")
	parts = strings.Split(href, "/")
	id := parts[len(parts)-1]

	return Fighter{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
	}
}
