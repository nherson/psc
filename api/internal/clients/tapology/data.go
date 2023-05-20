package tapology

import "time"

type Event struct {
	ID     string
	Name   string
	Date   time.Time
	Fights []Fight
}

type Fight struct {
	Fighter1   Fighter
	Fighter2   Fighter
	BoutNumber int
}

type Fighter struct {
	ID        string
	FirstName string
	LastName  string
}
