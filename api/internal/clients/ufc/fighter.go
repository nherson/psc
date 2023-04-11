package ufc

type Fighter struct {
	FighterID int     `json:"FighterId"`
	MMAID     int     `json:"MMAId"`
	Name      Name    `json:"Name"`
	Outcome   Outcome `json:"Outcome"`
}

type Name struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	NickName  string `json:"NickName"`
}

type Outcome struct {
	OutcomeID int    `json:"OutcomeId"`
	Outcome   string `json:"Outcome"`
}
