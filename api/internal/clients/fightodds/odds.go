package fightodds

type FightOdds struct {
	Fighter1 Fighter
	Fighter2 Fighter
}

type Fighter struct {
	ID        string
	FirstName string
	LastName  string
	BestOdds  int
}
