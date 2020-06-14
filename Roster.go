package main

import (
	"fmt"
	ga "github.com/tomcraven/goga"
)

// ShiftType will be encoded as bits
type ShiftType int

const (
	SHIFT_MORNING ShiftType = iota
	SHIFT_LATE
	SHIFT_NIGHT
)

type Roster struct {
	NumberOfDays      int
	NumberOfEmployees int
	ShiftTypes        []ShiftType
	genome            ga.IGenome
}

func NewRoster(genome ga.IGenome, employees int, days int) *Roster {
	return &Roster{
		NumberOfEmployees: employees,
		NumberOfDays:      days,
		genome:            genome,
	}
}

func (r Roster) GetFitness() int {
	fmt.Println("todo")
	return 0
}

func (r Roster) PrintSchedule() {
	fmt.Println("todo")
}
