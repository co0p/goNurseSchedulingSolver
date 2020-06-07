package goNurseSchedulingSolver

import (
	ga "github.com/tomcraven/goga"
)

type Roster struct {
	employees []Employee
	schedules []Schedule
}

func (r *Roster) GetFitness() int {
	panic("implement me")
}

func (r *Roster) SetFitness(i int) {
	panic("implement me")
}

func (r *Roster) GetBits() *ga.Bitset {
	panic("implement me")
}

func (r *Roster) String() string {
	return ""
}

func NewRosterFromGenome(g *ga.IGenome) *Roster {
	return &Roster{}
}

func NewRoster(employeeCount int, days int) Roster {
	employees := []Employee{}
	for i := 0; i < employeeCount; i++ {
		employees = append(employees, NewEmployee())
	}

	schedules := []Schedule{}
	for i := 0; i < days; i++ {
		schedules = append(schedules, Schedule{})
	}
	return Roster{
		employees: employees,
		schedules: schedules,
	}
}
