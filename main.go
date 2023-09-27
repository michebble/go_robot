package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/michebble/go_robot/robot"
)

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
	var rob robot.Robot
	lines, err := readLines(argsWithoutProg)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, line := range lines {
		switch line {
		case "LEFT":
			rob.TurnLeft()
		case "RIGHT":
			rob.TurnRight()
		case "MOVE":
			rob.Move()
		case "REPORT":
			report := rob.Report()
			fmt.Println(report)
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
				rob = *robot.NewRobot(placements[2], x, y)
			} else {
				fmt.Println("unknown command", line)
			}
		}
	}

}
