package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func lcm(x, y int) int {
	return (x * y) / gcd(x, y)
}

func lcmArray(arr []int) int {
	res := arr[0]
	for _, el := range arr {
		res = lcm(res, el)
	}
	return res
}

func solvePart1() {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	m := make(map[string][]string)

	for scanner.Scan() {
		userInput := scanner.Text()
		variableName := userInput[0:3]
		m[variableName] = []string{userInput[7:10], userInput[12:15]}

	}

	i := 0
	ans := 0
	curr := "AAA"
	for curr != "ZZZ" {
		if i == len(instructions) {
			i = 0
		}
		if instructions[i] == 'L' {
			curr = m[curr][0]
		} else {
			curr = m[curr][1]
		}
		i++
		ans++
	}
	fmt.Println(ans)
}

func solvePart2() {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	m := make(map[string][]string)

	for scanner.Scan() {
		userInput := scanner.Text()
		variableName := userInput[0:3]
		m[variableName] = []string{userInput[7:10], userInput[12:15]}

	}

	var res []int
	for key, _ := range m {
		if key[len(key)-1] == 'A' {
			i := 0
			t := 0
			curr := key
			for curr[len(curr)-1] != 'Z' {
				if i == len(instructions) {
					i = 0
				}
				if instructions[i] == 'L' {
					curr = m[curr][0]
				} else {
					curr = m[curr][1]
				}
				i++
				t++
			}
			res = append(res, t)
		}
	}
	fmt.Println(lcmArray(res))
}

func main() {
	solvePart1()
	solvePart2()
}
