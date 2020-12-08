package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Direct threaded interpreter:
// https://www.usenix.org/legacy/publications/library/proceedings/jvm01/gagnon/gagnon_html/node4.html

type Instr struct {
	icode string
	run   func(int, *Program)
	arg   int
}

type Program struct {
	instructions []Instr
	pc           int
	acc          int
	visitedPCs   map[int]bool
}

func newProgram(instrs []Instr) *Program {
	p := new(Program)
	p.instructions = instrs
	p.visitedPCs = make(map[int]bool)
	return p
}

func parseProgram(code []string) *Program {
	var instructions []Instr

	for _, line := range code {
		parts := strings.Split(line, " ")
		arg, _ := strconv.Atoi(parts[1])

		var icode string
		var run func(int, *Program)
		switch parts[0] {
		case "nop":
			icode = "nop"
			run = nop
		case "acc":
			icode = "acc"
			run = add
		case "jmp":
			icode = "jmp"
			run = jmp
		}

		instructions = append(instructions, Instr{icode: icode, run: run, arg: arg})
	}

	return newProgram(instructions)
}

func nop(_ int, p *Program) {
	p.visitedPCs[p.pc] = true
	p.pc++
	if p.pc == len(p.instructions) {
		return
	}
	if _, exists := p.visitedPCs[p.pc]; exists {
		return
	}
	nextInstr := p.instructions[p.pc]
	nextInstr.run(nextInstr.arg, p)
}

func add(arg int, p *Program) {
	p.acc += arg

	p.visitedPCs[p.pc] = true
	p.pc++
	if p.pc == len(p.instructions) {
		return
	}
	if _, exists := p.visitedPCs[p.pc]; exists {
		return
	}
	nextInstr := p.instructions[p.pc]
	nextInstr.run(nextInstr.arg, p)
}

func jmp(arg int, p *Program) {
	p.visitedPCs[p.pc] = true
	p.pc += arg
	if p.pc == len(p.instructions) {
		return
	}
	if _, exists := p.visitedPCs[p.pc]; exists {
		return
	}
	nextInstr := p.instructions[p.pc]
	nextInstr.run(nextInstr.arg, p)
}

func run(prog *Program) {
	// Start program by running first instruction
	firstInstr := prog.instructions[0]
	firstInstr.run(firstInstr.arg, prog)
}

func fixProgram(prog *Program) int {
	for i, currInstr := range prog.instructions {
		newInstrs := make([]Instr, len(prog.instructions))
		copy(newInstrs, prog.instructions)
		if currInstr.icode == "nop" {
			newInstrs[i].run = jmp
		} else if currInstr.icode == "jmp" {
			newInstrs[i].run = nop
		} else {
			continue
		}

		newProg := newProgram(newInstrs)
		run(newProg)
		if newProg.pc == len(newProg.instructions) {
			return newProg.acc
		}
	}

	return -1
}

func main() {
	file, _ := os.Open("/home/nimbus/advent-of-code-20/day-8/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var code []string
	for scanner.Scan() {
		code = append(code, scanner.Text())
	}

	prog := parseProgram(code)
	run(prog)
	fmt.Println("Part 1 Solution: ", prog.acc)
	fmt.Println("Part 2 Solution: ", fixProgram(prog))
}
