package robot_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/michebble/go_robot/robot"
)

func TestRobot(t *testing.T) {
	t.Parallel()
	_ = robot.Robot{
		Facing: "NORTH",
		X:      2,
		Y:      1,
	}
}

func TestNewRobot(t *testing.T) {
	t.Parallel()
	want := robot.Robot{
		Facing: "NORTH",
		X:      2,
		Y:      1,
	}
	got := *robot.NewRobot("NORTH", 2, 1)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestReport(t *testing.T) {
	t.Parallel()
	type testCase struct {
		r    robot.Robot
		want string
	}
	testCases := []testCase{
		{r: robot.Robot{Facing: "NORTH", X: 0, Y: 1}, want: "0,1,NORTH"},
		{r: robot.Robot{Facing: "SOUTH", X: 2, Y: 3}, want: "2,3,SOUTH"},
	}
	for _, tc := range testCases {
		got := tc.r.Report()

		if tc.want != got {
			t.Errorf("Expect %s, got %s", tc.want, got)
		}
	}

}

func TestTurnLeft(t *testing.T) {
	t.Parallel()
	type testCase struct {
		r    robot.Robot
		want robot.Robot
	}
	testCases := []testCase{
		{
			r:    robot.Robot{Facing: "NORTH", X: 0, Y: 1},
			want: robot.Robot{Facing: "WEST", X: 0, Y: 1},
		},
		{
			r:    robot.Robot{Facing: "WEST", X: 0, Y: 1},
			want: robot.Robot{Facing: "SOUTH", X: 0, Y: 1},
		},
		{
			r:    robot.Robot{Facing: "SOUTH", X: 0, Y: 1},
			want: robot.Robot{Facing: "EAST", X: 0, Y: 1},
		},
		{
			r:    robot.Robot{Facing: "EAST", X: 0, Y: 1},
			want: robot.Robot{Facing: "NORTH", X: 0, Y: 1},
		},
	}
	for _, tc := range testCases {
		tc.r.TurnLeft()

		if !cmp.Equal(tc.want, tc.r) {
			t.Error(cmp.Diff(tc.want, tc.r))
		}
	}
}

func TestTurnRight(t *testing.T) {
	t.Parallel()
	type testCase struct {
		r    robot.Robot
		want robot.Robot
	}
	testCases := []testCase{
		{
			r:    robot.Robot{Facing: "NORTH", X: 0, Y: 1},
			want: robot.Robot{Facing: "EAST", X: 0, Y: 1},
		},
		{
			r:    robot.Robot{Facing: "EAST", X: 0, Y: 1},
			want: robot.Robot{Facing: "SOUTH", X: 0, Y: 1},
		},
		{
			r:    robot.Robot{Facing: "SOUTH", X: 0, Y: 1},
			want: robot.Robot{Facing: "WEST", X: 0, Y: 1},
		},
		{
			r:    robot.Robot{Facing: "WEST", X: 0, Y: 1},
			want: robot.Robot{Facing: "NORTH", X: 0, Y: 1},
		},
	}
	for _, tc := range testCases {
		tc.r.TurnRight()

		if !cmp.Equal(tc.want, tc.r) {
			t.Error(cmp.Diff(tc.want, tc.r))
		}
	}
}

func TestMove(t *testing.T) {
	t.Parallel()
	type testCase struct {
		r    robot.Robot
		want robot.Robot
	}
	testCases := []testCase{
		{
			r:    robot.Robot{Facing: "NORTH", X: 2, Y: 2},
			want: robot.Robot{Facing: "NORTH", X: 3, Y: 2},
		},
		{
			r:    robot.Robot{Facing: "NORTH", X: 4, Y: 2},
			want: robot.Robot{Facing: "NORTH", X: 4, Y: 2},
		},
		{
			r:    robot.Robot{Facing: "EAST", X: 2, Y: 2},
			want: robot.Robot{Facing: "EAST", X: 2, Y: 3},
		},
		{
			r:    robot.Robot{Facing: "EAST", X: 2, Y: 4},
			want: robot.Robot{Facing: "EAST", X: 2, Y: 4},
		},
		{
			r:    robot.Robot{Facing: "SOUTH", X: 2, Y: 2},
			want: robot.Robot{Facing: "SOUTH", X: 1, Y: 2},
		},
		{
			r:    robot.Robot{Facing: "SOUTH", X: 0, Y: 2},
			want: robot.Robot{Facing: "SOUTH", X: 0, Y: 2},
		},
		{
			r:    robot.Robot{Facing: "WEST", X: 2, Y: 2},
			want: robot.Robot{Facing: "WEST", X: 2, Y: 1},
		},
		{
			r:    robot.Robot{Facing: "WEST", X: 2, Y: 0},
			want: robot.Robot{Facing: "WEST", X: 2, Y: 0},
		},
	}
	for _, tc := range testCases {
		tc.r.Move()

		if !cmp.Equal(tc.want, tc.r) {
			t.Error(cmp.Diff(tc.want, tc.r))
		}
	}
}
