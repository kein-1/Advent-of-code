package main

import (
	"fmt"
	"strings"
)

func PartTwo() {
	regex := `mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don\'t\(\)`
	data := loadData(regex)
	ans := parseData2(data)
	fmt.Println("The answer to par 2 is: ", ans)
}

func parseData2(data [][]string) int {
	ans := 0
	flag := true
	for _, v := range data {
		for _, str := range v {
			if str == "do()" {
				flag = true
				continue
			}
			if str == "don't()" {
				flag = false
				continue
			}
			if flag {
				after, _ := strings.CutPrefix(str, "mul(")
				final, _ := strings.CutSuffix(after, ")")
				numbers := strings.Split(final, ",")
				leftNumber := transform(numbers[0])
				rightNumber := transform(numbers[1])
				ans += leftNumber * rightNumber
			}
		}
	}
	return ans
}
