package main

import (
	"bufio"
	"fmt"
	"os"
)

type Form struct {
	combinedAnswers map[rune]bool
	answersInCommon string
}

func parseForm(lines []string) Form {
	form := Form{
		combinedAnswers: make(map[rune]bool),
		answersInCommon: lines[0],
	}

	var nextAnswersInCommon string
	for _, line := range lines {
		nextAnswersInCommon = ""
		for _, char := range line {
			form.combinedAnswers[char] = true

			for _, commonChar := range form.answersInCommon {
				if char == commonChar {
					nextAnswersInCommon += string(char)
					continue
				}
			}
		}
		form.answersInCommon = nextAnswersInCommon
	}

	return form
}

func solve(forms []Form) (int, int) {
	sumCombined, sumCommon := 0, 0
	for _, form := range forms {
		sumCombined += len(form.combinedAnswers)
		sumCommon += len(form.answersInCommon)
	}
	return sumCombined, sumCommon
}

func main() {
	file, _ := os.Open("/home/nimbus/advent-of-code-20/day-06/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var currLines []string
	var forms []Form
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			currLines = append(currLines, line)
		} else {
			// Form data finished
			forms = append(forms, parseForm(currLines))
			currLines = nil
		}
	}
	// Process last form
	forms = append(forms, parseForm(currLines))

	sumCombined, sumCommon := solve(forms)
	fmt.Println("Part 1 Solution: ", sumCombined)
	fmt.Println("Part 2 Solution: ", sumCommon)
}
