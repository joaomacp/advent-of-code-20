package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	value int
	child string
}

type TwoWayGraph struct {
	parents    map[string][]string
	children   map[string][]Edge
	nBagsCache map[string]int
}

func parseGraph(input []string) TwoWayGraph {
	parents := make(map[string][]string)
	children := make(map[string][]Edge)

	for _, line := range input {
		parts := strings.Split(line, " contain ")
		parentParts := strings.Split(parts[0], " ")
		parentColor := parentParts[0] + parentParts[1]

		childrenParts := strings.Split(parts[1], ", ")
		if childrenParts[0] != "no other bags." {
			for _, childString := range childrenParts {
				childParts := strings.Split(childString, " ")
				childValue, _ := strconv.Atoi(childParts[0])
				childColor := childParts[1] + childParts[2]

				parents[childColor] = append(parents[childColor], parentColor)
				children[parentColor] = append(children[parentColor], Edge{
					value: childValue,
					child: childColor,
				})
			}
		}
	}

	return TwoWayGraph{
		parents:    parents,
		children:   children,
		nBagsCache: make(map[string]int),
	}
}

func (g TwoWayGraph) countUniqueParents(node string) int {
	visited := make(map[string]bool)

	var queue []string
	queue = append(queue, node)

	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		for _, parentNode := range g.parents[currNode] {
			if _, ok := visited[parentNode]; !ok {
				visited[parentNode] = true
				queue = append(queue, parentNode)
			}
		}
	}

	return len(visited)
}

func (g TwoWayGraph) sumOfPathProducts(node string, accum int) int {
	// Memoization: reduces recursive calls 51->22
	if cachedValue, ok := g.nBagsCache[node]; ok {
		return cachedValue * accum
	}

	sum := 0
	for _, edge := range g.children[node] {
		mul := accum * edge.value
		sum += mul + g.sumOfPathProducts(edge.child, mul)
	}
	g.nBagsCache[node] = sum / accum
	return sum
}

func main() {
	file, _ := os.Open("/home/nimbus/advent-of-code-20/day-7/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	graph := parseGraph(input)
	fmt.Println("Part 1 Solution: ", graph.countUniqueParents("shinygold"))
	fmt.Println("Part 2 Solution: ", graph.sumOfPathProducts("shinygold", 1))
}
