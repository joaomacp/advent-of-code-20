package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vector4 struct {
	x int
	y int
	z int
	w int
}

type ConwayCubeGrid struct {
	activeCubes  map[Vector4]bool // Sparse "matrix"
	nActiveCubes int
	minVertex    Vector4 // min vertex in the 4D-cuboid space of active cubes
	maxVertex    Vector4 // max vertex in the 4D-cuboid space of active cubes
}

func newConwayCubeGrid(lines []string) *ConwayCubeGrid {
	grid := ConwayCubeGrid{
		activeCubes: make(map[Vector4]bool),
		minVertex:   Vector4{0, 0, 0, 0},
		maxVertex:   Vector4{len(lines[0]), len(lines), 0, 0},
	}

	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				grid.activeCubes[Vector4{x, y, 0, 0}] = true
				grid.nActiveCubes++
			}
		}
	}

	return &grid
}

func (c *ConwayCubeGrid) copy() *ConwayCubeGrid {
	gridCopy := ConwayCubeGrid{activeCubes: make(map[Vector4]bool)}

	for k, v := range c.activeCubes {
		gridCopy.activeCubes[k] = v
	}

	gridCopy.nActiveCubes = c.nActiveCubes
	gridCopy.minVertex = c.minVertex
	gridCopy.maxVertex = c.maxVertex

	return &gridCopy
}

func (c *ConwayCubeGrid) nActiveNeighbors(cube Vector4) (activeNeighbors int) {
	for x := cube.x - 1; x <= cube.x+1; x++ {
		for y := cube.y - 1; y <= cube.y+1; y++ {
			for z := cube.z - 1; z <= cube.z+1; z++ {
				for w := cube.w - 1; w <= cube.w+1; w++ {
					neighbor := Vector4{x, y, z, w}
					if neighbor != cube {
						active, exists := c.activeCubes[neighbor]
						if exists && active {
							activeNeighbors++
						}
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
				for w := c.minVertex.w - 1; w <= c.maxVertex.w+1; w++ {
					cube := Vector4{x, y, z, w}
					active, exists := c.activeCubes[cube]
					activeNeighbors := c.nActiveNeighbors(cube)
					if exists && active && (activeNeighbors < 2 || activeNeighbors > 3) {
						newGrid.activeCubes[cube] = false
						newGrid.nActiveCubes--
					}
					if (!exists || !active) && activeNeighbors == 3 {
						newGrid.activeCubes[cube] = true
						newGrid.nActiveCubes++

						// 4D-Cube activated: update 4D-cuboid area of active cubes
						if cube.x < c.minVertex.x {
							newGrid.minVertex.x = cube.x
						}
						if cube.y < c.minVertex.y {
							newGrid.minVertex.y = cube.y
						}
						if cube.z < c.minVertex.z {
							newGrid.minVertex.z = cube.z
						}
						if cube.w < c.minVertex.w {
							newGrid.minVertex.w = cube.w
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
						if cube.w > c.maxVertex.w {
							newGrid.maxVertex.w = cube.w
						}
					}
				}
			}
		}
	}

	return newGrid
}

func solvePartTwo(grid *ConwayCubeGrid) int {
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
	fmt.Println("Part 2 Solution: ", solvePartTwo(grid))
}
