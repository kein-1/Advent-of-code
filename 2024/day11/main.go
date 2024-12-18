package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	run()
}

func run() {

	stones := loadData()
	ans := parseData2(stones)
	// fmt.Println("part 1 answer: ", ans)

	fmt.Println("The answer to part 2 is:", ans)
}

func loadData() []string {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)

	stones := make([]string, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		stones = append(stones, line...)
	}

	return stones
}

// probably just at each iteartion, loop through our array , apply rules, and make a new one. Afterwards, set current array
// to the new array
func parseData(stones []string) int {
	i := 0

	for i < 25 {
		newStones := make([]string, 0)
		for _, stone := range stones {
			if stone == "0" {
				newStones = append(newStones, "1")
			} else if len(stone)%2 == 0 {
				leftHalf := stone[:len(stone)/2]
				rightHalf := parseStone(stone[len(stone)/2:])
				newStones = append(newStones, leftHalf, rightHalf)
			} else {
				i, _ := strconv.Atoi(stone)
				i *= 2024
				newStones = append(newStones, strconv.Itoa(i))
			}
		}
		stones = newStones
		i++
	}

	return len(stones)
}

// need some form of cache?
func parseData2(stones []string) int {

	// used to cache the reuslt of splitting a string so we dont need to recompute it
	cacheStrings := make(map[string][]string)

	i := 0

	cache := make(map[string]int, 0)
	for _, stone := range stones {
		cache[stone]++
	}

	for i < 75 {
		newCache := make(map[string]int, 0)
		for stone, count := range cache {
			updatedStone := applyRule(stone, cacheStrings)
			for _, v := range updatedStone {
				newCache[v] += count
			}
		}
		cache = newCache
		i++

	}
	ans := 0
	for _, v := range cache {
		ans += v
	}

	return ans
}

func applyRule(stone string, cacheStrings map[string][]string) []string {

	// maps to 1
	if stone == "0" {
		return []string{"1"}
	} else if len(stone)%2 == 0 {
		// return the split result so we dont need to recompute it
		value, ok := cacheStrings[stone]
		if ok {
			return value
		}
		leftHalf := stone[:len(stone)/2]
		rightHalf := parseStone(stone[len(stone)/2:])
		cacheStrings[stone] = []string{leftHalf, rightHalf}
		return []string{leftHalf, rightHalf}
	}
	i, _ := strconv.Atoi(stone)
	i *= 2024

	return []string{strconv.Itoa(i)}

}

func parseStone(stone string) string {
	for len(stone) > 1 && stone[0] == '0' {
		stone = stone[1:]
	}
	return stone
}
