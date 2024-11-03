package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Balls struct {
	red   int
	blue  int
	green int
}

func main() {

	fmt.Println("Game id power set sum:", sumPowerSet())
}

func sumPowerSet() int {

	ans := 0
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Split(text, ":")

		minBalls := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}

		games := strings.Split(strings.Trim(line[1], " "), ";")
		for _, v := range games {
			game := strings.Split(strings.Trim(v, " "), ",")

			var balls Balls
			for _, k := range game {
				value := strings.Split(strings.Trim(k, " "), " ")
				num, err := strconv.Atoi(value[0])
				if err != nil {
					panic("failed conversion")
				}
				switch value[1] {
				case "red":
					balls.red = num
				case "green":
					balls.green = num
				default:
					balls.blue = num
				}
			}
			check(balls, minBalls)
		}
		ans += minBalls["red"] * minBalls["green"] * minBalls["blue"]
	}
	return ans
}

func check(balls Balls, minBalls map[string]int) {

	if balls.red > minBalls["red"] {
		minBalls["red"] = balls.red
	}
	if balls.green > minBalls["green"] {
		minBalls["green"] = balls.green
	}
	if balls.blue > minBalls["blue"] {
		minBalls["blue"] = balls.blue
	}
}

// func sumGameID() int {

// 	ans := 0
// 	fi, err := os.Open("input.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer fi.Close()

// 	scanner := bufio.NewScanner(fi)
// 	for scanner.Scan() {
// 		text := scanner.Text()
// 		line := strings.Split(text, ":")
// 		id_line := strings.Split(line[0], " ")

// 		id, err := strconv.Atoi(id_line[1])
// 		if err != nil {
// 			panic("failed conversion")
// 		}

// 		games := strings.Split(strings.Trim(line[1], " "), ";")

// 		flag := true
// 		fmt.Println("game id:", id_line[1])
// 		for _, v := range games {
// 			game := strings.Split(strings.Trim(v, " "), ",")

// 			counter := map[string]int{
// 				"red":   0,
// 				"green": 0,
// 				"blue":  0,
// 			}
// 			for _, k := range game {
// 				value := strings.Split(strings.Trim(k, " "), " ")
// 				num, err := strconv.Atoi(value[0])
// 				if err != nil {
// 					panic("failed conversion")
// 				}
// 				counter[value[1]] = num
// 			}
// 			if !check(counter) {
// 				flag = false
// 			}
// 		}
// 		if flag {
// 			ans += id
// 		}
// 	}
// 	return ans
// }
