package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data := loadData()
	ans := getScore(data)
	fmt.Println("Total pts ", ans)
}

func loadData() []string {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)

	moves := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		moves = append(moves, line)
	}

	return moves
}

func getScore(data []string) int {
	moveMap := map[string]int{
		"A X": 0,
		"A Y": 3,
		"A Z": 6,
		"B X": 0,
		"B Y": 3,
		"B Z": 6,
		"C X": 0,
		"C Y": 3,
		"C Z": 6,
	}

	// a,x -> rock
	// b,y -> paper
	// c,z -> scissor

	sum := 0
	for _, v := range data {
		points, ok := moveMap[v]
		if !ok {
			panic("This key does not exist!")
		}
		sum += points + getShapeScore(v)
	}
	return sum

}

func getShapeScore(shape string) int {

	switch shape {
	case "A X":
		return 3
	case "A Y":
		return 1
	case "A Z":
		return 2
	case "B X":
		return 1
	case "B Y":
		return 2
	case "B Z":
		return 3
	case "C X":
		return 2
	case "C Y":
		return 3
	default:
		return 1
	}
}
