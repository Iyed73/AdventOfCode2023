package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(x int) int {
	file, err := os.Open("day11/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var matrix [][]rune
	var positions [][]int
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		var w []rune
		for j, el := range line {
			w = append(w, el)
			if el == '#' {
				positions = append(positions, []int{i, j})
			}
		}
		i++
		matrix = append(matrix, w)
	}

	var c []int
	var r []int
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		ok := true
		for j := 0; j < m; j++ {
			if matrix[i][j] == '#' {
				ok = false
				break
			}
		}
		if ok {
			r = append(r, i)
		}
	}
	for i := 0; i < m; i++ {
		ok := true
		for j := 0; j < n; j++ {
			if matrix[j][i] == '#' {
				ok = false
				break
			}
		}
		if ok {
			c = append(c, i)
		}
	}

	ans := 0
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			x1, y1 := positions[i][0], positions[i][1]
			x2, y2 := positions[j][0], positions[j][1]
			rCount := 0
			for _, el := range r {
				if min(x1, x2) <= el && el <= max(x1, x2) {
					rCount++
				}
			}

			cCount := 0
			for _, el := range c {
				if min(y1, y2) <= el && el <= max(y1, y2) {
					cCount++
				}
			}

			ans += x*rCount + x*cCount + max(x1, x2) - min(x1, x2) + max(y1, y2) - min(y1, y2)
		}
	}
	return ans
}

func solvePart1() {
	fmt.Println(solve(1))
}

func solvePart2() {
	fmt.Println(solve(99999))
}

func main() {
	solvePart1()
	solvePart2()
}
