package ufc

type Event struct {
	EventID   int       `json:"EventId"`
	Name      string    `json:"Name"`
	FightCard FightCard `json:"FightCard"`
}
