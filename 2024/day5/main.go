package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// PartOne()
	PartTwo()
}

type Data struct {
	beforeRules map[int]map[int]bool
	updates     [][]int
}

func PartOne() {

	data := loadData()
	for k, v := range data.beforeRules {
		fmt.Println("Map key:", k)
		fmt.Println("Subseuent before numbers:", len(v))
	}
	ans := parseData(data)
	fmt.Println("The answer to part 1 is:", ans)
}

func loadData() Data {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)

	data := Data{
		beforeRules: make(map[int]map[int]bool),
		updates:     make([][]int, 0),
	}

	beforeMap := make(map[int]map[int]bool)
	flag := false
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			flag = true
			continue
		}
		// false flag, so its rules : for each number, put all numbers that must appear before it
		if !flag {
			line := scanner.Text()
			parts := strings.Split(line, "|")

			// Convert to integers
			beforeValue, err1 := strconv.Atoi(parts[0])
			afterValue, err2 := strconv.Atoi(parts[1])

			if err1 != nil || err2 != nil {
				log.Fatalf("Error converting values: %v, %v", err1, err2)
			}

			// Ensure the map for the after value exists
			if _, exists := beforeMap[afterValue]; !exists {
				beforeMap[afterValue] = make(map[int]bool)
			}

			// Add the before value to the prerequisites of the after value
			beforeMap[afterValue][beforeValue] = true
		} else {
			r := strings.Split(line, ",")
			numbers := make([]int, 0)
			for _, str := range r {
				number, err := strconv.Atoi(str)
				if err != nil {
					panic("Error converting")
				}
				numbers = append(numbers, number)
			}
			data.updates = append(data.updates, numbers)
		}
	}
	data.beforeRules = beforeMap
	return data
}

// The fourth update, 75,97,47,61,53, is not in the correct order: it would print 75 before 97, which violates the rule 97|75.
// idea: maybe keep a map of before, where we map a key and a map of values, where these values are all all values that must appear
// before our current value
// i.e 75|47, 75|61, 75|53, and 75|29. -> this means for 47, 75 must appear before it
// then, we can actually loop the update backwards : for each value except the last value, just check if the values
// to the right of it are supposed to be before it; i.e if we have something like 53|61, then 53 must come before 61.
// if we loop it in reverse, then we see that when we are at 61, then the rule states we have failed since 53 is printed after 61
// then for 47 : we just check if 61 and 53 are supposed to come before it; if not, we are good to go

func parseData(data Data) int {
	ans := 0
	for _, updates := range data.updates {
		flag := true
		evaluatedNumbers := make([]int, 0)
		for i := len(updates) - 1; i >= 0; i-- {
			if i == len(updates)-1 {
				evaluatedNumbers = append(evaluatedNumbers, updates[i])
			} else {
				// otherwise, check all evaluated numbers with the currnet number; if
				if !checkBefore(data.beforeRules, updates[i], evaluatedNumbers) {

					flag = false
					break
				}
				evaluatedNumbers = append(evaluatedNumbers, updates[i])
			}
		}
		if flag {
			middle := updates[len(updates)/2]
			ans += middle
		}
	}
	return ans
}

func checkBefore(rules map[int]map[int]bool, number int, evaluatedNumbers []int) bool {
	// run this number against the rule with all evaluated numbers
	for _, numb := range evaluatedNumbers {

		beforeNumbers, ok := rules[number]
		if ok {
			// that means this number has a rule; so check all the numbers that must come before it
			// if the evaluated number is in there, this means this number shouldve came before our current number
			// in the updates
			_, ok := beforeNumbers[numb]
			if ok {
				return false
			}
		}
	}
	return true
}
