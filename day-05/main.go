package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func getSeatId(seat string) int {
	rowSpecification := seat[:7]
	colSpecification := seat[7:]

	low, high := 0, 127
	for _, direction := range rowSpecification {
		if direction == 'F' {
			high = (high + low) / 2
		} else {
			low = (high+low)/2 + 1
		}
	}
	row := low

	low, high = 0, 7
	for _, direction := range colSpecification {
		if direction == 'L' {
			high = (high + low) / 2
		} else {
			low = (high+low)/2 + 1
		}
	}
	col := low

	return row*8 + col
}

func processSeats(seats []string) (int, int, int) {
	min, max, sum := math.MaxInt32, 0, 0

	for _, seat := range seats {
		seatId := getSeatId(seat)

		sum += seatId
		if seatId < min {
			min = seatId
		}
		if seatId > max {
			max = seatId
		}
	}

	return min, max, sum
}

func getMissingId(minSeatId int, maxSeatId int, sumIdsFound int) int {
	sumIdsAll := 0
	for i := minSeatId; i <= maxSeatId; i++ {
		sumIdsAll += i
	}
	return sumIdsAll - sumIdsFound
}

func main() {
	file, _ := os.Open("/home/nimbus/advent-of-code-20/day-05/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var seats []string
	for scanner.Scan() {
		seat := scanner.Text()
		seats = append(seats, seat)
	}

	minSeatId, maxSeatId, sumIdsFound := processSeats(seats)

	fmt.Println("Part 1 Solution: ", maxSeatId)
	fmt.Println("Part 2 Solution: ", getMissingId(minSeatId, maxSeatId, sumIdsFound))
}
