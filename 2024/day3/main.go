package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// 28498913 too low
func main() {
	// PartOne()
	PartTwo()
}

func PartOne() {
	regex := `mul\(\d+,\d+\)`

	data := loadData(regex)
	ans := parseData(data)
	fmt.Println("The answer to par 1 is: ", ans)
}

func loadData(regex string) [][]string {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}
	var strings [][]string
	defer fi.Close()
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(regex)
		allStrings := re.FindAllString(line, -1)
		strings = append(strings, allStrings)
	}
	return strings
}

func parseData(data [][]string) int {
	ans := 0
	for _, v := range data {
		for _, str := range v {
			after, _ := strings.CutPrefix(str, "mul(")
			final, _ := strings.CutSuffix(after, ")")
			numbers := strings.Split(final, ",")
			leftNumber := transform(numbers[0])
			rightNumber := transform(numbers[1])
			ans += leftNumber * rightNumber
		}
	}
	return ans
}

func transform(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		panic("error converting")
	}
	return number
}
