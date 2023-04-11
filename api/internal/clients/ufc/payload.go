package ufc

type EventPayload struct {
	Event Event `json:"LiveEventDetail"`
}

type FightPayload struct {
	Fight Fight `json:"LiveFightDetail"`
}
