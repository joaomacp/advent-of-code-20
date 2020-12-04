package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solvePartOne(input []int) int {
	for i := 0; i < len(input)-1; i++ {
		curr := input[i]
		for j := i + 1; j < len(input); j++ {
			if curr+input[j] == 2020 {
				return curr * input[j]
			}
		}
	}
	return -1
}

func solvePartTwo(input []int) int {
	for i := 0; i < len(input)-2; i++ {
		curr_i := input[i]
		for j := i + 1; j < len(input)-1; j++ {
			curr_j := input[j]
			for k := i + 2; k < len(input); k++ {
				if curr_i+curr_j+input[k] == 2020 {
					return curr_i * curr_j * input[k]
				}
			}
		}
	}
	return -1
}

func main() {
	file, _ := os.Open("/home/nimbus/advent-of-code-20/day-1/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var input []int
	for scanner.Scan() {
		entry, _ := strconv.Atoi(scanner.Text())
		input = append(input, entry)
	}

	fmt.Println("Part 1 Solution: ", solvePartOne(input))

	fmt.Println("Part 2 Solution: ", solvePartTwo(input))
}
