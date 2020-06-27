package main

import (
	ga "github.com/tomcraven/goga"
	"testing"
)

func TestNewRoster_should_create_valid_roster_from_genome(t *testing.T) {
	employeeCount := 9
	days := 23

	bitset := ga.Bitset{}
	bitset.Create(employeeCount * days)
	bitset.SetAll(0)
	genome := ga.NewGenome(bitset)

	r := NewRoster(genome, employeeCount, days)

	// we should have one 10 schedules
	if len(r.schedule) != employeeCount {
		t.Errorf("expected %d schedules, got %d instead", employeeCount, len(r.schedule))
	}

	for _, v := range r.schedule {
		if len(v.shifts) != days {
			t.Errorf("expected %d for each shift, got %d instead", days, len(v.shifts))
		}
	}
}
