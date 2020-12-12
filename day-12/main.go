package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Ship struct {
	dirs     []rune
	dirIndex int
	east     int
	north    int
	waypoint []int
}

func newShip() *Ship {
	return &Ship{
		dirs:     []rune{'E', 'S', 'W', 'N'},
		waypoint: []int{10, 1},
	}
}

func (s *Ship) processInstructionDirect(instr rune, value int) {
	switch instr {
	case 'N':
		s.north += value
	case 'S':
		s.north -= value
	case 'E':
		s.east += value
	case 'W':
		s.east -= value
	case 'R':
		for i := 0; i < value/90; i++ {
			s.dirIndex = (s.dirIndex + 1) % 4
		}
	case 'L':
		for i := 0; i < value/90; i++ {
			s.dirIndex = s.dirIndex - 1
			if s.dirIndex == -1 {
				s.dirIndex = 3
			}
		}
	case 'F':
		s.processInstructionDirect(s.dirs[s.dirIndex], value)
	}
}

func (s *Ship) processInstructionWaypoint(instr rune, value int) {
	switch instr {
	case 'N':
		s.waypoint[1] += value
	case 'S':
		s.waypoint[1] -= value
	case 'E':
		s.waypoint[0] += value
	case 'W':
		s.waypoint[0] -= value
	case 'R':
		for i := 0; i < value/90; i++ {
			east := s.waypoint[0]
			s.waypoint[0] = s.waypoint[1]
			s.waypoint[1] = -east
		}
	case 'L':
		for i := 0; i < value/90; i++ {
			north := s.waypoint[1]
			s.waypoint[1] = s.waypoint[0]
			s.waypoint[0] = -north
		}
	case 'F':
		s.east += s.waypoint[0] * value
		s.north += s.waypoint[1] * value
	}
}

func (s *Ship) navigate(navLines []string, waypointNavigation bool) {
	for _, line := range navLines {
		instruction := line[0]
		value, _ := strconv.Atoi(line[1:])

		if waypointNavigation {
			s.processInstructionWaypoint(rune(instruction), value)
		} else {
			s.processInstructionDirect(rune(instruction), value)
		}
	}
}

func manhattanDistance(x int, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}

	return x + y
}

func main() {
	file, _ := os.Open("C:/Users/Nimbus/Documents/advent-of-code-20/day-12/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var navLines []string
	for scanner.Scan() {
		line := scanner.Text()
		navLines = append(navLines, line)
	}

	ship := newShip()
	ship.navigate(navLines, false)
	fmt.Println("Part 1 Solution: ", manhattanDistance(ship.east, ship.north))

	ship = newShip()
	ship.navigate(navLines, true)
	fmt.Println("Part 2 Solution: ", manhattanDistance(ship.east, ship.north))
}
