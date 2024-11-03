package main

import (
	"bufio"
	"fmt"
	"os"
)

// Sliding window problem
func main() {
	data := loadData()
	ans := getMarkers(data)
	fmt.Println("The first marker appears after: ", ans)
}

func loadData() string {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)
	var data string
	for scanner.Scan() {
		data = scanner.Text()
	}
	return data
}

func getMarkers(data string) int {
	// this is sliding window

	markers := map[rune]int{}
	start := 0
	for end, v := range data {
		val, ok := markers[v]
		if !ok {
			markers[v] = 1
		} else {
			markers[v] = val + 1
		}

		for markers[v] > 1 {
			fmt.Println("This letter and count is greater than 1:", string(v), markers[v])
			sChar := rune(data[start])
			sVal, ok := markers[sChar]
			if !ok {
				panic("this key does not exist!")
			}
			if sVal == 1 {
				delete(markers, rune(data[start]))
			} else {
				markers[sChar] = sVal - 1
			}
			start++
		}

		if len(markers) == 14 {
			debugMap(markers)
			return end + 1
		}

	}

	return 0
}

func debugMap(m map[rune]int) {
	for k, v := range m {
		fmt.Println(string(k), v)
	}
}
