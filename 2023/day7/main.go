package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 254583379 wrong
// 254271577 wrong; too high
// 254244561 wrong; too high

var ranking = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 0, // changed to 0 for part 2
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

type Hand struct {
	bid     int
	ranking int
	cards   string
}

type Hands []Hand

func (h Hands) Len() int {
	return len(h)
}

func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hands) Less(i, j int) bool {
	// same rank, so return by card value
	if h[i].ranking == h[j].ranking {
		for k := 0; k < 5; k++ {
			c1 := string(h[i].cards[k])
			c2 := string(h[j].cards[k])
			if ranking[c1] == ranking[c2] {
				continue
			}
			return ranking[c1] < ranking[c2]
		}
	}
	return h[i].ranking < h[j].ranking
}

func main() {
	data := loadData()
	hand := parseData2(data)
	ans := calculateWinnings(hand)
	fmt.Println("The answer is:", ans)
}

func loadData() []string {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)
	data := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	return data
}

func parseData(data []string) []Hand {

	hands := Hands{}

	for _, v := range data {
		line := strings.Split(v, " ")
		bid, err := strconv.Atoi(line[1])
		if err != nil {
			panic("error converting")
		}

		c := Hand{
			bid:     bid,
			ranking: getCardStrength(line[0]),
			cards:   line[0],
		}

		hands = append(hands, c)
	}

	sort.Sort(hands)

	return hands
}

func getCardStrength(cards string) int {
	hand := map[rune]int{}
	for _, c := range cards {
		val, ok := hand[c]
		if !ok {
			hand[c] = 1
		} else {
			hand[c] = val + 1
		}
	}

	// greedy? check rank from strongest to weakest
	// mayb determine based on map size

	// weakest
	if len(hand) == 5 {
		return 1
	}

	// one pair
	if len(hand) == 4 {
		return 2
	}

	// could be two pair, or 3 of a kind
	if len(hand) == 3 {
		for _, v := range hand {
			if v == 3 {
				return 4
			}
		}
		return 3
	}

	// full house or 4 of a kind - size 2
	if len(hand) == 2 {
		for _, v := range hand {
			if v == 4 {
				return 6
			}
		}
		return 5
	}
	// 5 of a kind
	return 7
}

func calculateWinnings(hands Hands) int {

	ans := 0
	for i, h := range hands {
		ans += (i + 1) * h.bid
	}
	return ans
}

// Part 2: with joker

func parseData2(data []string) []Hand {

	hands := Hands{}

	for _, v := range data {
		line := strings.Split(v, " ")
		bid, err := strconv.Atoi(line[1])
		if err != nil {
			panic("error converting")
		}

		c := Hand{
			bid:     bid,
			ranking: getCardStrength2(line[0]),
			cards:   line[0],
		}

		hands = append(hands, c)
	}

	sort.Sort(hands)
	fmt.Println("Sorted hands: ")
	for _, v := range hands {
		fmt.Println(v)
	}
	return hands
}

func getCardStrength2(cards string) int {

	jokers := 0
	hand := map[rune]int{}
	for _, c := range cards {
		val, ok := hand[c]
		if c == 'J' {
			jokers++
		}
		if !ok {
			hand[c] = 1
		} else {
			hand[c] = val + 1
		}
	}
	// weakest
	if len(hand) == 5 {
		switch jokers {
		case 0:
			return 1
		default:
			return 2
		}
	}

	// one pair
	if len(hand) == 4 {
		// ex: jj234
		// ex: j2234
		// this becomes a 3 of a kind or a 2 pair; greedily choose 3 of kind

		if jokers == 1 || jokers == 2 {
			return 4
		}
		return 2
	}

	// could be two pair, or 3 of a kind
	if len(hand) == 3 {
		threeOfKind := false
		for _, v := range hand {
			if v == 3 {
				// QQQJA - 3 of a kind - either 0 or 1 joker
				threeOfKind = true
				break
			}
		}
		// coud be 3334A
		// coud be JJJ23
		// coud be JJ233
		// coud be JJ223
		// coud be J2333
		if threeOfKind {
			if jokers == 0 {
				return 4
			}
			return 6
		}

		// must be 2 pairs
		// KKAA2 -  2 pair - 0 joker
		// KTJJT - 2 pair - 2 joker -> 4 of a kind
		// KTJJT - 2 pair - 2 joker -> 4 of a kind
		// KJTTT
		// KJTTK

		if jokers == 0 {
			return 3
		}

		if jokers == 1 {
			return 5
		}
		return 6
	}

	// full house or 4 of a kind - size 2
	if len(hand) == 2 {

		fourOfaKind := false
		for _, v := range hand {
			if v == 4 {
				fourOfaKind = true
				break
			}
		}

		// JJJJ4 - 4 joker
		// 44445 - 0 joker
		// J4444 - 1 joker
		// JJJ44 - 1 joker
		if fourOfaKind {
			if jokers == 0 {
				return 6
			}
			return 7
		}

		// fullhouse
		// JJJ22 - 3 jokers
		// JJ222 - 2 jokers
		// 33222 - 0 jokers
		if jokers == 0 {
			return 5
		}
		return 7

	}
	// 5 of a kind
	return 7
}
