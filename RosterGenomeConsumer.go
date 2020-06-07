package goNurseSchedulingSolver

import (
	ga "github.com/tomcraven/goga"
	"log"
)

type RosterGenomeConsumer struct {
	iteration int
}

func (e *RosterGenomeConsumer) OnElite(g *ga.IGenome) {
	roster := NewRosterFromGenome(g)
	log.Printf("[%d] = %d - %s", e.iteration, roster.GetFitness(), roster)
	e.iteration++
}
