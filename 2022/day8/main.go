package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Size struct {
	treeVal  int
	maxUp    int
	maxDown  int
	maxLeft  int
	maxRight int
}

type View struct {
	treeVal    int
	upTrees    int
	downTrees  int
	leftTrees  int
	rightTrees int
}

func main() {

	x := 99
	y := 99

	trees := loadData(x, y)
	// maxArr := parseData(trees, x, y)
	// ans := computeVisibleTrees(maxArr, x, y)
	view := parseData2(trees, x, y)
	ans := computeMaxScenicScore(view, x, y)
	fmt.Println("Trees : ", ans)
}

func loadData(x int, y int) [][]int {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)
	data := make([][]int, x)
	for i := range x {
		data[i] = make([]int, y)
	}

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, v := range line {
			data[row][i] = int(v - '0')
		}
		row += 1
	}
	return data
}

// Brute force
// maybe at each index, we get its biggest left and right values
// same thing : its top and bottom values;
// for left : if current is bigger than biggest in left, then we can def see the tree
// for right: if current is bigger than biggest in right, we can def see tree
func parseData(trees [][]int, x int, y int) [][]Size {

	// get max in each row and each column for each value
	// we can store this in a 2D array; each slot will be a Size struct
	// representing the max left, max right, max up, max down of this tree
	// then we can determine whether the tree is visible

	size := make([][]Size, x)
	for i := range x {
		size[i] = make([]Size, y)
	}

	for i := range x {
		for j := range y {
			tree := trees[i][j]

			s := Size{
				treeVal:  tree,
				maxUp:    getMaxCol(trees, 0, i, j),
				maxDown:  getMaxCol(trees, i+1, x, j),
				maxLeft:  getMaxRow(trees[i], 0, j),
				maxRight: getMaxRow(trees[i], j+1, y),
			}
			size[i][j] = s
		}
	}
	return size
}

func computeVisibleTrees(size [][]Size, x int, y int) int {
	ans := 0
	for i := range x {
		for j := range y {
			// edges
			if i == 0 || j == 0 || i == x-1 || j == y-1 {
				ans += 1
				continue
			}
			s := size[i][j]
			if s.treeVal > s.maxLeft || s.treeVal > s.maxRight || s.treeVal > s.maxUp || s.treeVal > s.maxDown {
				fmt.Println("This tree is visible. its val is:", s.treeVal)
				fmt.Println("row and col:", i, j)
				fmt.Println("maxUp:", s.maxUp)
				fmt.Println("maxDown", s.maxDown)
				fmt.Println("maxLeft:", s.maxLeft)
				fmt.Println("maxRight:", s.maxRight)
				fmt.Println("")
				ans += 1
			}
		}
	}
	return ans
}

func getMaxRow(data []int, start int, end int) int {

	ans := -1
	for i := start; i < end; i++ {
		if data[i] > ans {
			ans = data[i]
		}
	}

	if ans == -1 {
		return math.MaxInt
	}

	return ans
}

func getMaxCol(data [][]int, start int, end int, col int) int {
	ans := -1
	for i := start; i < end; i++ {
		if data[i][col] > ans {
			ans = data[i][col]
		}
	}
	if ans == -1 {
		return math.MaxInt
	}
	return ans
}

// part 2 : we can store the max value on each side, plus its distance

func parseData2(trees [][]int, x int, y int) [][]View {

	view := make([][]View, x)
	for i := range x {
		view[i] = make([]View, y)
	}

	for i := range x {
		for j := range y {
			tree := trees[i][j]
			s := View{
				treeVal:    tree,
				upTrees:    getTreeSeenColumn(trees, 0, i-1, j, tree, "up"),
				downTrees:  getTreeSeenColumn(trees, i+1, x, j, tree, "down"),
				leftTrees:  getTreeSeenRow(trees[i], 0, j-1, tree, "left"),
				rightTrees: getTreeSeenRow(trees[i], j+1, y, tree, "right"),
			}
			view[i][j] = s
		}
	}
	return view
}

func computeMaxScenicScore(view [][]View, x int, y int) int {
	ans := math.MinInt

	for i := range x {
		for j := range y {

			if i == 0 || j == 0 || i+1 == x || j+1 == y {
				continue
			}

			v := view[i][j]

			score := v.upTrees * v.downTrees * v.leftTrees * v.rightTrees

			result := fmt.Sprintf("This tree is at row %d and col %d and its value is %d", i, j, v.treeVal)
			fmt.Println(result)
			fmt.Println("The up seen is: ", v.upTrees)
			fmt.Println("The down seen is: ", v.downTrees)
			fmt.Println("The left seen is: ", v.leftTrees)
			fmt.Println("The right seen is: ", v.rightTrees)
			fmt.Println("The score is: ", score)
			fmt.Println("")
			if score > ans {
				ans = score
			}
		}
	}
	return ans
}

func getTreeSeenRow(data []int, start int, end int, tree int, dir string) int {
	ans := 0

	// expand right to left
	if dir == "left" {
		for i := end; i >= start; i-- {
			ans++
			if data[i] >= tree {
				return ans
			}
		}
	} else {
		// expand left to right
		for i := start; i < end; i++ {
			ans++
			if data[i] >= tree {
				return ans
			}
		}
	}

	return ans
}

func getTreeSeenColumn(data [][]int, start int, end int, col int, tree int, dir string) int {
	ans := 0

	// expand upwards
	if dir == "up" {
		for i := end; i >= start; i-- {
			ans++
			if data[i][col] >= tree {
				return ans
			}
		}
	} else {
		// expand downwards
		for i := start; i < end; i++ {
			ans++
			if data[i][col] >= tree {
				return ans
			}
		}
	}

	return ans
}
