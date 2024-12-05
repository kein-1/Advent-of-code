package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	left       []int
	rightCount map[int]int
}

func PartTwo() {
	data := loadData2()
	ans := parseData2(data)
	fmt.Println("The answer to part 2 is: ", ans)
}

func newData() *Data {
	var data Data
	data.rightCount = make(map[int]int)
	return &data
}

func loadData2() *Data {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)
	data := newData()
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

		data.left = append(data.left, leftNumber)
		val, ok := data.rightCount[rightNumber]
		if !ok {
			data.rightCount[rightNumber] = 1
		} else {
			data.rightCount[rightNumber] = val + 1
		}
	}

	return data
}

func parseData2(data *Data) int {
	ans := 0
	for _, val := range data.left {
		count, ok := data.rightCount[val]
		if ok {
			ans += val * count
		}
	}
	return ans
}
