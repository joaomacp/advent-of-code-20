package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solvePartOne(ourArrivalTime int, busIds []int) int {
	minWaitTime := math.MaxInt32
	bestBusId := 0

	for _, busId := range busIds {
		if busId != -1 {
			lastBusPassed := ourArrivalTime % busId

			if lastBusPassed == 0 {
				return 0 // 0*ID = 0
			}

			waitTime := busId - lastBusPassed
			if waitTime < minWaitTime {
				minWaitTime = waitTime
				bestBusId = busId
			}
		}
	}

	return minWaitTime * bestBusId
}

// PART 2: Chinese remainder theorem:
// https://www.geeksforgeeks.org/chinese-remainder-theorem-set-2-implementation/

// Modular multiplicative inverse: (a*x) % m = 1; x = moduloInverse(a,m)
// Using extended Euclidean algorithm (from link above)
func moduloInverse(a int64, m int64) int64 {
	if m == 1 {
		return 0
	}

	var m0, x0, x1, t, q int64
	m0 = m
	x1 = 1

	// Apply extended Euclid Algorithm
	for a > 1 {
		q = a / m
		t = m
		m = a % m
		a = t
		t = x0
		x0 = x1 - q*x0
		x1 = t
	}

	if x1 < 0 {
		x1 += m0
	}

	return x1
}

func chineseRemainder(nums []int, rems []int) int64 {
	// Product of all nums
	prod := int64(1)
	for _, num := range nums {
		prod *= int64(num)
	}

	var result int64
	for i, num := range nums {
		pp := prod / int64(num)
		result += int64(rems[i]) * moduloInverse(pp, int64(num)) * pp
	}

	return result % prod
}

func solvePartTwo(busIds []int) int64 {
	var nums, rems []int

	for i, busId := range busIds {
		if busId != -1 {
			nums = append(nums, busId)
			if i == 0 {
				rems = append(rems, 0)
			} else {
				rems = append(rems, busId-i)
			}
		}
	}

	return chineseRemainder(nums, rems)
}

func main() {
	file, _ := os.Open("C:/Users/Nimbus/Documents/advent-of-code-20/day-13/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	ourArrivalTime, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	buses := scanner.Text()

	var busIds []int
	for _, bus := range strings.Split(buses, ",") {
		if bus == "x" {
			busIds = append(busIds, -1)
		} else {
			busId, _ := strconv.Atoi(bus)
			busIds = append(busIds, busId)
		}
	}

	fmt.Println("Part 1 Solution: ", solvePartOne(ourArrivalTime, busIds))
	fmt.Println("Part 2 Solution: ", solvePartTwo(busIds))
}
