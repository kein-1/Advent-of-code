package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	data := loadData()
	ans := getMax2(data)
	fmt.Println("Total calories: ", ans)
}

func loadData() []string {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)

	calories := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		calories = append(calories, line)
	}

	return calories
}

func getMax(data []string) int {
	ans := math.MinInt
	sum := 0
	for _, v := range data {
		if len(v) == 0 {
			if sum > ans {
				ans = sum
			}
			sum = 0
		} else {
			integer, err := strconv.Atoi(v)
			if err != nil {
				panic("Error converting in calories")
			}
			sum += integer
		}
	}
	return ans
}

func getMax2(data []string) int {
	ans := []int{}
	sum := 0
	for _, v := range data {
		if len(v) == 0 {
			ans = append(ans, sum)
			sum = 0
		} else {
			integer, err := strconv.Atoi(v)
			if err != nil {
				panic("Error converting in calories")
			}
			sum += integer
		}
	}
	sort.Ints(ans)

	total := 0
	lastThree := ans[len(ans)-3:]
	for _, v := range lastThree {
		total += v
	}

	return total
}
