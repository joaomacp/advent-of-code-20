package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Policy struct {
	low    int
	high   int
	letter rune
}

func parseInputLine(line string) (Policy, string) {
	var policy Policy

	parts := strings.Split(line, ": ")
	policyString, password := parts[0], parts[1]

	parts = strings.Split(policyString, "-")
	policy.low, _ = strconv.Atoi(parts[0])

	parts = strings.Split(parts[1], " ")
	policy.high, _ = strconv.Atoi(parts[0])
	policy.letter = rune(parts[1][0])

	return policy, password
}

func isValidPart1(password string, policy Policy) bool {
	occurrences := 0

	for _, char := range password {
		if char == policy.letter {
			occurrences++
		}
	}

	return (policy.low <= occurrences) && (occurrences <= policy.high)
}

func isValidPart2(password string, policy Policy) bool {
	letterAtLow := rune(password[policy.low-1]) == policy.letter
	letterAtHigh := rune(password[policy.high-1]) == policy.letter

	return letterAtLow != letterAtHigh
}

func main() {
	file, _ := os.Open("/home/nimbus/advent-of-code-20/day-02/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	validPasswordsPart1 := 0
	validPasswordsPart2 := 0

	for scanner.Scan() {
		policy, password := parseInputLine(scanner.Text())
		if isValidPart1(password, policy) {
			validPasswordsPart1++
		}
		if isValidPart2(password, policy) {
			validPasswordsPart2++
		}
	}

	fmt.Println("Part 1 Solution: ", validPasswordsPart1)
	fmt.Println("Part 2 Solution: ", validPasswordsPart2)
}
