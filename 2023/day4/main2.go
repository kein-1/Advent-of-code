package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Numbers struct {
	winningNumbs map[int]bool
	total        int
}

func main() {
	// 122359 wrong
	data := loadData()
	ans := calculateValue(data)
	fmt.Println(ans)
}

func loadData() []string {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()

	data := []string{}
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		data = append(data, line)
	}
	return data
}

func calculateValue(data []string) int {

	cardMap := make(map[int]int)

	for i, v := range data {
		cardMap[i] = cardMap[i] + 1
		winningNumbersMap, currNumbers := parseNumbers(v)
		matchCounter := 0
		for _, v := range currNumbers {
			if len(v) == 0 {
				continue
			}
			_, ok := winningNumbersMap[v]
			if ok {
				matchCounter += 1
			}
		}
		for j := 1; j <= matchCounter; j++ {

			cardMap[i+j] = cardMap[i+j] + cardMap[i]
		}

		fmt.Println("after card i and map: ", i, cardMap)
	}

	ans := 0
	for _, v := range cardMap {
		ans += v
	}
	fmt.Println(cardMap)
	return ans
}

func parseNumbers(str string) (map[string]bool, []string) {
	winningNumbsMap := map[string]bool{}
	str_ := strings.Split(str, "|")
	left := strings.Split(str_[0], ":")
	_, winningNumbStr := left[0], strings.Trim(left[1], " ")

	numbers := strings.Split(winningNumbStr, " ")
	for _, v := range numbers {
		winningNumbsMap[v] = true
	}

	right := strings.Split(strings.Trim(str_[1], " "), " ")

	return winningNumbsMap, right
}
