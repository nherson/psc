package ufc

type Fight struct {
	FightID     int        `json:"FightId"`
	FightOrder  int        `json:"FightOrder"`
	CardSegment string     `json:"CardSegment"`
	Fighters    []Fighter  `json:"Fighters"`
	Result      *Result    `json:"Result"`
	FightStats  FightStats `json:"FightStats"`
}
