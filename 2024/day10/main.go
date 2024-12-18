package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	run()
}

func run() {

	grid := loadData()
	// ans := parseData(grid)

	ans := parseData2(grid)

	fmt.Println("The answer to part 2 is:", ans)
}

func loadData() [][]int {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)

	grid := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0)
		for _, v := range line {
			numb := int(v - '0')
			row = append(row, numb)
		}
		grid = append(grid, row)
	}

	return grid
}

// run a dfs at each 0. for a valid, gradual increase, we need current value to be 1 more than prev
// keep a visited at only the final location; basically seeing how many different 9's we can reach by dfs-ing at 0
// so if we reach that 9, just mark it
func parseData(grid [][]int) int {
	ans := 0
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == 0 {

				visited := make([][]int, 0)
				for i := 0; i < len(grid); i++ {
					row := make([]int, len(grid[0]))
					visited = append(visited, row)
				}

				fmt.Printf("The trail head at %d and %d is %d:\n", i, j, grid[i][j])
				count := dfs(grid, visited, -1, i, j)
				ans += count
				fmt.Printf("The trail head resulted in: %d\n\n", count)
			}
		}
	}
	return ans
}

// dfs
func dfs(grid, visited [][]int, prev, x, y int) int {
	// not in bounds so invalid
	if !inBounds(grid, x, y) {
		return 0
	}
	curr := grid[x][y]
	// invalid path since not gradual increase so go back
	if curr-1 != prev {
		return 0
	}

	// valid path since we reached the end, so go back
	if curr == 9 && prev == 8 {
		if visited[x][y] == 1 {
			return 0
		}
		visited[x][y] = 1
		return 1
	}

	dir := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	// now explore 4 ways by dfs each direction
	ans := 0
	for _, v := range dir {
		x_, y_ := v[0], v[1]
		newX := x + x_
		newY := y + y_
		ans += dfs(grid, visited, curr, newX, newY)
	}

	return ans
}

func parseData2(grid [][]int) int {
	ans := 0
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == 0 {

				fmt.Printf("The trail head at %d and %d is %d:\n", i, j, grid[i][j])
				count := dfs2(grid, -1, i, j)
				ans += count
				fmt.Printf("The trail head resulted in: %d\n\n", count)
			}
		}
	}
	return ans
}

// dfs for part 2 - now different from part 1; no need to make distinctions. we are just counting number of
// ways to reach 9
func dfs2(grid [][]int, prev, x, y int) int {
	// not in bounds so invalid
	if !inBounds(grid, x, y) {
		return 0
	}
	curr := grid[x][y]
	// invalid path since not gradual increase so go back
	if curr-1 != prev {
		return 0
	}

	// valid path since we reached the end, so go back
	if curr == 9 && prev == 8 {
		return 1
	}

	dir := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	// now explore 4 ways by dfs each direction
	ans := 0
	for _, v := range dir {
		x_, y_ := v[0], v[1]
		newX := x + x_
		newY := y + y_
		ans += dfs2(grid, curr, newX, newY)
	}

	return ans
}

func inBounds(grid [][]int, x, y int) bool {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return false
	}
	return true
}
