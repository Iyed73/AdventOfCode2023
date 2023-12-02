package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func solvePart1() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	id := 1
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		w := strings.Fields(line)
		check := true
		for i, el := range w {
			if isInteger(el) {
				num, _ := strconv.Atoi(el)
				if (num > 12 && strings.Contains(w[i+1], "red")) || (num > 13 && strings.Contains(w[i+1], "green")) || (num > 14 && strings.Contains(w[i+1], "blue")) {
					check = false
					break
				}
			}
		}
		if check {
			ans += id
		}
		id += 1
	}
	fmt.Println(ans)
}

func solvePart2() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		w := strings.Fields(line)
		dic := map[string]int{}
		for i, el := range w {
			if isInteger(el) {
				num, _ := strconv.Atoi(el)
				switch {
				case strings.Contains(w[i+1], "red"):
					dic["red"] = max(num, dic["red"])
				case strings.Contains(w[i+1], "green"):
					dic["green"] = max(num, dic["green"])
				case strings.Contains(w[i+1], "blue"):
					dic["blue"] = max(num, dic["blue"])
				}
			}
		}
		x := dic["blue"] * dic["green"] * dic["red"]
		ans += x
	}
	fmt.Println(ans)
}

func main() {
	solvePart1()
	solvePart2()
}
