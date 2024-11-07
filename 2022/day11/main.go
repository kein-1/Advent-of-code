package main

import (
	"fmt"
)

type Monkey struct {
	items   []int
	ifTrue  string
	ifFalse string
}

func main() {

	data := LoadDataTest()
	ans := parseData(data)
	fmt.Println("Answer is: ", ans)
}

func parseData(monkeys map[string]Monkey) int {

	// inspections := map[string]int{
	// 	"Monkey 0": 0,
	// 	"Monkey 1": 0,
	// 	"Monkey 2": 0,
	// 	"Monkey 3": 0,
	// 	"Monkey 4": 0,
	// 	"Monkey 5": 0,
	// 	"Monkey 6": 0,
	// 	"Monkey 7": 0,
	// }

	// names := []string{
	// 	"Monkey 0",
	// 	"Monkey 1",
	// 	"Monkey 2",
	// 	"Monkey 3",
	// 	"Monkey 4",
	// 	"Monkey 5",
	// 	"Monkey 6",
	// 	"Monkey 7",
	// }

	inspections := map[string]int{
		"Monkey 0": 0,
		"Monkey 1": 0,
		"Monkey 2": 0,
		"Monkey 3": 0,
	}

	names := []string{
		"Monkey 0",
		"Monkey 1",
		"Monkey 2",
		"Monkey 3",
	}

	fmt.Println(monkeys)

	for i := 0; i < 1000; i++ {
		// for each monkey:
		for _, name := range names {
			// inspect items
			v := monkeys[name]
			items := v.items
			if len(items) == 0 {
				continue
			}
			length := inspections[name]
			length += len(items)
			inspections[name] = length

			for len(items) > 0 {
				x := items[0]
				items = items[1:]
				// wLevel := GetOperation(name, x) / 3/ 3
				wLevel := GetOperationTest(name, x)
				if GetThrowRuleTest(name, wLevel) {

					monkey := monkeys[v.ifTrue]
					monkey.items = append(monkey.items, wLevel)
					monkeys[v.ifTrue] = monkey
				} else {
					monkey := monkeys[v.ifFalse]
					monkey.items = append(monkey.items, wLevel)
					monkeys[v.ifFalse] = monkey
				}
			}
			v.items = items
			monkeys[name] = v
		}

		fmt.Printf("End of round %d \n", i+1)
		// for _, v := range names {
		// 	fmt.Printf("%s holds : %v \n", v, monkeys[v].items)
		// 	fmt.Println("")
		// }

	}
	fmt.Println("insepctions: ", inspections)

	a, b := 0, 0
	for _, v := range inspections {
		if v > a {
			b = a
			a = v
		} else if v > a {
			a = v
		}
	}

	fmt.Println("a,b: ", a, b)

	return a * b
}
