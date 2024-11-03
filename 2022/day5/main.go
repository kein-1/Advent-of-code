package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	quantity int
	start    int
	end      int
}

func main() {
	boxes, moves := loadData()
	// computeMoves(boxes, moves)
	computeMoves2(boxes, moves)
	ans := getStr(boxes)
	fmt.Println("The str is: ", ans)
}

func loadData() (map[int][]string, []Move) {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)

	boxes := map[int][]string{
		1: {"F", "C", "J", "P", "H", "T", "W"},
		2: {"G", "R", "V", "F", "Z", "J", "B", "H"},
		3: {"H", "P", "T", "R"},
		4: {"Z", "S", "N", "P", "H", "T"},
		5: {"N", "V", "F", "Z", "H", "J", "C", "D"},
		6: {"P", "M", "G", "F", "W", "D", "Z"},
		7: {"M", "V", "Z", "W", "S", "J", "D", "P"},
		8: {"N", "D", "S"},
		9: {"D", "Z", "S", "F", "M"},
	}
	moves := []Move{}
	flag := false

	for scanner.Scan() {

		line := scanner.Text()
		if len(line) == 0 {
			flag = true
			continue
		}
		if flag {
			l := strings.Split(line, " ")
			temp := []int{}
			for _, v := range l {
				v_, err := strconv.Atoi(v)
				if err != nil {
					continue
				}
				temp = append(temp, v_)
			}

			m := Move{
				quantity: temp[0],
				start:    temp[1],
				end:      temp[2],
			}
			moves = append(moves, m)
		}
	}
	return boxes, moves
}

func computeMoves(boxes map[int][]string, moves []Move) {
	for _, v := range moves {
		for i := 0; i < v.quantity; i++ {

			from := boxes[v.start]
			to := boxes[v.end]

			last := from[len(from)-1]
			boxes[v.start] = from[:len(from)-1]
			to = append(to, last)
			boxes[v.end] = to
		}
	}
}

func computeMoves2(boxes map[int][]string, moves []Move) {
	for _, v := range moves {

		from := boxes[v.start]
		to := boxes[v.end]
		last := from[len(from)-v.quantity:]
		boxes[v.start] = from[:len(from)-v.quantity]
		to = append(to, last...)
		boxes[v.end] = to
	}
}

func getStr(boxes map[int][]string) string {
	str := ""
	for i := 1; i <= 9; i++ {
		v := boxes[i]
		last := v[len(v)-1]
		str += last
	}
	return str
}
