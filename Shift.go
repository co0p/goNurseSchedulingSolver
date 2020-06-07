package goNurseSchedulingSolver

type ShiftType int

const (
	Morning ShiftType = iota
	Afternoon
	Night
)

func (s ShiftType) String() string {
	return [...]string{"Morning", "Afternoon", "Night"}[s]
}
