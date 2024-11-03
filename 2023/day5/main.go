package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	dest   int
	source int
	range_ int
}

type Seed struct {
	min int
	max int
}

func main() {
	seeds, nameMaps := loadData2()
	ans := parseSeeds(seeds, nameMaps)
	fmt.Println("lowest location:", ans)
}

func loadData2() ([]Seed, map[string][]Range) {
	fi, err := os.Open("input2.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	row := 0
	scanner := bufio.NewScanner(fi)
	var seeds []Seed
	maps := map[string][]Range{}
	var name string

	for scanner.Scan() {
		// foor each map: we create the struct, then its just na array of
		// range sttructs; we have source start, dest start, and the offest

		line := scanner.Text()
		if row == 0 {
			seeds = getSeeds2(line)
			row += 1
			continue
		}
		// empty line so remove name
		if len(line) == 0 {
			continue
		} else if strings.Contains(line, "map") {
			name = strings.Split(line, " ")[0]
			maps[name] = []Range{}
		} else {
			r := Range{}
			mappings := strings.Split(line, " ")
			for i, v := range mappings {
				integer, err := strconv.Atoi(v)
				if err != nil {
					panic("Error converting in seeds!")
				}
				switch i {
				case 0:
					r.dest = integer
				case 1:
					r.source = integer
				default:
					r.range_ = integer
				}
			}
			ranges, ok := maps[name]
			if ok {
				ranges = append(ranges, r)
				maps[name] = ranges
			}
		}
	}
	return seeds, maps
}

func getSeeds2(line string) []Seed {
	seeds := []Seed{}
	seedArr := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")

	i := 0
	for i < len(seedArr) {
		seed, range_ := seedArr[i], seedArr[i+1]
		seedNumb, err := strconv.Atoi(seed)
		if err != nil {
			panic("Error converting")
		}

		rangeNumb, err := strconv.Atoi(range_)
		if err != nil {
			panic("error converting")
		}

		s := Seed{
			min: seedNumb,
			max: seedNumb + rangeNumb,
		}
		seeds = append(seeds, s)
		i += 2
	}
	fmt.Println("Seeds:", seeds)

	return seeds
}

func parseSeeds(seeds []Seed, maps map[string][]Range) int {
	mappingSequence := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}
	fmt.Println("Seeds:", seeds)

	ans := []int{}

	for _, seed := range seeds {
		a, b := seed.min, seed.max
		for {
			resultA := a
			resultB := b
			for _, s := range mappingSequence {
				resultA = getMapping(resultA, maps[s])
				resultB = getMapping(resultB, maps[s])
			}

			if resultA < resultB {
				ans = append(ans, resultA)
				break
			}
		}

	}
	return ans
}

func getMapping(start int, ranges []Range) int {
	for _, r := range ranges {
		if r.source <= start && start <= r.source+r.range_-1 {
			// apply offset
			return start - r.source + r.dest
		}
	}
	return start
}

func debugPrint(seeds []int, maps map[string][]Range) {
	fmt.Println(seeds)
	for k, v := range maps {
		fmt.Println("Map name:", k)
		for _, value := range v {
			fmt.Println("range struct:", value)
		}
	}
}

// part1

func getSeeds(line string) []int {
	seeds := []int{}
	seedArr := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
	for _, v := range seedArr {
		seedNumb, err := strconv.Atoi(v)
		if err != nil {
			panic("Error converting!")
		}
		seeds = append(seeds, seedNumb)
	}
	return seeds
}

func loadData() ([]int, map[string][]Range) {
	fi, err := os.Open("input2.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	row := 0
	scanner := bufio.NewScanner(fi)
	var seeds []int
	maps := map[string][]Range{}
	var name string

	for scanner.Scan() {
		// foor each map: we create the struct, then its just na array of
		// range sttructs; we have source start, dest start, and the offest

		line := scanner.Text()
		if row == 0 {
			seeds = getSeeds(line)
			row += 1
			continue
		}
		// empty line so remove name
		if len(line) == 0 {
			continue
		} else if strings.Contains(line, "map") {
			name = strings.Split(line, " ")[0]
			maps[name] = []Range{}
		} else {
			r := Range{}
			mappings := strings.Split(line, " ")
			for i, v := range mappings {
				integer, err := strconv.Atoi(v)
				if err != nil {
					panic("Error converting in seeds!")
				}
				switch i {
				case 0:
					r.dest = integer
				case 1:
					r.source = integer
				default:
					r.range_ = integer
				}
			}
			ranges, ok := maps[name]
			if ok {
				ranges = append(ranges, r)
				maps[name] = ranges
			}
		}
	}
	return seeds, maps
}
