package ufc

type FightStats []FighterStats

type FighterStats struct {
	FighterID        int    `json:"FighterId"`
	SigStrikesLanded int    `json:"SigStrikesLanded"`
	Knockdowns       int    `json:"Knockdowns"`
	TakedownsLanded  int    `json:"TakedownsLanded"`
	ControlTime      string `json:"ControlTime"`
}

func (f *FighterStats) ControlTimeSeconds() int {
	return timeStringToSeconds(f.ControlTime)
}
