package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data := loadData()
	ans := getSum(data)
	fmt.Println("Total priority sum ", ans)
}

func loadData() [][]string {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)

	compartment := [][]string{}

	count := 0
	temp := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		temp = append(temp, line)
		count++
		if count%3 == 0 {
			compartment = append(compartment, temp)
			temp = []string{}
		}
	}

	return compartment
}

func getSum(data [][]string) int {
	ans := 0

	for _, group := range data {
		a, b, c := group[0], group[1], group[2]
		commonMap := map[rune]bool{}

		aMap := map[rune]int{}
		for _, c := range a {
			val, ok := aMap[c]
			if !ok {
				aMap[c] = 1
			} else {
				aMap[c] = val + 1
			}
		}

		for _, c := range b {
			_, ok := aMap[c]
			if ok {
				commonMap[c] = true
			}
		}

		for _, c := range c {
			_, ok := commonMap[c]
			if ok {
				ans += getPriority(c)
				break
			}
		}

	}
	return ans
}

func getPriority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	}
	return int(r - 'A' + 27)
}

// func getSum(data []string) int {
// 	ans := 0
// 	for _, v := range data {
// 		length := len(v)
// 		left := v[:length/2]
// 		right := v[length/2:]

// 		leftMap := map[rune]int{}
// 		for _, c := range left {
// 			val, ok := leftMap[c]
// 			if !ok {
// 				leftMap[c] = 1
// 			} else {
// 				leftMap[c] = val + 1
// 			}
// 		}

// 		for _, c := range right {
// 			_, ok := leftMap[c]
// 			if ok {
// 				ans += getPriority(c)
// 				break
// 			}
// 		}
// 	}
// 	return ans
// }

// func getPriority(r rune) int {
// 	if r >= 'a' && r <= 'z' {
// 		return int(r - 'a' + 1)
// 	}
// 	return int(r - 'A' + 27)
// }
