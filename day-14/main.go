package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type InitProgram struct {
	instructions []string
	mask         string
	mem          map[int]int
}

func (p InitProgram) sumMemValues() (sum int) {
	for _, value := range p.mem {
		sum += value
	}
	return
}

func (p InitProgram) run(partTwo bool) {
	for _, instr := range p.instructions {
		if instr[0:7] == "mask = " {
			p.mask = instr[7:]
		} else {
			closeBracket := strings.IndexAny(instr, "]")
			memAddress, _ := strconv.Atoi(instr[4:closeBracket])
			valueToWrite, _ := strconv.Atoi(strings.Split(instr, " ")[2])

			if !partTwo {
				p.mem[memAddress] = applyMaskPartOne(valueToWrite, p.mask)
			} else {
				for _, addr := range applyMaskPartTwo(memAddress, p.mask) {
					p.mem[addr] = valueToWrite
				}
			}
		}
	}
}

func applyMaskPartOne(n int, mask string) int {
	for i, char := range mask {
		pos := 35 - i
		if char == '0' {
			n &^= 1 << pos
		} else if char == '1' {
			n |= 1 << pos
		}
	}
	return n
}

func applyMaskPartTwo(originalAddress int, mask string) (addressesToSet []int) {
	var floatingBits []int
	for i, char := range mask {
		pos := 35 - i
		if char == '1' {
			originalAddress |= 1 << pos
		} else if char == 'X' {
			originalAddress &^= 1 << pos
			floatingBits = append(floatingBits, pos)
		}
	}

	for seq := 0; seq < powInt(2, len(floatingBits)); seq++ {
		address := originalAddress
		for bitIndex, bit := range floatingBits {
			bitPos := len(floatingBits) - 1 - bitIndex
			if seq&powInt(2, bitPos) != 0 {
				address |= 1 << bit
			}
		}
		addressesToSet = append(addressesToSet, address)
	}

	return
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func main() {
	file, _ := os.Open("C:/Users/Nimbus/Documents/advent-of-code-20/day-14/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var instructions []string
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, line)
	}

	prog := InitProgram{instructions: instructions, mem: make(map[int]int)}

	prog.run(false)
	fmt.Println("Part 1 Solution: ", prog.sumMemValues())

	prog.mem = make(map[int]int) // Reset mem
	prog.run(true)
	fmt.Println("Part 2 Solution", prog.sumMemValues())
}
