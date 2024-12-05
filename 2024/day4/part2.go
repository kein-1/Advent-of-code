package main

import (
	"fmt"
)

func PartTwo() {
	data := loadData()
	ans := parseData2(data)
	fmt.Println("The answer to par 2 is: ", ans)
}

// maybe we start DFS at each A;
//
// for each A, we just check : top left, top right, botttom left, bottom right

// Possible:

// S.S
// .A.
// M.M

// S.M
// .A.
// S.M

// M.M
// .A.
// S.S

// M.S
// .A.
// M.S

func parseData2(data [][]string) int {
	ans := 0

	for x := range len(data) {
		for y := range len(data[0]) {
			if data[x][y] == "A" {
				result := helper(data, x, y)

				ans += result
			}
		}
	}
	return ans
}

func helper(data [][]string, x int, y int) int {

	// check:
	// S.S
	// .A.
	// M.M

	if inBounds(data, x-1, y-1) && data[x-1][y-1] == "S" && inBounds(data, x-1, y+1) && data[x-1][y+1] == "S" && inBounds(data, x+1, y-1) && data[x+1][y-1] == "M" && inBounds(data, x+1, y+1) && data[x+1][y+1] == "M" {
		return 1
	}
	// check
	// S.M
	// .A.
	// S.M

	if inBounds(data, x-1, y-1) && data[x-1][y-1] == "S" && inBounds(data, x-1, y+1) && data[x-1][y+1] == "M" && inBounds(data, x+1, y-1) && data[x+1][y-1] == "S" && inBounds(data, x+1, y+1) && data[x+1][y+1] == "M" {
		return 1
	}

	// check
	// M.M
	// .A.
	// S.S

	if inBounds(data, x-1, y-1) && data[x-1][y-1] == "M" && inBounds(data, x-1, y+1) && data[x-1][y+1] == "M" && inBounds(data, x+1, y-1) && data[x+1][y-1] == "S" && inBounds(data, x+1, y+1) && data[x+1][y+1] == "S" {
		return 1
	}

	// check
	// M.S
	// .A.
	// M.S
	if inBounds(data, x-1, y-1) && data[x-1][y-1] == "M" && inBounds(data, x-1, y+1) && data[x-1][y+1] == "S" && inBounds(data, x+1, y-1) && data[x+1][y-1] == "M" && inBounds(data, x+1, y+1) && data[x+1][y+1] == "S" {
		return 1
	}
	return 0

}

func inBounds(data [][]string, x, y int) bool {
	// if we out of bounds, in grid, return
	if x < 0 || x >= len(data) || y < 0 || y >= len(data[0]) {
		return false
	}
	return true
}
