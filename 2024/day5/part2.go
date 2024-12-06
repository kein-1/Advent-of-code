package main

import (
	"fmt"
	"sort"
)

func PartTwo() {
	data := loadData()
	ans := parseData2(data)
	fmt.Println("The answer to part 2 is:", ans)
}

func parseData2(data Data) int {
	ans := 0

	for _, updates := range data.updates {
		evaluatedNumbers := make([]int, 0)
		for i := len(updates) - 1; i >= 0; i-- {
			if i == len(updates)-1 {
				evaluatedNumbers = append(evaluatedNumbers, updates[i])
			} else {

				if !checkBefore(data.beforeRules, updates[i], evaluatedNumbers) {

					// Incorrect order, so make this correct and then
					// get its value
					ans += orderUpdates(data.beforeRules, updates)

					break
				}
				evaluatedNumbers = append(evaluatedNumbers, updates[i])
			}
		}

	}
	return ans
}

// probably topological sort in part 2 to re-order
func orderUpdates(rules map[int]map[int]bool, updates []int) int {
	// Initialize degrees map
	degrees := make(map[int]int)

	// Compute in-degrees and count dependencies
	for _, current := range updates {
		for _, other := range updates {
			if current == other {
				continue
			}
			befores := rules[current]
			_, ok := befores[other]
			if ok {
				degrees[current]++
			}
		}
	}
	// Sort updates based on in-degrees
	sort.Slice(updates, func(i, j int) bool {
		return degrees[updates[i]] < degrees[updates[j]]
	})

	return updates[len(updates)/2]

}
