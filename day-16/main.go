package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	range1 []int
	range2 []int
}

func (r Range) containsValue(v int) bool {
	return (v >= r.range1[0] && v <= r.range1[1]) ||
		(v >= r.range2[0] && v <= r.range2[1])
}

type Ticket struct {
	values []int
}

func parseTicket(valuesString string) (ticket Ticket) {
	for _, valueString := range strings.Split(valuesString, ",") {
		value, _ := strconv.Atoi(valueString)
		ticket.values = append(ticket.values, value)
	}
	return
}

type Notes struct {
	fieldRanges map[string]Range
}

func (n *Notes) addFieldRange(fieldRangeString string) {
	parts := strings.Split(fieldRangeString, ": ")
	field := parts[0]
	ranges := strings.Split(parts[1], " or ")
	range1s := strings.Split(ranges[0], "-")
	range2s := strings.Split(ranges[1], "-")

	range1 := make([]int, 2)
	range1[0], _ = strconv.Atoi(range1s[0])
	range1[1], _ = strconv.Atoi(range1s[1])

	range2 := make([]int, 2)
	range2[0], _ = strconv.Atoi(range2s[0])
	range2[1], _ = strconv.Atoi(range2s[1])

	n.fieldRanges[field] = Range{range1: range1, range2: range2}
}

type PossibilityMatrixEntry struct {
	nPossibilities int
	possibilities  []bool
}

func (e *PossibilityMatrixEntry) firstPossibility() int {
	for i, value := range e.possibilities {
		if value {
			return i
		}
	}
	return -1
}

func solvePartOne(nearbyTickets []Ticket, notes *Notes) (sumInvalidValues int, validTickets []Ticket) {
	for _, ticket := range nearbyTickets {
		validTicket := true
		for _, value := range ticket.values {
			validValue := false
			for _, r := range notes.fieldRanges {
				if r.containsValue(value) {
					validValue = true
					break
				}
			}
			if !validValue {
				validTicket = false
				sumInvalidValues += value
			}
		}
		if validTicket {
			validTickets = append(validTickets, ticket)
		}
	}

	return
}

func selectPossibilities(tickets []Ticket, notes *Notes) [][]string {
	var allFieldNames []string
	for fieldName := range notes.fieldRanges {
		allFieldNames = append(allFieldNames, fieldName)
	}

	possibilities := make([][]string, len(notes.fieldRanges))

	for pos := 0; pos < len(possibilities); pos++ {
		possibilities[pos] = allFieldNames
		for _, ticket := range tickets {
			var newPossibilities []string
			for _, possibility := range possibilities[pos] {
				if notes.fieldRanges[possibility].containsValue(ticket.values[pos]) {
					newPossibilities = append(newPossibilities, possibility)
				}
			}
			possibilities[pos] = newPossibilities
		}
	}

	return possibilities
}

func filterPossibilities(possibilities [][]string) map[string]int {
	possibilityMatrix := make(map[string]*PossibilityMatrixEntry)

	for i, possibilityList := range possibilities {
		for _, possibility := range possibilityList {
			if _, ok := possibilityMatrix[possibility]; !ok {
				possibilityMatrix[possibility] = &PossibilityMatrixEntry{
					nPossibilities: 0,
					possibilities:  make([]bool, len(possibilities)),
				}
			}
			possibilityMatrix[possibility].nPossibilities++
			possibilityMatrix[possibility].possibilities[i] = true
		}
	}

	uniquePossibilities := make(map[string]int)

	for len(uniquePossibilities) < len(possibilities) {
		for field, entry := range possibilityMatrix {
			if entry.nPossibilities == 1 {
				if _, ok := uniquePossibilities[field]; !ok {
					possibility := entry.firstPossibility()
					uniquePossibilities[field] = possibility
					for k, v := range possibilityMatrix {
						if k != field && v.possibilities[possibility] {
							v.possibilities[possibility] = false
							v.nPossibilities--
						}
					}
				}
			}
		}
	}

	return uniquePossibilities
}

func solvePartTwo(ourTicket Ticket, tickets []Ticket, notes *Notes) int {
	possibilities := selectPossibilities(tickets, notes)
	uniquePossibilities := filterPossibilities(possibilities)

	result := 1
	for field, position := range uniquePossibilities {
		if len(field) > 9 && field[:9] == "departure" {
			result *= ourTicket.values[position]
		}
	}
	return result
}

func main() {
	file, _ := os.Open("C:/Users/Nimbus/Documents/advent-of-code-20/day-16/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	notes := &Notes{fieldRanges: make(map[string]Range)}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			notes.addFieldRange(line)
		} else {
			break
		}
	}

	scanner.Scan()
	scanner.Scan()
	ourTicket := parseTicket(scanner.Text())

	scanner.Scan()
	scanner.Scan()

	var nearbyTickets []Ticket
	for scanner.Scan() {
		nearbyTickets = append(nearbyTickets, parseTicket(scanner.Text()))
	}

	sumInvalidValues, validTickets := solvePartOne(nearbyTickets, notes)
	fmt.Println("Part 1 Solution: ", sumInvalidValues)
	fmt.Println("Part 2 Solution", solvePartTwo(ourTicket, validTickets, notes))
}
