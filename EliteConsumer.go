package goNurseSchedulingSolver

import (
	ga "github.com/tomcraven/goga"
)

type EliteConsumer struct{}

func (EliteConsumer) OnElite(genome *ga.IGenome) {
	panic("implement me")
}
