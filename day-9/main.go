package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func firstInvalidNum(data []int64) int64 {
	for currIndex := 25; currIndex < len(data); currIndex++ {
		currNum := data[currIndex]
		valid := false

		for i := currIndex - 25; i < currIndex-1; i++ {
			for j := i + 1; j < currIndex; j++ {
				if data[i] != data[j] && data[i]+data[j] == currNum {
					valid = true
				}
			}
		}

		if !valid {
			return currNum
		}
	}
	return 0
}

func findWeakness(data []int64, invalidNum int64) int64 {
	lowIndex, highIndex := 0, 1
	currSum := data[lowIndex] + data[highIndex]

	for {
		diff := currSum - invalidNum
		if diff == 0 {
			break
		}

		if diff > 0 {
			currSum -= data[lowIndex]
			lowIndex++
			if lowIndex == highIndex {
				highIndex++
				currSum += data[highIndex]
			}
		}

		if diff < 0 {
			highIndex++
			currSum += data[highIndex]
		}
	}

	var min, max int64 = math.MaxInt64, 0
	for i := lowIndex; i <= highIndex; i++ {
		if data[i] < min {
			min = data[i]
		}
		if data[i] > max {
			max = data[i]
		}
	}
	return min + max
}

func main() {
	file, _ := os.Open("C:/Users/Nimbus/Documents/advent-of-code-20/day-9/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var xmasData []int64
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		xmasData = append(xmasData, int64(num))
	}

	firstInvalidNum := firstInvalidNum(xmasData)
	fmt.Println("Part 1 Solution: ", firstInvalidNum)
	fmt.Println("Part 2 Solution: ", findWeakness(xmasData, firstInvalidNum))
}
