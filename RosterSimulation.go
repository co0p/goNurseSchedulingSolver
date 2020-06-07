package goNurseSchedulingSolver

import (
	ga "github.com/tomcraven/goga"
	"log"
)

type Employee struct {
	maxHours int
}

func NewEmployee() Employee {
	return Employee{
		maxHours: 40,
	}
}

type Assignment struct {
	Employee Employee
	Shift    ShiftType
}

type Schedule struct {
	weekDay     int
	assignments []Assignment
}

type RosterSimulation struct {
	roster Roster
}

func NewRosterSimulation(days int, employeeCount int) *RosterSimulation {
	return &RosterSimulation{
		roster: NewRoster(employeeCount, days),
	}
}

func (s *RosterSimulation) OnBeginSimulation() {
	log.Printf("Begin Simulation")
}

func (s *RosterSimulation) Simulate(genome *ga.IGenome) {
	panic("implement me")
}

func (s *RosterSimulation) OnEndSimulation() {
	log.Printf("Begin Simulation")
}

func (s *RosterSimulation) ExitFunc(genome *ga.IGenome) bool {
	panic("implement me")
}
