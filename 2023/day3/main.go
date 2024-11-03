package main

// #TODO - COMPLETE IN FUTURE

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// load it all into a matrix, then scan;
	// when we find the start of a number, we can find the whole number, then just mark whether along the way any part
	// of the number is adjacent to a symbol. Once we get to a non number, we know our current numbe is done, so just reset

	// 467..114..
	// ...*......
	// ..35..633.
	// ......#...
	// 617*......
	// .....+.58.
	// ..592.....
	// ......755.
	// ...$.*....
	// .664.598..

	// 1,296,374 - incorrect
	// 538,293

	matrix := loadData("input.txt")

	ans := 0
	for i, v := range matrix {
		flag := false
		number := 0
		for j := range v {
			// fmt.Println("Number and flag: ", number, flag)
			val := matrix[i][j]

			if unicode.IsDigit(val) {
				if checkNumber(matrix, i, j) {
					flag = true
				}
				number = number*10 + int(val-'0')
			} else {
				if flag {
					fmt.Println("this is valid number we are adding: ", number)
					ans += number
					fmt.Println("curr sum: ", ans)
				}
				flag = false
				number = 0
			}
		}
	}
	fmt.Println("total:", ans)

}

func checkNumber(matrix [][]rune, x int, y int) bool {
	// up down left right top left top right bottom right bottom left
	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
	}

	for _, dir := range directions {
		newX := x + dir[0]
		newY := y + dir[1]
		if newX < 0 || newX >= len(matrix) || newY < 0 || newY >= len(matrix[0]) {
			continue
		}
		if matrix[newX][newY] == 46 {
			continue
		}
		if isSpecialRune(matrix[newX][newY]) {
			return true
		}
	}
	return false
}

func isSpecialRune(r rune) bool {
	switch r {
	case '+', '%', '*', '$', '#', '/', '=', '@', '-', '&':
		return true
	default:
		return false
	}
}

func loadData(fileName string) [][]rune {
	fi, err := os.Open(fileName)
	if err != nil {
		panic("Error opening")
	}

	defer fi.Close()

	matrix := make([][]rune, 140)
	for i, _ := range matrix {
		matrix[i] = make([]rune, 140)
	}

	scanner := bufio.NewScanner(fi)
	row := 0
	for scanner.Scan() {
		text := scanner.Text()

		for i, r := range text {
			matrix[row][i] = r
		}
		row++
	}
	return matrix
}
