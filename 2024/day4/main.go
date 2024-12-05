package main

import (
	"bufio"
	"fmt"
	"os"
)

// 28498913 too low
func main() {
	// PartOne()
	PartTwo()
}

func PartOne() {

	data := loadData()
	ans := parseData(data)
	fmt.Println("The answer to par 1 is: ", ans)
	// for _, v := range data {
	// 	fmt.Println(v)
	// }
}

func loadData() [][]string {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}
	data := [][]string{}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for _, v := range line {
			row = append(row, string(v))
		}
		data = append(data, row)
	}
	return data
}

// dfs each time we see an X; keep an index of where we are in [XMAS]
// if currenti s not equal to index char of [XMAS], return
// otherwise, keep going.
// Important note: we should be going in a single direction; so if we came from top, we should continue dfs-downards
// if we came from top left, we dfs towards bottom left etc

func parseData(data [][]string) int {

	ans := 0
	word := "XMAS"
	for x := range len(data) {
		for y := range len(data[0]) {
			if data[x][y] == "X" {
				result := dfs(data, 0, word, x, y)
				ans += result
			}
		}
	}
	return ans
}

func dfs(data [][]string, currIndex int, word string, x int, y int) int {

	direction := []string{
		"T", "TR", "R", "BR", "B", "BL", "L", "TL",
	}

	ans := 0
	for _, dir := range direction {
		result := dfsHelper(data, currIndex, word, dir, x, y)
		ans += result
	}
	return ans

}

func dfsHelper(data [][]string, currIndex int, word string, direction string, x int, y int) int {

	// went past the length of last word so we good
	if currIndex >= len(word) {
		return 1
	}

	// if we out of bounds, in grid, return
	if x < 0 || x >= len(data) || y < 0 || y >= len(data[0]) {
		return 0
	}

	if data[x][y] != string(word[currIndex]) {
		return 0
	}

	// fmt.Println("Current direciton and letter:", direction, string(word[currIndex]))

	// dfs direction
	switch direction {

	// came from top left
	// case "TL":
	// 	return dfs(data, currIndex+1, word, "TL", x+1, y+1)

	// came from top
	case "T":
		return dfsHelper(data, currIndex+1, word, direction, x-1, y)

	// came from top right
	case "TR":
		return dfsHelper(data, currIndex+1, word, direction, x-1, y+1)

	// came from right
	case "R":
		return dfsHelper(data, currIndex+1, word, direction, x, y+1)

	// came from bottom right
	case "BR":
		return dfsHelper(data, currIndex+1, word, direction, x+1, y+1)

	// came from bottom
	case "B":
		return dfsHelper(data, currIndex+1, word, direction, x+1, y)

	// came from bottom left
	case "BL":
		return dfsHelper(data, currIndex+1, word, direction, x+1, y-1)

	// came from left
	case "L":
		return dfsHelper(data, currIndex+1, word, direction, x, y-1)

	// came from top left
	default:
		return dfsHelper(data, currIndex+1, word, direction, x-1, y-1)
	}

}
