package main

import (
	"bufio"
	"fmt"
	"os"
)

type Slope struct {
	right int
	down  int
}

type Problem struct {
	input  []string
	slopes []Slope
}

func (p Problem) treesHit(slope Slope) int {
	height, width := len(p.input), len(p.input[0])

	treesHit := 0

	for row, column := 0, 0; row < height; row, column = row+slope.down, (column+slope.right)%width {
		if p.input[row][column] == '#' {
			treesHit++
		}
	}

	return treesHit
}

func (p Problem) solve() int {
	product := 1
	for _, slope := range p.slopes {
		product *= p.treesHit(slope)
	}
	return product
}

func main() {
	file, _ := os.Open("/home/nimbus/advent-of-code-20/day-3/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var input []string
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	partOneSlopes := []Slope{{right: 3, down: 1}}

	partTwoSlopes := []Slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	problemPartOne := Problem{input, partOneSlopes}
	problemPartTwo := Problem{input, partTwoSlopes}

	fmt.Println("Part 1 Solution: ", problemPartOne.solve())
	fmt.Println("Part 2 Solution: ", problemPartTwo.solve())
}
