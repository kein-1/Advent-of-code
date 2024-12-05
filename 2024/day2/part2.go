package main

import (
	"fmt"
)

type Data struct {
	left       []int
	rightCount map[int]int
}

func PartTwo() {
	data := loadData()
	ans := parseData2(data)
	fmt.Println("The answer to part 2 is: ", ans)
}

func parseData2(data [][]int) int {
	ans := 0
	for _, row := range data {
		reportStatus, index1, index2 := checkReport(row)

		if reportStatus {
			ans++
		} else {
			// Bad report; now decide which element to remove: if we have both index, either element is contributing to the problem
			// so we try both
			sliceOne := updateSlice(row, index1)
			sliceTwo := updateSlice(row, index2)
			sliceThree := updateSlice(row, 0)

			check1, _, _ := checkReport(sliceOne)
			check2, _, _ := checkReport(sliceTwo)
			check3, _, _ := checkReport(sliceThree)
			if check1 || check2 || check3 {
				ans++
			}

			// check removing first element
			fmt.Println("")
		}
	}
	return ans
}

func checkReport(row []int) (bool, int, int) {
	var upwards bool
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
				return false, index, index + 1
			}
		}
		if !checkDifference(curr, next) {
			return false, index, index + 1
		}
	}
	return true, -1, -1
}

func updateSlice(row []int, index int) []int {
	row2 := make([]int, len(row))
	copy(row2, row)
	slice := append(row2[:index], row2[index+1:]...)
	return slice
}
