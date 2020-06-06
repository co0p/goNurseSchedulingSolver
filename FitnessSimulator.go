package goNurseSchedulingSolver

import ga "github.com/tomcraven/goga"

type FitnessSimulator struct{}

func (s *FitnessSimulator) OnBeginSimulation() {
	panic("implement me")
}

func (s *FitnessSimulator) Simulate(genome *ga.IGenome) {
	panic("implement me")
}

func (s *FitnessSimulator) OnEndSimulation() {
	panic("implement me")
}

func (s *FitnessSimulator) ExitFunc(genome *ga.IGenome) bool {
	panic("implement me")
}
