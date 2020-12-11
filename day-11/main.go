package main

import (
	"bufio"
	"fmt"
	"os"
)

type GridState struct {
	grid          [][]rune
	occupiedSeats int
}

func (g GridState) print() {
	fmt.Println("Occupied: ", g.occupiedSeats)
	for _, row := range g.grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func (g GridState) copy() *GridState {
	grid := make([][]rune, len(g.grid))
	for i := range g.grid {
		grid[i] = make([]rune, len(g.grid[i]))
		copy(grid[i], g.grid[i])
	}

	return &GridState{grid, g.occupiedSeats}
}

type SeatCalculator interface {
	nOccupiedSeats(state *GridState, rowIndex int, colIndex int) int
}

type AdjacentSeatCalculator struct{}

func (p AdjacentSeatCalculator) nOccupiedSeats(state *GridState, rowIndex int, colIndex int) (adjacentOccupied int) {
	for r := rowIndex - 1; r <= rowIndex+1; r++ {
		if r >= 0 && r < len(state.grid) {
			for c := colIndex - 1; c <= colIndex+1; c++ {
				if c >= 0 && c < len(state.grid[0]) {
					if !(rowIndex == r && colIndex == c) && state.grid[r][c] == '#' {
						adjacentOccupied++
					}
				}
			}
		}
	}

	return
}

type VisibleSeatCalculator struct{}

func (p VisibleSeatCalculator) nOccupiedSeats(state *GridState, rowIndex int, colIndex int) (visibleOccupied int) {
	for dirY := -1; dirY <= 1; dirY++ {
		for dirX := -1; dirX <= 1; dirX++ {
			if dirX == 0 && dirY == 0 {
				continue
			}

			for mult := 1; ; mult++ {
				r := rowIndex + mult*dirY
				c := colIndex + mult*dirX
				if r < 0 || r >= len(state.grid) ||
					c < 0 || c >= len(state.grid[0]) ||
					state.grid[r][c] == 'L' {
					break
				}
				if state.grid[r][c] == '#' {
					visibleOccupied++
					break
				}
			}
		}
	}

	return
}

type Rules struct {
	occupiedThreshold int
	seatCalculator    SeatCalculator
}

func (r Rules) applyRules(currState *GridState) (newState *GridState, changesMade int) {
	newState = currState.copy()

	for rowIndex, row := range currState.grid {
		for colIndex, char := range row {
			if char == 'L' && r.seatCalculator.nOccupiedSeats(currState, rowIndex, colIndex) == 0 {
				newState.grid[rowIndex][colIndex] = '#'
				newState.occupiedSeats++
				changesMade++
			} else if char == '#' && r.seatCalculator.nOccupiedSeats(currState, rowIndex, colIndex) >= r.occupiedThreshold {
				newState.grid[rowIndex][colIndex] = 'L'
				newState.occupiedSeats--
				changesMade++
			}
		}
	}

	return
}

func solve(state *GridState, rules Rules) int {
	var changesMade int

	for {
		//state.print()
		state, changesMade = rules.applyRules(state)

		if changesMade == 0 {
			return state.occupiedSeats
		}
	}
}

func main() {
	file, _ := os.Open("C:/Users/Nimbus/Documents/advent-of-code-20/day-11/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	gridState := GridState{grid: grid, occupiedSeats: 0}

	partOneRules := Rules{
		occupiedThreshold: 4,
		seatCalculator:    AdjacentSeatCalculator{},
	}
	partTwoRules := Rules{
		occupiedThreshold: 5,
		seatCalculator:    VisibleSeatCalculator{},
	}

	fmt.Println("Part 1 Solution: ", solve(&gridState, partOneRules))
	fmt.Println("Part 2 Solution: ", solve(&gridState, partTwoRules))
}
