package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	direction string
	count     int
}

type Spot struct {
	x int
	y int
}

// 9719 - incorrect
// 438 - incorrect
func main() {

	moves := loadData()

	grid := make([][]int, 10000)
	for i := range 1000 {
		grid[i] = make([]int, 10000)
	}

	grid[500][500] = 2

	// tailSpots := simulate(moves, grid)

	tailSpots := simulate2(moves)
	fmt.Println("Tail spots:", len(tailSpots))

}

func loadData() []Move {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)

	moves := []Move{}
	for scanner.Scan() {
		line := scanner.Text()
		l_ := strings.Split(line, " ")

		move, err := strconv.Atoi(l_[1])
		if err != nil {
			panic("failed conversion")
		}
		m := Move{
			direction: l_[0],
			count:     move,
		}

		moves = append(moves, m)
	}
	return moves
}

// seems like a greedy problem; to touch, we want either same row or column but if not, then diagon
// so we can check at end of each move from head if we are touching; if not, then determine the move for tail
// brute force seems to be just create a grid as large as possible then start head and tail in the center
// use 1 for tail, 2 for head; each time head/tail moves, just mark the grid spot with its value
// what we can do is keep a hashmap of all the places that are unique

func simulate(moves []Move, grid [][]int) map[Spot]bool {

	spots := map[Spot]bool{}

	s := Spot{
		x: 500,
		y: 500,
	}
	spots[s] = true
	// start of h and t
	x := 500
	y := 500

	a := 500
	b := 500

	for _, m := range moves {
		dir, count := m.direction, m.count
		for i := 0; i < count; i++ {
			grid[x][y] = 0
			switch dir {
			case "U":
				x--
			case "D":
				x++
			case "L":
				y--
			default:
				y++
			}
			grid[x][y] = 2
			// If fail, we need to move tail and add its spot
			// either move tail in same row, same col, or diagon

			if !checkRange(grid, a, b) {
				// updates tail to tag head
				moveDirection(grid, x, y, &a, &b)

				s := Spot{
					x: a,
					y: b,
				}
				spots[s] = true
			}
		}
	}
	return spots

}

// Checks if tail is touching head or not
func checkRange(grid [][]int, a int, b int) bool {

	// we are overlapping
	if grid[a][b] == 2 {
		return true
	}

	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
	}

	for _, v := range directions {
		x, y := v[0], v[1]
		newX := a + x
		newY := b + y
		if grid[newX][newY] == 2 {
			return true
		}
	}
	return false
}

func moveDirection(grid [][]int, x int, y int, a *int, b *int) {
	// in same row
	if x == *a {
		// head behind tail; move tail
		if y < *b {
			*b = *b - 1
		} else {
			// head ahead of tail
			*b = *b + 1
		}
	} else if y == *b {
		// head abvoe tail; move tail up
		if x < *a {
			*a = *a - 1
		} else {
			// head below tail, move ail down
			*a = *a + 1
		}
	} else {
		// diagonal; so we simulate a move to see if it will be touching

		diagonals := [][]int{
			{-1, -1},
			{-1, 1},
			{1, 1},
			{1, -1},
		}

		for _, v := range diagonals {
			u, v := v[0], v[1]
			newX := *a + u
			newY := *b + v
			if checkRange(grid, newX, newY) {
				*a = newX
				*b = newY
				return
			}
		}
	}

}

// Part 2 :
// Seems like if a knot moves, we just need to check if the next knot needs to move;
// if not, we can finish there
// otherwise, we move that knot to touch the leader knot, then check whether the next knot
// needs to do the same. continue until we reach knot 9, which is tail
// so what we can do: store a map of position 0-9, where 1-9 are the remaining knots;then
// just loop and move. might not even need a grid here;

func simulate2(moves []Move) map[Spot]bool {

	positions := map[int]Spot{
		0: {x: 10, y: 10},
		1: {x: 10, y: 10},
		2: {x: 10, y: 10},
		3: {x: 10, y: 10},
		4: {x: 10, y: 10},
		5: {x: 10, y: 10},
		6: {x: 10, y: 10},
		7: {x: 10, y: 10},
		8: {x: 10, y: 10},
		9: {x: 10, y: 10},
	}

	spots := map[Spot]bool{}

	s := Spot{
		x: 10,
		y: 10,
	}

	spots[s] = true

	for _, m := range moves {
		dir, count := m.direction, m.count
		for i := 0; i < count; i++ {

			h := positions[0]

			switch dir {
			case "U":
				h.x--
			case "D":
				h.x++
			case "L":
				h.y--
			default:
				h.y++
			}
			positions[0] = h

			// now check for each subsequent
			for i := 1; i <= 9; i++ {
				leader := positions[i-1]
				follower := positions[i]
				// if false, we need to move follower to close to leader

				result := checkRange2(leader, follower)
				fmt.Println("leader", leader)
				fmt.Println("follower", follower)
				fmt.Println("result:", result)
				if result {
					break
				}
				fmt.Println("Moving!")
				move(leader, &follower)
				if i == 9 {
					spots[follower] = true
				}
				positions[i] = follower

				fmt.Println("follower is now:", follower)
				fmt.Println("")
			}
			fmt.Println(" ")
			// debugGrid(positions)
		}

		fmt.Println("Current positions: ")
		fmt.Println(positions)
		fmt.Println(spots)
		fmt.Println(" ")
	}
	return spots
}

// checks the range; if false, we need to move follower
func checkRange2(leader Spot, follower Spot) bool {
	a, b := leader.x, leader.y
	u, v := follower.x, follower.y

	if a == b && u == v {
		return true
	}

	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
	}

	for _, dir := range directions {
		x, y := dir[0], dir[1]
		newX := u + x
		newY := v + y

		if newX == a && newY == b {
			return true
		}
	}
	return false

}

// Move follower to be touching leader
func move(leader Spot, follower *Spot) {
	leaderRow, leaderCol := leader.x, leader.y
	followerRow, followerCol := follower.x, follower.y

	// if
	if leaderRow == followerRow {
		// follower behind leader; move follower right
		if leaderCol < followerCol {
			followerCol--
		} else {
			// follower ahead of leader; move follower left
			followerCol++
		}
	} else if leaderCol == followerCol {
		// follower above leader; move follower down
		if leaderRow < followerRow {
			followerRow--
		} else {
			// follower below leader; move follower down
			followerRow++
		}
	} else {
		// diagonal; so we simulate a move to see if it will be touching

		diagonals := [][]int{
			{-1, -1},
			{-1, 1},
			{1, 1},
			{1, -1},
		}

		for _, v_ := range diagonals {
			m, n := v_[0], v_[1]
			x := followerRow + m
			y := followerCol + n

			s := Spot{
				x: x,
				y: y,
			}

			if checkRange2(leader, s) {
				follower.x = s.x
				follower.y = s.y
				return
			}
		}
	}

	follower.x = followerRow
	follower.y = followerCol
}

func debugGrid(positions map[int]Spot) {
	grid := make([][]int, 50)
	for i := range 50 {
		grid[i] = make([]int, 50)
	}

	for k, v := range positions {
		x, y := v.x, v.y
		grid[x][y] = k
	}

	for x, row := range grid {
		for y, value := range row {
			if grid[x][y] == 0 {
				fmt.Printf(" . ")
			} else {
				fmt.Printf("%3d ", value) // %3d aligns values in a 3-character-wide field
			}
		}
		fmt.Println() // Move to the next line after each row
	}

}
