package robot

import "fmt"

type Robot struct {
	Facing string
	X      int
	Y      int
}

func NewRobot(facing string, x, y int) *Robot {
	r := Robot{Facing: facing, X: x, Y: y}
	return &r
}

func (r *Robot) Report() string {
	position := fmt.Sprintf("%d,%d,%s", r.X, r.Y, r.Facing)
	return position
}

func (r *Robot) TurnLeft() {
	switch r.Facing {
	case "NORTH":
		r.Facing = "WEST"
	case "WEST":
		r.Facing = "SOUTH"
	case "SOUTH":
		r.Facing = "EAST"
	case "EAST":
		r.Facing = "NORTH"
	}
}

func (r *Robot) TurnRight() {
	switch r.Facing {
	case "NORTH":
		r.Facing = "EAST"
	case "EAST":
		r.Facing = "SOUTH"
	case "SOUTH":
		r.Facing = "WEST"
	case "WEST":
		r.Facing = "NORTH"
	}
}

func (r *Robot) Move() {
	switch r.Facing {
	case "NORTH":
		r.X = increment(r.X, 4)
	case "EAST":
		r.Y = increment(r.Y, 4)
	case "SOUTH":
		r.X = decrement(r.X, 0)
	case "WEST":
		r.Y = decrement(r.Y, 0)
	}
}

// increment returns the input value incremented by 1 if it does not exceed the bound.
// Otherwise is returns the input value.
func increment(input, bound int) int {
	i := input + 1
	if i <= bound {
		return i
	}
	return input
}

// decrement returns the input value decremented by 1 if it does not exceed the bound.
// Otherwise is returns the input value.
func decrement(input, bound int) int {
	i := input - 1
	if i >= bound {
		return i
	}
	return input
}
