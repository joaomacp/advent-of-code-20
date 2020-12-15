package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(startingNums []string, nThNumSpoken int) int {
	mem := make(map[int]int)

	for i, numString := range startingNums {
		num, _ := strconv.Atoi(numString)
		mem[num] = i
	}

	currNum := 0
	for i := len(startingNums); i < nThNumSpoken-1; i++ {
		lastNum := currNum
		if lastIndex, ok := mem[currNum]; ok {
			currNum = i - lastIndex
		} else {
			currNum = 0
		}
		mem[lastNum] = i
	}

	return currNum
}

func main() {
	file, _ := os.Open("C:/Users/Nimbus/Documents/advent-of-code-20/day-15/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	input := scanner.Text()

	fmt.Println("Part 1 Solution: ", solve(strings.Split(input, ","), 2020))
	fmt.Println("Part 2 Solution: ", solve(strings.Split(input, ","), 30000000))

}
