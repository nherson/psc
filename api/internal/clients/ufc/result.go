package ufc

type Result struct {
	Method      string `json:"Method"`
	EndingRound int    `json:"EndingRound"`

	// EndingTime is the amount of time that elapsed in the round, NOT the time left on the clock
	EndingTime string `json:"EndingTime"`
}
