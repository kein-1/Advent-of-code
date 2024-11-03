package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := loadData()
	ans := getPairs(data)
	fmt.Println("Total sum ", ans)
}

func loadData() []string {
	fi, err := os.Open("input2.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)

	data := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	return data
}

func getPairs(data []string) int {
	ans := 0
	for _, v := range data {
		r := strings.Split(v, ",")
		left, right := strings.Split(r[0], "-"), strings.Split(r[1], "-")
		s, err := strconv.Atoi(left[0])
		if err != nil {
			panic("failed conversion")
		}
		e, err := strconv.Atoi(left[1])
		if err != nil {
			panic("failed conversion")
		}

		a, err := strconv.Atoi(right[0])
		if err != nil {
			panic("failed conversion")
		}

		b, err := strconv.Atoi(right[1])
		if err != nil {
			panic("failed conversion")
		}
		if checkOverlap(s, e, a, b) {
			ans++
		}
	}
	return ans
}

func checkInterval(s int, e int, a int, b int) bool {

	if s <= a && e >= b {
		return true
	}
	if a <= s && b >= e {
		return true
	}
	return false
}

func checkOverlap(s int, e int, a int, b int) bool {

	if e < a {
		return false
	}
	if b < s {
		return false
	}
	return true
}
