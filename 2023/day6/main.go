package main

import (
	"fmt"
	"os"
)

type Race struct {
	time     int
	distance int
}

func main() {
	data := loadData2()
	ans := parseData(data)
	fmt.Println("The answer is:", ans)
}

func loadData() []Race {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	// scanner := bufio.NewScanner(fi)
	// data := []Race{
	// 	{time: 7, distance: 9},
	// 	{time: 15, distance: 40},
	// 	{time: 30, distance: 200},
	// }
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	data2 := []Race{
		{time: 34, distance: 204},
		{time: 90, distance: 1713},
		{time: 89, distance: 1210},
		{time: 86, distance: 1780},
	}

	return data2
}

func parseData(races []Race) int {
	// speed would be hold time/s. travel would be
	// time - hold time * speed
	// brute force : run loop from [0 to time], inclusive

	ans := 1
	for _, race := range races {

		count := 0

		for i := 0; i <= race.time; i++ {
			speed := i
			travelTime := race.time - i
			distance := speed * travelTime
			if distance > race.distance {
				count++
			}
		}
		ans *= count
	}
	return ans

}

// part2
func loadData2() []Race {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	// scanner := bufio.NewScanner(fi)
	// data := []Race{
	// 	{time: 71530, distance: 940200},
	// }
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	data2 := []Race{
		{time: 34908986, distance: 204171312101780},
	}

	return data2
}
