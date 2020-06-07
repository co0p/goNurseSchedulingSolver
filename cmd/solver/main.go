package main

import (
	"fmt"
	ga "github.com/tomcraven/goga"
	"goNurseSchedulingSolver"
	"runtime"
	"time"
)

func main() {

	var populationSize = 100

	algorithm := ga.NewGeneticAlgorithm()
	algorithm.Simulator = goNurseSchedulingSolver.NewRosterSimulation(31, 5)
	algorithm.BitsetCreate = &goNurseSchedulingSolver.RosterInitializer{}
	algorithm.EliteConsumer = &goNurseSchedulingSolver.RosterGenomeConsumer{}

	algorithm.Mater = ga.NewMater(
		[]ga.MaterFunctionProbability{
			{P: 1.0, F: ga.TwoPointCrossover},
			{P: 1.0, F: ga.Mutate},
			{P: 1.0, F: ga.UniformCrossover, UseElite: true},
		},
	)
	algorithm.Selector = ga.NewSelector(
		[]ga.SelectorFunctionProbability{
			{P: 1.0, F: ga.Roulette},
		},
	)

	numThreads := 4
	runtime.GOMAXPROCS(numThreads)
	algorithm.Init(populationSize, numThreads)

	startTime := time.Now()
	algorithm.Simulate()
	fmt.Println(time.Since(startTime))
}
