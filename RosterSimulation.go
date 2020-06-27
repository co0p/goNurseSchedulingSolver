package main

import (
	"fmt"
	ga "github.com/tomcraven/goga"
	"math/rand"
)

type RosterSimulation struct {
	simulationCount     int
	NumberOfSimulations int
	NumberOfEmployees   int
	NumberOfDays        int
	PopulationSize      int
}

// Go initializes a random roster
func (r RosterSimulation) Go() ga.Bitset {
	size := r.NumberOfDays * r.NumberOfEmployees * 3
	bitset := ga.Bitset{}
	bitset.Create(size)
	for i := 0; i < size; i++ {
		bitset.Set(i, rand.Intn(2))
	}
	return bitset
}

func (r *RosterSimulation) OnBeginSimulation() {
	r.simulationCount++
	if r.NumberOfSimulations < 1 {
		panic("NumberOfSimulations must be greater than 0")
	}
}

// Simulate assigns a fitness value to the given genome
func (r *RosterSimulation) Simulate(genome *ga.IGenome) {
	roster := NewRoster(*genome, r.NumberOfEmployees, r.NumberOfDays)
	fitness := roster.GetFitness()
	(*genome).SetFitness(fitness)
}

// OnElite prints the current elite on every simulation interation
func (r *RosterSimulation) OnElite(genome *ga.IGenome) {
	roster := NewRoster(*genome, r.NumberOfEmployees, r.NumberOfDays)
	fmt.Printf("** [%d] simulation **\n", r.simulationCount)
	roster.PrintSchedule()
}

// ExitFunc defines when to stop the simulation
func (r *RosterSimulation) ExitFunc(genome *ga.IGenome) bool {
	return r.simulationCount >= r.NumberOfSimulations
}

func (r *RosterSimulation) OnEndSimulation() {}
