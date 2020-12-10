package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type FieldValidator interface {
	validate(value string) bool
}

type IntRangeValidator struct {
	minValue int
	maxValue int
}

func (v IntRangeValidator) validate(value string) bool {
	valueInt, err := strconv.Atoi(value)
	return err == nil && v.minValue <= valueInt && valueInt <= v.maxValue
}

type HeightValidator struct {
	inchValidator IntRangeValidator
	cmValidator   IntRangeValidator
}

func (v HeightValidator) validate(value string) bool {
	leadingValue := value[:len(value)-2]
	lastTwoChars := value[len(value)-2:]

	if lastTwoChars == "in" {
		return v.inchValidator.validate(leadingValue)
	} else if lastTwoChars == "cm" {
		return v.cmValidator.validate(leadingValue)
	} else {
		return false
	}
}

type EyeColorValidator struct {
	eyeColors []string
}

func (v EyeColorValidator) validate(value string) bool {
	for _, color := range v.eyeColors {
		if value == color {
			return true
		}
	}
	return false
}

type RegexValidator struct {
	regexPattern string
}

func (v RegexValidator) validate(value string) bool {
	matched, err := regexp.MatchString(v.regexPattern, value)
	return err == nil && matched
}

type Passport struct {
	fields map[string]string
}

func parsePassport(textLines []string) Passport {
	p := Passport{fields: make(map[string]string)}
	for _, line := range textLines {
		for _, field := range strings.Split(line, " ") {
			fieldParts := strings.Split(field, ":")
			p.fields[fieldParts[0]] = fieldParts[1]
		}
	}
	return p
}

func (p Passport) containsRequiredFields(requiredFields []string) bool {
	for _, requiredField := range requiredFields {
		if _, ok := p.fields[requiredField]; !ok {
			return false
		}
	}
	return true
}

func (p Passport) isValid(fieldValidators map[string]FieldValidator) bool {
	for field, validator := range fieldValidators {
		if !validator.validate(p.fields[field]) {
			return false
		}
	}
	return true
}

func main() {
	file, _ := os.Open("/home/nimbus/advent-of-code-20/day-04/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	fieldValidators := map[string]FieldValidator{
		"byr": IntRangeValidator{minValue: 1920, maxValue: 2002},
		"iyr": IntRangeValidator{minValue: 2010, maxValue: 2020},
		"eyr": IntRangeValidator{minValue: 2020, maxValue: 2030},
		"hgt": HeightValidator{
			inchValidator: IntRangeValidator{minValue: 59, maxValue: 76},
			cmValidator:   IntRangeValidator{minValue: 150, maxValue: 193},
		},
		"hcl": RegexValidator{regexPattern: "^#[a-f0-9]{6}$"},
		"ecl": EyeColorValidator{eyeColors: []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}},
		"pid": RegexValidator{regexPattern: "^[0-9]{9}$"},
	}

	passportsWithReqFields := 0
	validPassports := 0
	var currLines []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			currLines = append(currLines, line)
		} else {
			// Passport data finished
			passport := parsePassport(currLines)
			if passport.containsRequiredFields(requiredFields) {
				passportsWithReqFields++
				if passport.isValid(fieldValidators) {
					validPassports++
				}
			}
			currLines = nil
		}
	}

	// Process last passport
	passport := parsePassport(currLines)
	if passport.containsRequiredFields(requiredFields) {
		passportsWithReqFields++
		if passport.isValid(fieldValidators) {
			validPassports++
		}
	}

	fmt.Println("Part 1 Solution: ", passportsWithReqFields)
	fmt.Println("Part 2 Solution: ", validPassports)
}
