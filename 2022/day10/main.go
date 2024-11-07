package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	name  string
	value int
}

type Sprite struct {
	a, b, c int
}

func main() {

	data := loadData()
	// ans := parseData(data)
	ans := parseData2(data)
	fmt.Println("Answer is: ", ans)
}

func loadData() []Instruction {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)

	instructions := []Instruction{}
	for scanner.Scan() {

		line := scanner.Text()
		var i Instruction
		if line == "noop" {
			i = Instruction{name: "noop"}
		} else {
			l_ := strings.Split(line, " ")

			val, err := strconv.Atoi(l_[1])
			if err != nil {
				panic("Error conversion!")
			}

			i = Instruction{
				name:  l_[0],
				value: val,
			}
		}
		instructions = append(instructions, i)
	}
	return instructions
}

func parseData(instructions []Instruction) int {
	// each iteration in instruction, if we see noop, just increment cycle
	// then we should check a separate queue to see if we need to process;
	// if we see a addx V, we add this to the queue and increment the cycle by 1

	// ex:
	// 	noop
	// addx 3
	// addx -5

	// noop - do nothing, cycle++(2). do not check queue since nothing in there
	// addx3 - counter++(3). do not check queue since nothing there; add this to queue
	// addx5 - counter++(4). remove from queue and add 3; now we have V = 4. add addx5 to queue
	// no mroe instructions, but we have queue still; remove from queue

	ans := 0
	counter := 0
	register := 1
	cycle := 1
	flag := false
	for counter < len(instructions) {

		instruction := instructions[counter]
		fmt.Println("Instruction is: ", instruction)
		if instruction.name != "noop" {
			if flag {
				fmt.Printf("In second tick. Before update: Register %d || cycle %d || ans %d \n", register, cycle, ans)
				register += instruction.value
				flag = false
				counter++
				fmt.Printf("In second tick. After update: Register %d || cycle %d || ans %d \n", register, cycle, ans)
			} else {
				fmt.Printf("In first tick. Register %d || cycle %d || ans %d \n", register, cycle, ans)
				flag = true
			}
		} else {
			counter++
			fmt.Printf("In noop. nothing happening here. Register %d || cycle %d || ans %d \n", register, cycle, ans)
		}

		fmt.Println("")

		cycle++
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			ans += register * cycle
		}

	}

	fmt.Println("Final cycle and register: ", cycle, register)

	return ans
}

// part2

func parseData2(instructions []Instruction) int {
	// each iteration in instruction, if we see noop, just increment cycle
	// then we should check a separate queue to see if we need to process;
	// if we see a addx V, we add this to the queue and increment the cycle by 1

	// ex:
	// 	noop
	// addx 3
	// addx -5

	// noop - do nothing, cycle++(2). do not check queue since nothing in there
	// addx3 - counter++(3). do not check queue since nothing there; add this to queue
	// addx5 - counter++(4). remove from queue and add 3; now we have V = 4. add addx5 to queue
	// no mroe instructions, but we have queue still; remove from queue
	s := Sprite{
		a: 0,
		b: 1,
		c: 2,
	}
	ans := 0
	counter := 0
	register := 1
	cycle := 1
	flag := false
	var sb strings.Builder
	for counter < len(instructions) {

		instruction := instructions[counter]

		if cycle-1 == s.a || cycle-1 == s.b || cycle-1 == s.c {
			sb.WriteString("#")
		} else {
			sb.WriteString(".")
		}
		// fmt.Println("Current cycle , register, and str:", cycle, register, sb.String())
		// fmt.Println("s is: ", s)
		// fmt.Println("")

		if instruction.name != "noop" {
			if flag {
				// fmt.Printf("In second tick. Before update: Register %d || cycle %d || ans %d \n", register, cycle, ans)
				register += instruction.value
				flag = false
				counter++

				s.a = register - 1
				s.b = register
				s.c = register + 1

				// fmt.Printf("In second tick. After update: Register %d || cycle %d || ans %d \n", register, cycle, ans)
			} else {
				// fmt.Printf("In first tick. Register %d || cycle %d || ans %d \n", register, cycle, ans)
				flag = true
			}
		} else {
			counter++
			// fmt.Printf("In noop. nothing happening here. Register %d || cycle %d || ans %d \n", register, cycle, ans)
		}
		cycle++

		if cycle == 41 {
			fmt.Println(sb.String())

			sb.Reset()
			cycle = 1
		}

	}

	return ans
}
