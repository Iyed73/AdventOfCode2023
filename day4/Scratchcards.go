package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func find(l []int, number int) bool {
	for _, element := range l {
		if element == number {
			return true
		}
	}
	return false
}

func solvePart1() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		w := strings.Fields(line)
		var winningNumbers []int
		var numbers []int
		ok := true
		for _, el := range w {
			if el == "|" {
				ok = false
			} else if ok && isInteger(el) {
				num, _ := strconv.Atoi(el)
				winningNumbers = append(winningNumbers, num)
			} else if isInteger(el) {
				num, _ := strconv.Atoi(el)
				numbers = append(numbers, num)
			}
		}
		x := 0
		for _, el := range winningNumbers {
			if find(numbers, el) {
				x++
			}
		}
		x--
		if x >= 0 {
			ans += int(math.Pow(2.0, float64(x)))
		}
	}
	fmt.Println(ans)
}

func solvePart2() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	ans := 0
	i := 0
	var cards [220]int
	for scanner.Scan() {
		line := scanner.Text()
		w := strings.Fields(line)
		var winningNumbers []int
		var numbers []int
		ok := true
		for _, el := range w {
			if el == "|" {
				ok = false
			} else if ok && isInteger(el) {
				num, _ := strconv.Atoi(el)
				winningNumbers = append(winningNumbers, num)
			} else if isInteger(el) {
				num, _ := strconv.Atoi(el)
				numbers = append(numbers, num)
			}
		}
		x := 0
		for _, el := range winningNumbers {
			if find(numbers, el) {
				x++
			}
		}
		for j := i + 1; j <= i+x; j++ {
			cards[j] += cards[i] + 1
		}
		i++
	}
	for _, el := range cards {
		ans += el
	}
	ans += 219
	fmt.Println(ans)
}

func main() {
	solvePart1()
	solvePart2()
}
