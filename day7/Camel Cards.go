package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// todo: improve code quality

func jokerProcess(r, pair, three, four, five int) int {
	if r == 1 && four == 1 {
		five = 1
		four = 0
	} else if r == 1 && three == 1 {
		four = 1
		three = 0
	} else if r == 2 && three == 1 {
		five = 1
		three = 0
		pair = 0
	} else if r == 1 && pair >= 1 {
		three = 1
		pair--
	} else if r == 2 && pair == 2 {
		four = 1
		pair = 0
	} else if r == 3 && pair == 1 {
		five = 1
		pair = 0
		three = 0
	} else if r == 2 {
		three = 1
		pair = 0
	} else if r == 4 {
		five = 1
		four = 0
	} else if r == 3 {
		four = 1
		three = 0
	} else if r == 1 {
		pair++
	}
	return five*1000 + four*100 + three*10 + pair*1
}

func calculate(card string, joker bool) int {
	m := make(map[rune]int)
	for _, el := range card {
		m[el]++
	}
	five := 0
	four := 0
	three := 0
	pair := 0
	for _, value := range m {
		switch value {
		case 5:
			five++
		case 4:
			four++
		case 3:
			three++
		case 2:
			pair++
		}
	}
	if joker {
		return jokerProcess(m['J'], pair, three, four, five)

	}
	return five*1000 + four*100 + three*10 + pair*1
}

func check(card1, card2 string, joker bool) bool {
	if calculate(card1, joker) > calculate(card2, joker) {
		return true
	} else if calculate(card1, joker) < calculate(card2, joker) {
		return false
	} else {
		dic := map[rune]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
		if joker {
			dic['J'] = 1
		}
		for i := 0; i < 5; i++ {
			if dic[rune(card1[i])] > dic[rune(card2[i])] {
				return true
			} else if dic[rune(card1[i])] < dic[rune(card2[i])] {
				return false
			}
		}
	}
	return true
}

func solvePart1() {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var hands []string
	var bids []int
	for scanner.Scan() {
		line := scanner.Text()
		w := strings.Fields(line)
		hand := w[0]
		bid, _ := strconv.Atoi(w[1])
		hands = append(hands, hand)
		bids = append(bids, bid)
	}

	ok := true
	for ok {
		ok = false
		for i := 0; i < len(hands)-1; i++ {
			if check(hands[i], hands[i+1], false) {
				hands[i], hands[i+1] = hands[i+1], hands[i]
				bids[i], bids[i+1] = bids[i+1], bids[i]
				ok = true
			}
		}
	}

	ans := 0
	for i := 0; i < 1000; i++ {
		ans += (i + 1) * bids[i]
	}
	fmt.Println(ans)
}

func solvePart2() {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var hands []string
	var bids []int
	for scanner.Scan() {
		line := scanner.Text()
		w := strings.Fields(line)
		hand := w[0]
		bid, _ := strconv.Atoi(w[1])
		hands = append(hands, hand)
		bids = append(bids, bid)
	}

	ok := true
	for ok {
		ok = false
		for i := 0; i < len(hands)-1; i++ {
			if check(hands[i], hands[i+1], true) {
				hands[i], hands[i+1] = hands[i+1], hands[i]
				bids[i], bids[i+1] = bids[i+1], bids[i]
				ok = true
			}
		}
	}

	ans := 0
	for i := 0; i < 1000; i++ {
		ans += (i + 1) * bids[i]
	}
	fmt.Println(ans)
}

func main() {
	solvePart1()
	solvePart2()
}
