package main

import (
	"bufio"
	"fmt"
	"os"
)

func solvePart1() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		var first, last int
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				first = int(line[i] - '0')
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				last = int(line[i] - '0')
				break
			}
		}
		ans += first*10 + last
	}
	fmt.Println(ans)
}

func solvePart2() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	ans := 0
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for scanner.Scan() {
		line := scanner.Text()
		first, last := -1, -1
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				first = int(line[i] - '0')
			}
			for j, el := range numbers {
				if len(el)+i <= len(line) {
					num := line[i : i+len(el)]
					if num == el {
						first = j + 1
						break
					}
				}
			}
			if first != -1 {
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				last = int(line[i] - '0')
			}
			for j, el := range numbers {
				if len(el)+i <= len(line) {
					num := line[i : i+len(el)]
					if num == el {
						last = j + 1
						break
					}
				}
			}
			if last != -1 {
				break
			}
		}
		ans += first*10 + last
	}
	fmt.Println(ans)
}

func main() {
	solvePart1()
	solvePart2()
}
