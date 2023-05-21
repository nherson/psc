package fuzzy

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/nherson/psc/api/ent"
)

type config struct {
	db                  *ent.Client
	promptChoices       int
	confidenceThreshold float32
	stringMetric        strutil.StringMetric
}

type Option func(*config)

func WithDB(db *ent.Client) Option {
	return func(cfg *config) {
		cfg.db = db
	}
}

func WithPromptChoices(n int) Option {
	return func(cfg *config) {
		cfg.promptChoices = n
	}
}

func WithConfidenceThreshold(threshold float32) Option {
	return func(cfg *config) {
		cfg.confidenceThreshold = threshold
	}
}

func WithStringMetric(m strutil.StringMetric) Option {
	return func(cfg *config) {
		cfg.stringMetric = m
	}
}

type Matcher struct {
	db                  *ent.Client
	promptChoices       int
	confidenceThreshold float32
	stringMetric        strutil.StringMetric

	// implement a very primitive cache to avoid excessive db lookups
	fighters                []*ent.Fighter
	fighterRefreshTimestamp time.Time
	lock                    sync.Mutex
}

func NewMatcher(opts ...Option) (*Matcher, error) {
	cfg := &config{
		promptChoices:       5,
		confidenceThreshold: .80,
		stringMetric:        metrics.NewSorensenDice(),
	}
	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.db == nil {
		return nil, errors.New("no db provided")
	}

	return &Matcher{
		db:                  cfg.db,
		promptChoices:       cfg.promptChoices,
		confidenceThreshold: cfg.confidenceThreshold,
		stringMetric:        cfg.stringMetric,
	}, nil
}

func (m *Matcher) getFighters(ctx context.Context) ([]*ent.Fighter, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	now := time.Now()

	if m.fighters == nil || now.After(m.fighterRefreshTimestamp) {
		fighters, err := m.db.Fighter.Query().All(ctx)
		if err != nil {
			return nil, err
		}

		m.fighters = fighters
		m.fighterRefreshTimestamp = now.Add(10 * time.Minute)
	}

	return m.fighters, nil
}
