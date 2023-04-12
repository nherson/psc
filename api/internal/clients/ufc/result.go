package ufc

type Result struct {
	Method      string `json:"Method"`
	EndingRound int    `json:"EndingRound"`

	// EndingTime is the amount of time that elapsed in the round, NOT the time left on the clock
	EndingTime string `json:"EndingTime"`
}

// EndingTimeSeconds turns a time string like '3:40' into an integer number of seconds like 220
func (r *Result) EndingTimeSeconds() int {
	return timeStringToSeconds(r.EndingTime)
}
