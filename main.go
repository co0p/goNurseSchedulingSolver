package main

import (
	"fmt"
	ga "github.com/tomcraven/goga"
	"runtime"
	"time"
)

func main() {

	simulation := RosterSimulation{
		NumberOfSimulations: 100,
		PopulationSize:      20,
		NumberOfEmployees:   3,
		NumberOfDays:        31,
	}

	// mater defines how to combine genomes
	mater := ga.NewMater(
		[]ga.MaterFunctionProbability{
			{P: 1.0, F: ga.TwoPointCrossover},
			{P: 1.0, F: ga.Mutate},
			{P: 1.0, F: ga.UniformCrossover, UseElite: true},
		},
	)

	// selector defines how to select genomes from which the elite is being taken
	selector := ga.NewSelector(
		[]ga.SelectorFunctionProbability{
			{P: 1.0, F: ga.Roulette},
		},
	)

	algorithm := ga.NewGeneticAlgorithm()
	algorithm.Simulator = &simulation
	algorithm.EliteConsumer = &simulation
	algorithm.BitsetCreate = &simulation
	algorithm.Selector = selector
	algorithm.Mater = mater

	numThreads := 4
	runtime.GOMAXPROCS(numThreads)
	algorithm.Init(simulation.PopulationSize, numThreads)

	startTime := time.Now()
	algorithm.Simulate()
	fmt.Println(time.Since(startTime))
}
