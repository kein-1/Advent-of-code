package main

import "fmt"

func PartTwo() {
	g := loadData()
	ans := parseData2(g)
	fmt.Println("The answer to part 2 is:", ans)
}

// brute force : put obstacle in each possible position, then check if in cycle
// for cycle : maybe if visited count is like 10 for each elem, we are stuck
// in a loop
func parseData2(g Game) int {
	ans := 0
	grid := g.grid
	visited := getVisited(g)

	for k := range visited {

		grid[k.x][k.y] = '#'
		if runLoop(grid, g.startingPoint) {
			ans++
		}
		grid[k.x][k.y] = '.'
	}

	// get all visited

	return ans
}

func getVisited(g Game) map[Coordinate]bool {

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
	fmt.Println("Length of visited:", len(visited))
	return visited
}

func runLoop(grid [][]rune, start Coordinate) bool {

	direction := "UP"
	visited := make(map[Coordinate]int, 0)
	visited[start] = 1
	currPosition := start

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
		if inBounds(grid, newX, newY) && grid[newX][newY] == '#' {
			direction = changeDirection(direction)
		} else {
			// no obstacle; so update position
			currPosition = Coordinate{
				x: newX,
				y: newY,
			}
		}

		if !inBounds(grid, currPosition.x, currPosition.y) {
			return false
		}

		visited[currPosition]++
		if visited[currPosition] >= 5 {
			return true
		}
	}
}
