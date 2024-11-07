package main

func LoadDataTest() map[string]Monkey {
	return map[string]Monkey{
		"Monkey 0": {
			items:   []int{79, 98},
			ifTrue:  "Monkey 2",
			ifFalse: "Monkey 3",
		},
		"Monkey 1": {
			items:   []int{54, 65, 75, 74},
			ifTrue:  "Monkey 2",
			ifFalse: "Monkey 0",
		},
		"Monkey 2": {
			items:   []int{79, 60, 97},
			ifTrue:  "Monkey 1",
			ifFalse: "Monkey 3",
		},
		"Monkey 3": {
			items:   []int{74},
			ifTrue:  "Monkey 0",
			ifFalse: "Monkey 1",
		},
	}
}

func LoadData() map[string]Monkey {
	return map[string]Monkey{
		"Monkey 0": {
			items:   []int{91, 54, 70, 61, 64, 64, 60, 85},
			ifTrue:  "Monkey 5",
			ifFalse: "Monkey 2",
		},
		"Monkey 1": {
			items:   []int{82},
			ifTrue:  "Monkey 4",
			ifFalse: "Monkey 3",
		},
		"Monkey 2": {
			items:   []int{84, 93, 70},
			ifTrue:  "Monkey 5",
			ifFalse: "Monkey 1",
		},
		"Monkey 3": {
			items:   []int{78, 56, 85, 93},
			ifTrue:  "Monkey 6",
			ifFalse: "Monkey 7",
		},
		"Monkey 4": {
			items:   []int{64, 57, 81, 95, 52, 71, 58},
			ifTrue:  "Monkey 7",
			ifFalse: "Monkey 3",
		},
		"Monkey 5": {
			items:   []int{58, 71, 96, 58, 68, 90},
			ifTrue:  "Monkey 4",
			ifFalse: "Monkey 1",
		},
		"Monkey 6": {
			items:   []int{56, 99, 89, 97, 81},
			ifTrue:  "Monkey 0",
			ifFalse: "Monkey 2",
		},
		"Monkey 7": {
			items:   []int{68, 72},
			ifTrue:  "Monkey 6",
			ifFalse: "Monkey 0",
		},
	}
}
