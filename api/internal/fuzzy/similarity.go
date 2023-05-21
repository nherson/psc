package fuzzy

import "github.com/nherson/psc/api/ent"

type similarityOutput struct {
	dbFighter *ent.Fighter
	fullName  string
	score     float64
}

type similarities []similarityOutput

func (s similarities) Len() int           { return len(s) }
func (s similarities) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s similarities) Less(i, j int) bool { return s[i].score > s[j].score }
