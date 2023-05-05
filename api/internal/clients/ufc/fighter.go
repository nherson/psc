package ufc

type Fighter struct {
	FighterID int  `json:"FighterId"`
	MMAID     int  `json:"MMAId"`
	Name      Name `json:"Name"`

	// Relevant in fight outcomes
	Outcome Outcome `json:"Outcome"`
	Corner  string  `json:"Corner"`
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
