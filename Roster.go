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

type Assignment struct {
	employee int
	shifts   []ShiftType
}

type Roster struct {
	NumberOfDays      int
	NumberOfEmployees int
	ShiftTypes        []ShiftType
	schedule          []Assignment
}

func NewRoster(genome ga.IGenome, employees int, days int) *Roster {
	r := &Roster{
		NumberOfEmployees: employees,
		NumberOfDays:      days,
		schedule:          []Assignment{},
	}
	// for each employee go over the bitset and extract the current shift assigment for the whole period of time
	bits := genome.GetBits()
	for i := 0; i < employees; i++ {
		assignment := Assignment{
			employee: i,
			shifts:   []ShiftType{},
		}
		for idx := i; idx < bits.GetSize(); idx += employees {
			chromosome := bits.Get(idx)
			shift := ShiftType(chromosome)
			assignment.shifts = append(assignment.shifts, shift)
		}
		r.schedule = append(r.schedule, assignment)
	}

	return r
}

func (r Roster) GetFitness() int {

	// TODO calculate fitness or r's schedule

	return 0
}

func (r Roster) PrintSchedule() {

	for _, s := range r.schedule {
		fmt.Printf("%d\t: ", s.employee)
		for _, a := range s.shifts {
			fmt.Printf("%d | ", a)
		}
		fmt.Printf("\n")
	}
}
