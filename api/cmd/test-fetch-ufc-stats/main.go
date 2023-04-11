package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/nherson/psc/api/internal/clients/ufc"
)

func main() {
	c := ufc.NewCacheClient()

	stats, err := c.EventByID(1135)
	if err != nil {
		panic(err)
	}

	spew.Dump(stats)

	// fightStats, err := c.FightByID(6458)
	// if err != nil {
	// 	panic(err)
	// }

	// spew.Dump(fightStats)

}
