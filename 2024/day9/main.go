package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	PartOne()
}

func PartOne() {

	data := loadData()
	fmt.Println("Data is:", data)
	ans := parseData2(data)
	fmt.Println("The answer to part 1 is:", ans)
}

func loadData() []int {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)

	data := make([]int, 0)
	// 23 33 13 31 21 41 41 31 40 2
	idNumb := 0
	flag := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			digit := (line[i] - '0')
			data = populateData(idNumb, int(digit), data, flag)
			flag++
			if flag%2 == 0 {
				idNumb++
			}
		}
	}

	return data
}

func populateData(id int, digit int, data []int, flag int) []int {

	// if flag is even, then it must be file size; otheriwse, it must be free space

	i := 0
	for i < digit {
		if flag%2 == 0 {
			data = append(data, id)
		} else {
			data = append(data, -1)
		}
		i++

	}

	return data
}

// use two pointer to move stuff
func parseData(data []int) int {

	// move blocks
	start, end := 0, len(data)-1
	for {
		// find first empty space
		for start < len(data) && data[start] != -1 {
			start++
		}

		// find first block
		for end >= 0 && data[end] == -1 {
			end--
		}

		if start >= end {
			break
		}

		data[start], data[end] = data[end], data[start]
		start++
		end--
	}

	ans := 0
	for i, v := range data {
		if v == -1 {
			continue
		}
		ans += i * v
	}
	return ans
}

// now moving blocks; so what we need to do is find the size of the file we want to move, then check offset
// maybe we brute force : for each file, just start scanning from left to see if we can find a spot for it

// alternative solution : run two heaps or some sort where we store positions of the largest space that can fit the largest id size.
// more complex to implement
func parseData2(data []int) int {

	end := len(data) - 1
	for end >= 0 {
		start := 0
		// find first block
		for end >= 0 && data[end] == -1 {
			end--
		}
		if end < 0 {
			break
		}

		fileID := data[end]
		size := getFileSize(data, fileID, end)
		// returns where  start is (the first available empty spot that can fit)
		cursor, available := checkSpace(data, start, size, end)
		if available {
			moveContent(data, size, cursor, end, fileID)
		}
		end -= size
	}

	ans := 0
	for i, v := range data {

		if v == -1 {
			continue
		}
		ans += i * v
	}
	return ans
}

// get file size; can find it based on the id
func getFileSize(data []int, id int, position int) int {
	counter := 0
	for position >= 0 && data[position] == id {
		position--
		counter++
	}
	return counter
}

// find the first available space
func checkSpace(data []int, start int, size int, end int) (int, bool) {
	for start < end {
		i := 0
		flag := true
		// check offset first
		for i < size {
			if data[start+i] != -1 {
				flag = false
				break
			}
			i++
		}
		if flag {
			return start, true
		}
		// otherwise, move start to start + i, since this is the first spot where we found aother file. then, we find the next empty space from here
		start += i
		for start < len(data) && data[start] != -1 {
			start++
		}
	}
	return start, false
}

func moveContent(data []int, size int, start int, end int, fileID int) {
	i := 0
	for i < size {
		data[start+i] = fileID
		data[end] = -1
		end--
		i++
	}
}
