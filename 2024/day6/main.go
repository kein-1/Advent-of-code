package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// PartOne()
	PartTwo()
}

type Game struct {
	grid          [][]rune
	startingPoint Coordinate
}

type Coordinate struct {
	x int
	y int
}

func PartOne() {

	g := loadData()
	ans := parseData(g)
	fmt.Println("The answer to part 1 is:", ans)
}

func loadData() Game {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)

	g := Game{}
	grid := make([][]rune, 0)

	rowNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, 0)
		for col, v := range line {
			row = append(row, v)
			if v == '^' {
				startingPoint := Coordinate{
					x: rowNumber,
					y: col,
				}
				g.startingPoint = startingPoint
			}
		}
		rowNumber++
		grid = append(grid, row)
	}
	g.grid = grid

	return g
}

// keep track of direction and where we are, and keep a visited map for unique locations
// at each iteartion, check surroundings; if obstacle in direction, turn 90
// otherwise no obstacle, so continue;
// when we exit: out of bounds
func parseData(g Game) int {

	direction := "UP"
	visited := make(map[Coordinate]bool, 0)
	visited[g.startingPoint] = true
	currPosition := g.startingPoint
	for {
		var newX int
		var newY int
		var dir []int
		switch direction {
		case "UP":
			dir = []int{-1, 0}
		case "DOWN":
			dir = []int{1, 0}
		case "LEFT":
			dir = []int{0, -1}
		default:
			dir = []int{0, 1}
		}
		// compute the next coordinate; then check if there is obstacle
		newX = currPosition.x + dir[0]
		newY = currPosition.y + dir[1]
		if inBounds(g.grid, newX, newY) && g.grid[newX][newY] == '#' {
			direction = changeDirection(direction)
		} else {
			// no obstacle; so update position
			currPosition = Coordinate{
				x: newX,
				y: newY,
			}
		}

		if !inBounds(g.grid, currPosition.x, currPosition.y) {
			break
		}

		visited[currPosition] = true
	}

	return len(visited)
}

func changeDirection(direction string) string {
	var dir string
	switch direction {
	case "UP":
		dir = "RIGHT"
	case "DOWN":
		dir = "LEFT"
	case "LEFT":
		dir = "UP"
	default:
		dir = "DOWN"
	}
	return dir
}

// check if we left the grid
func inBounds(grid [][]rune, x, y int) bool {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return false
	}
	return true
}
