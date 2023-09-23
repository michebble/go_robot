package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type robot struct {
	Facing string
	// might be nice to convert this into an X and Y instead of an array
	Coords [2]int
}

func newRobot(facing string, x, y int) *robot {
	coords := [2]int{x, y}
	r := robot{Facing: facing, Coords: coords}
	return &r
}

func (r *robot) report() {
	position := fmt.Sprintf("%d,%d,%s", r.Coords[0], r.Coords[1], r.Facing)
	fmt.Println(position)
}

func (r *robot) turnLeft() {
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

func (r *robot) turnRight() {
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

func (r *robot) move() {
	switch r.Facing {
	case "NORTH":
		r.Coords[1] += 1
	case "EAST":
		r.Coords[0] += 1
	case "SOUTH":
		r.Coords[1] -= 1
	case "WEST":
		r.Coords[0] -= 1
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	argsWithoutProg := os.Args[1]
	var rob robot
	lines, err := readLines(argsWithoutProg)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, line := range lines {
		switch line {
		case "LEFT":
			rob.turnLeft()
		case "RIGHT":
			rob.turnRight()
		case "MOVE":
			rob.move()
		case "REPORT":
			rob.report()
		default:
			// This should be moved to before iterating the lines slice
			// pop the first element and check if it is a placement,
			// if not quit.
			if strings.HasPrefix(line, "PLACE") {
				// needs error handling
				placement, _ := strings.CutPrefix(line, "PLACE ")
				placements := strings.Split(placement, ",")
				x, _ := strconv.Atoi(placements[0])
				y, _ := strconv.Atoi(placements[1])
				rob = *newRobot(placements[2], x, y)
			} else {
				fmt.Println("unknown command", line)
			}
		}
	}

}
