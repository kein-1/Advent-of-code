package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	instructions, mapping := loadData()
	// ans := parseData(instructions, mapping)
	ans := parseData2(instructions, mapping)
	fmt.Println("The answer is:", ans)

}

type Node struct {
	left  string
	right string
}

func loadData() (string, map[string]Node) {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)
	var instructions string
	mapping := map[string]Node{}
	flag := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			flag = false
			continue
		}

		if flag {
			instructions = line
		} else {
			move := strings.Split(line, " = ")
			directions := strings.ReplaceAll(strings.ReplaceAll(move[1], "(", ""), ")", "")
			d := strings.Split(directions, ", ")
			mapping[move[0]] = Node{
				left:  d[0],
				right: d[1],
			}
		}
	}

	return instructions, mapping
}

func parseData(instructions string, mapping map[string]Node) int {

	counter := 0
	ans := 0
	curr := "AAA"
	for {
		index := counter % len(instructions)
		n := mapping[curr]
		if instructions[index] == 'R' {
			curr = n.right
		} else {
			curr = n.left
		}
		ans++
		counter++
		if curr == "ZZZ" {
			break
		}
	}
	return ans
}

// part2
func parseData2(instructions string, mapping map[string]Node) int {

	// get all nodes that end with A, then continuously loop until
	// answerNodes is same length as nodes; after each iteartion, clear out anwer nodes
	// basically can run BFS

	ans := 0
	counter := 0

	queue := []string{}
	nodes := []string{}
	for k, _ := range mapping {
		if k[2] == 'A' {
			queue = append(queue, k)
		}
	}

	startLength := len(queue)
	fmt.Println("queue:", queue)

	for {
		index := counter % len(instructions)
		direction := instructions[index]
		length := len(queue)

		for i := 0; i < length; i++ {
			elem, queue_ := queue[0], queue[1:]
			queue = queue_
			node := mapping[elem]
			var str string
			if direction == 'R' {
				str = node.right
			} else {
				str = node.left
			}

			queue = append(queue, str)
			if str[2] == 'Z' {
				nodes = append(nodes, str)
			}
		}
		ans++
		counter++
		if len(nodes) == startLength {
			break
		} else {
			nodes = []string{}
		}
		fmt.Println("queue:", queue)

	}
	fmt.Println("The final result:", nodes)
	return ans

}
