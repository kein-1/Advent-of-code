package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type List struct {
	left  []int
	right []int
}

func main() {
	// PartOne()
	PartTwo()
}

func PartOne() {
	list := loadData()
	ans := parseData(list)
	fmt.Println("The answer to par 1 is: ", ans)
}

func loadData() List {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)
	var list List
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "   ")

		leftNumber, err := strconv.Atoi(arr[0])
		if err != nil {
			panic("Error converting left")
		}
		rightNumber, err := strconv.Atoi(arr[1])
		if err != nil {
			panic("Error converting right")
		}

		list.left = append(list.left, leftNumber)
		list.right = append(list.right, rightNumber)
	}
	slices.Sort(list.left)
	slices.Sort(list.right)

	return list
}

func parseData(list List) int {
	ans := 0
	for index := range list.left {
		leftNumb := list.left[index]
		rightNumb := list.right[index]
		ans += int(math.Abs(float64(leftNumb - rightNumb)))
	}
	return ans
}
