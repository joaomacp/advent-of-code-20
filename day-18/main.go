package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func shuntingYard(tokens []rune, precedence map[rune]int) (output []rune) {
	var opStack []rune

	for _, token := range tokens {
		if token == '+' || token == '*' {
			for len(opStack) > 0 {
				top := len(opStack) - 1
				if (opStack[top] == '+' || opStack[top] == '*') &&
					precedence[opStack[top]] >= precedence[token] {
					output = append(output, opStack[top])
					opStack = opStack[:top]
				} else {
					break
				}
			}
			opStack = append(opStack, token)
		} else if token == '(' {
			opStack = append(opStack, token)
		} else if token == ')' {
			for opStack[len(opStack)-1] != '(' {
				top := len(opStack) - 1
				output = append(output, opStack[top])
				opStack = opStack[:top]
			}
			top := len(opStack) - 1
			opStack = opStack[:top]
		} else { // token is a number
			output = append(output, token)
		}
	}
	for len(opStack) > 0 {
		top := len(opStack) - 1
		output = append(output, opStack[top])
		opStack = opStack[:top]
	}

	return output
}

func executePostfix(tokens []rune) int {
	var operandStack []int

	for _, token := range tokens {
		if token == '+' || token == '*' {
			top := len(operandStack) - 1
			op1 := operandStack[top]
			operandStack = operandStack[:top]
			top--
			op2 := operandStack[top]
			operandStack = operandStack[:top]

			if token == '+' {
				operandStack = append(operandStack, op1+op2)
			} else if token == '*' {
				operandStack = append(operandStack, op1*op2)
			}
		} else { // token is a number
			intToken, _ := strconv.Atoi(string(token))
			operandStack = append(operandStack, intToken)
		}
	}

	return operandStack[0]
}

func solve(expressions []string, operatorPrecedence map[rune]int) (sum int) {
	for _, expr := range expressions {
		var tokens []rune
		for _, char := range expr {
			if char != ' ' {
				tokens = append(tokens, char)
			}
		}
		postfixExpr := shuntingYard(tokens, operatorPrecedence)
		sum += executePostfix(postfixExpr)
	}

	return
}

func main() {
	file, _ := os.Open("C:/Users/Nimbus/Documents/advent-of-code-20/day-18/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var expressions []string
	for scanner.Scan() {
		expressions = append(expressions, scanner.Text())
	}

	operatorPrecedence := make(map[rune]int)

	// Part 1: same precedence
	operatorPrecedence['+'] = 2
	operatorPrecedence['*'] = 2
	fmt.Println("Part 1 Solution: ", solve(expressions, operatorPrecedence))

	// Part 2: addition evaluated before multiplication
	operatorPrecedence['*'] = 1
	fmt.Println("Part 2 Solution: ", solve(expressions, operatorPrecedence))
}
