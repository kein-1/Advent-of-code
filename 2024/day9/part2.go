package main

import (
	"fmt"
	"strconv"
)

func PartTwo() {
	data := loadData()
	ans := parseData2(data)
	fmt.Println("The answer to part 2 is:", ans)
}

func parseData2(data []Row) int {
	ans := 0

	for _, v := range data {
		if dfs2(v.value, 0, 0, v.numbers, "+") {
			ans += v.value
		}
	}

	return ans
}

// just another option of || - this means concat this with what we currently have
func dfs2(value int, total int, currIndex int, numbers []int, operation string) bool {

	if currIndex >= len(numbers) {
		return value == total
	}

	currentElement := numbers[currIndex]

	currSum := total
	if operation == "+" {
		currSum += currentElement
	} else if operation == "*" {
		currSum *= currentElement
	} else {
		currSum = concat(total, currentElement)
	}

	operations := []string{
		"+", "*", "||",
	}

	for _, operation_ := range operations {
		if dfs2(value, currSum, currIndex+1, numbers, operation_) {
			return true
		}
	}
	return false

}

func concat(numb1 int, numb2 int) int {

	str1 := strconv.Itoa(numb1)
	str2 := strconv.Itoa(numb2)

	result := str1 + str2

	r, err := strconv.Atoi(result)
	if err != nil {
		panic("error converting concat")
	}
	return r
}
