package main

func GetOperation(name string, value int) int {
	switch name {
	case "Monkey 0":
		return value * 13
	case "Monkey 1":
		return value + 7
	case "Monkey 2":
		return value + 2
	case "Monkey 3":
		return value * 2
	case "Monkey 4":
		return value * value
	case "Monkey 5":
		return value + 6
	case "Monkey 6":
		return value + 1
	default:
		return value + 8
	}
}

func GetThrowRule(name string, value int) bool {
	switch name {
	case "Monkey 0":
		return value%2 == 0
	case "Monkey 1":
		return value%13 == 0
	case "Monkey 2":
		return value%5 == 0
	case "Monkey 3":
		return value%3 == 0
	case "Monkey 4":
		return value%11 == 0
	case "Monkey 5":
		return value%17 == 0
	case "Monkey 6":
		return value%7 == 0
	default:
		return value%19 == 0
	}
}

func GetOperationTest(name string, value int) int {
	switch name {
	case "Monkey 0":
		return value * 19
	case "Monkey 1":
		return value + 6
	case "Monkey 2":
		return value * value
	default:
		return value + 3
	}
}

func GetThrowRuleTest(name string, value int) bool {
	switch name {
	case "Monkey 0":
		return value%23 == 0
	case "Monkey 1":
		return value%19 == 0
	case "Monkey 2":
		return value%13 == 0
	default:
		return value%17 == 0
	}
}
