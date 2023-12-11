package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solvePart1() {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		var w []int
		var r [][]int

		for _, val := range values {
			num, _ := strconv.Atoi(val)
			w = append(w, num)
		}
		ans := w[len(w)-1]

		for {
			zeroCount := 0
			var d []int

			for j := 0; j < len(w)-1; j++ {
				d = append(d, w[j+1]-w[j])
				if d[j] == 0 {
					zeroCount++
				}
			}

			w = d
			r = append(r, w)
			ans += w[len(w)-1]

			if zeroCount == len(w) {
				break
			}
		}
		res += ans

	}

	fmt.Println(res)
}

func solvePart2() {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	res := 0

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		var w []int

		for _, val := range values {
			num, _ := strconv.Atoi(val)
			w = append(w, num)
		}

		var r []int
		r = append(r, w[0])
		for {
			zeroCount := 0
			var d []int

			for j := 0; j < len(w)-1; j++ {
				d = append(d, w[j+1]-w[j])
				if d[j] == 0 {
					zeroCount++
				}
			}

			w = d
			r = append(r, w[0])

			if zeroCount == len(w) {
				break
			}
		}
		ans := 0
		for j := len(r) - 1; j >= 0; j-- {
			ans = r[j] - ans
		}
		res += ans
	}

	fmt.Println(res)
}

func main() {
	solvePart1()
	solvePart2()
}
