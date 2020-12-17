package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vector3 struct {
	x int
	y int
	z int
}

type ConwayCubeGrid struct {
	activeCubes  map[Vector3]bool // Sparse "matrix"
	nActiveCubes int
	minVertex    Vector3 // min vertex in the cuboid space of active cubes
	maxVertex    Vector3 // max vertex in the cuboid space of active cubes
}

func newConwayCubeGrid(lines []string) *ConwayCubeGrid {
	grid := ConwayCubeGrid{
		activeCubes: make(map[Vector3]bool),
		minVertex:   Vector3{0, 0, 0},
		maxVertex:   Vector3{len(lines[0]), len(lines), 0},
	}

	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				grid.activeCubes[Vector3{x, y, 0}] = true
				grid.nActiveCubes++
			}
		}
	}

	return &grid
}

func (c *ConwayCubeGrid) copy() *ConwayCubeGrid {
	gridCopy := ConwayCubeGrid{activeCubes: make(map[Vector3]bool)}

	for k, v := range c.activeCubes {
		gridCopy.activeCubes[k] = v
	}

	gridCopy.nActiveCubes = c.nActiveCubes
	gridCopy.minVertex = c.minVertex
	gridCopy.maxVertex = c.maxVertex

	return &gridCopy
}

func (c *ConwayCubeGrid) nActiveNeighbors(cube Vector3) (activeNeighbors int) {
	for x := cube.x - 1; x <= cube.x+1; x++ {
		for y := cube.y - 1; y <= cube.y+1; y++ {
			for z := cube.z - 1; z <= cube.z+1; z++ {
				neighbor := Vector3{x, y, z}
				if neighbor != cube {
					active, exists := c.activeCubes[neighbor]
					if exists && active {
						activeNeighbors++
					}
				}
			}
		}
	}
	return
}

func (c *ConwayCubeGrid) doIteration() *ConwayCubeGrid {
	newGrid := c.copy()

	for x := c.minVertex.x - 1; x <= c.maxVertex.x+1; x++ {
		for y := c.minVertex.y - 1; y <= c.maxVertex.y+1; y++ {
			for z := c.minVertex.z - 1; z <= c.maxVertex.z+1; z++ {
				cube := Vector3{x, y, z}
				active, exists := c.activeCubes[cube]
				activeNeighbors := c.nActiveNeighbors(cube)
				if exists && active && (activeNeighbors < 2 || activeNeighbors > 3) {
					newGrid.activeCubes[cube] = false
					newGrid.nActiveCubes--
				}
				if (!exists || !active) && activeNeighbors == 3 {
					newGrid.activeCubes[cube] = true
					newGrid.nActiveCubes++

					// Cube activated: update cuboid area of active cubes
					if cube.x < c.minVertex.x {
						newGrid.minVertex.x = cube.x
					}
					if cube.y < c.minVertex.y {
						newGrid.minVertex.y = cube.y
					}
					if cube.z < c.minVertex.z {
						newGrid.minVertex.z = cube.z
					}
					if cube.x > c.maxVertex.x {
						newGrid.maxVertex.x = cube.x
					}
					if cube.y > c.maxVertex.y {
						newGrid.maxVertex.y = cube.y
					}
					if cube.z > c.maxVertex.z {
						newGrid.maxVertex.z = cube.z
					}
				}
			}
		}
	}

	return newGrid
}

func solvePartOne(grid *ConwayCubeGrid) int {
	for i := 0; i < 6; i++ {
		grid = grid.doIteration()
	}
	return grid.nActiveCubes
}

func main() {
	file, _ := os.Open("C:/Users/Nimbus/Documents/advent-of-code-20/day-17/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	grid := newConwayCubeGrid(lines)
	fmt.Println("Part 1 Solution: ", solvePartOne(grid))
}
