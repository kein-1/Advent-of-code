package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// PartOne()
	PartTwo()
}

func PartOne() {
	data := loadData()
	ans := parseData(data)
	fmt.Println("The answer to par 1 is: ", ans)
}

func loadData() [][]int {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)
	data := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		row := []int{}
		for _, val := range arr {
			number, err := strconv.Atoi(val)
			if err != nil {
				panic("Error converting!")
			}
			row = append(row, number)
		}
		data = append(data, row)
	}
	return data
}

func parseData(data [][]int) int {
	ans := 0
	for _, row := range data {
		var upwards bool
		flag := true
		for index := range len(row) - 1 {
			curr := row[index]
			next := row[index+1]

			// check progression
			if index == 0 {
				if curr > next {
					upwards = false
				} else {
					upwards = true
				}
			} else {
				if !checkMonotocity(curr, next, upwards) {
					flag = false
					break
				}
			}
			if !checkDifference(curr, next) {
				flag = false
				break
			}
		}
		if flag {
			ans++
		}
	}
	return ans
}

func checkMonotocity(x int, y int, upwards bool) bool {
	if upwards {
		return x < y
	}
	return x > y
}

func checkDifference(x int, y int) bool {
	result := int(math.Abs(float64(x - y)))
	if result >= 1 && result <= 3 {
		return true
	}
	return false
}
