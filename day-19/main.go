package main

import (
	"bufio"
	"fmt"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
	"os"
	"strings"
)

// For recursive regex, golang-pkg-pcre is used, which requires libpcre++-dev

func solveRules(rules map[string]string) string {
	hasMoreRules := true
	for hasMoreRules {
		r := rules["0"]

		hasMoreRules = false
		for _, token := range strings.Split(r, " ") {
			if token != "" && token != "|" && token[0] != '(' && token[0] != ')' && token[0] != '"' && token[0] != '*' && token[0] != '+' && token[0] != '?' {
				hasMoreRules = true
				// Replace all, including overlapping
				for strings.Count(rules["0"], " "+token+" ") > 0 {
					rules["0"] = strings.ReplaceAll(rules["0"], " "+token+" ", " ("+rules[token]+") ")
				}
			}
		}
	}

	rules["0"] = strings.ReplaceAll(rules["0"], " ", "")
	rules["0"] = strings.ReplaceAll(rules["0"], "\"", "")
	return "^" + rules["0"] + "$"
}

func solve(rules map[string]string, testStrings []string) (matchingStrings int) {
	regexString := solveRules(rules)
	regex := pcre.MustCompile(regexString, 0)

	for _, testStr := range testStrings {
		if regex.MatcherString(testStr, 0).Matches() {
			matchingStrings++
		}
	}
	return
}

func main() {
	file, _ := os.Open("/home/nimbus/advent-of-code-20/day-19/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rules := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lineParts := strings.Split(line, ":")
		rules[lineParts[0]] = lineParts[1] + " "
	}

	var testStrings []string
	for scanner.Scan() {
		testStrings = append(testStrings, scanner.Text())
	}

	fmt.Println("Part 1 Solution: ", solve(rules, testStrings))

	// Reset rule 0
	rules["0"] = " 8 11 "

	// Set new rules
	rules["8"] = " ( 42 ) + "
	rules["11"] = " ( ?<recursion> 42 (?&recursion)? 31 ) "

	fmt.Println("Part 2 Solution: ", solve(rules, testStrings))
}
