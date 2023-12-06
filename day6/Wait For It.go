package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solvePart1() {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var time []int
	var distance []int

	scanner.Scan()
	line1 := strings.Fields(scanner.Text())
	for _, el := range line1[1:] {
		num, _ := strconv.Atoi(el)
		time = append(time, num)
	}

	scanner.Scan()
	line2 := strings.Fields(scanner.Text())
	for _, el := range line2[1:] {
		num, _ := strconv.Atoi(el)
		distance = append(distance, num)
	}

	ans := 1
	for i := 0; i < len(time); i++ {
		x := 0
		for j := 1; j < time[i]; j++ {
			r := (time[i] - j) * j
			if r > distance[i] {
				x += 1
			}
		}
		ans *= x
	}
	fmt.Println(ans)
}

func solvePart2() {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var time, distance int
	var stringTime, stringDistance string

	scanner.Scan()
	line1 := strings.Fields(scanner.Text())
	for _, el := range line1[1:] {
		stringTime += el
	}
	time, _ = strconv.Atoi(stringTime)

	scanner.Scan()
	line2 := strings.Fields(scanner.Text())
	for _, el := range line2[1:] {
		stringDistance += el
	}
	distance, _ = strconv.Atoi(stringDistance)

	ans := 0
	for j := 1; j < time; j++ {
		r := (time - j) * j
		if r > distance {
			ans += 1
		}
	}
	fmt.Println(ans)
}

func main() {
	solvePart1()
	solvePart2()
}
