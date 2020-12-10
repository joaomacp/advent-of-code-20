package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var cache = make(map[int]int)

func countDistribution(adapters []int) int {
	oneJoltDiffs, threeJoltDiffs := 0, 0
	for i := 0; i < len(adapters)-1; i++ {
		switch adapters[i+1] - adapters[i] {
		case 1:
			oneJoltDiffs++
		case 3:
			threeJoltDiffs++
		}
	}
	threeJoltDiffs++

	return oneJoltDiffs * threeJoltDiffs
}

func countArrangements(adapters []int, index int) int {
	// Memoization
	if cachedValue, ok := cache[index]; ok {
		return cachedValue
	}

	if index == len(adapters)-1 {
		return 1
	}

	sum := 0
	for i := index + 1; i < len(adapters) && adapters[i] <= adapters[index]+3; i++ {
		sum += countArrangements(adapters, i)
	}
	cache[index] = sum
	return sum
}

func main() {
	file, _ := os.Open("C:/Users/Nimbus/Documents/advent-of-code-20/day-10/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var adapters []int
	adapters = append(adapters, 0)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, num)
	}

	sort.Ints(adapters)
	fmt.Println("Part 1 Solution: ", countDistribution(adapters))
	fmt.Println("Part 2 Solution: ", countArrangements(adapters, 0))
}
