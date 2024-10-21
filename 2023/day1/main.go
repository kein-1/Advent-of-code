package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// brute force : find index of lowest and highest..?
func main() {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer fi.Close()

	var ans int = 0
	reader := bufio.NewReader(fi)
	for {
		str, err := reader.ReadString(byte('\n'))
		str = strings.TrimSuffix(str, "\n")
		if err != nil {
			if err == io.EOF && len(str) > 0 {
				result := getLowest(str)*10 + getHighest(str)
				ans += result
				fmt.Printf("Str: %s , Calib: %d \n\n", str, result)
				fmt.Println("Reached EOF")
				break
			}
			fmt.Println("read:", str)
		}
		result := getLowest(str)*10 + getHighest(str)
		ans += result
		fmt.Printf("Str: %s , Calib: %d \n\n", str, result)
	}
	fmt.Println("The answer is:", ans)

}

// get the actual string val, check for lowest
func getLowest(str string) int {
	var lowest int = len(str)
	var numbStr string = ""
	strings_ := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, v := range strings_ {
		if index := strings.Index(str, v); index != -1 {
			if index < lowest {
				lowest = index
				numbStr = v
			}
		}
	}
	return convertFormat(numbStr)
}

// get the actual string val, check for lowest
func getHighest(str string) int {
	var highest int = -1
	var numbStr string = ""
	strings_ := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, v := range strings_ {
		if index := strings.LastIndex(str, v); index != -1 {
			if index > highest {
				highest = index
				numbStr = v
			}
		}
	}
	return convertFormat(numbStr)
}

func convertFormat(str string) int {
	stringToNumber := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	value, ok := stringToNumber[str]
	if !ok {
		fmt.Println("key does not exist")
		return 0
	}
	return value
}

// // Part 1
// func main() {
// 	fi, err := os.Open("input.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer fi.Close()

// 	var ans int = 0
// 	reader := bufio.NewReader(fi)
// 	for {
// 		str, err := reader.ReadString(byte('\n'))
// 		str = strings.TrimSuffix(str, "\n")
// 		if err != nil {
// 			if err == io.EOF && len(str) > 0 {
// 				ans += getCalib(str)
// 				fmt.Println("Reached EOF")
// 				break
// 			}
// 			fmt.Println("read:", str)
// 		}
// 		result := getCalib(str)
// 		ans += result
// 		fmt.Printf("Str: %s , Calib: %d \n\n", str, result)
// 	}
// 	fmt.Println("The answer is:", ans)

// }

// func getCalib(str string) int {
// 	ans := 0
// 	var s int = 0
// 	var t int = 0
// 	for _, v := range str {
// 		if v >= '0' && v <= '9' {
// 			s = int(v) - '0'
// 			break
// 		}
// 	}

// 	for i := len(str) - 1; i >= 0; i-- {
// 		if str[i] >= '0' && str[i] <= '9' {
// 			t = int(str[i]) - '0'
// 			break
// 		}
// 	}

// 	fmt.Println("s and t are:", s, t)
// 	ans = s*10 + t

// 	return ans
// }
