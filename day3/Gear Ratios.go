package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var directions = [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

func solvePart1() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	matrix := make([][]rune, 150)
	for i := range matrix {
		matrix[i] = make([]rune, 150)
		for j := range matrix[i] {
			matrix[i][j] = '.'
		}
	}

	k := 1
	for scanner.Scan() {
		line := scanner.Text()
		for j := 1; j <= len(line); j++ {
			matrix[k][j] = rune(line[j-1])
		}
		k++
	}

	ans := 0
	ok := false
	curr := 0
	for i := 1; i <= 140; i++ {
		for j := 1; j <= 140; j++ {
			if j == 1 && ok {
				ok = false
				ans += curr
				curr = 0
			}
			if j == 1 {
				curr = 0
			}
			if '0' <= matrix[i][j] && matrix[i][j] <= '9' {
				curr = curr*10 + int(matrix[i][j]-'0')
				for _, el := range directions {
					r := matrix[el[0]+i][el[1]+j]
					if r != '.' && !unicode.IsDigit(r) && !unicode.IsLetter(r) {
						ok = true
					}
				}
			} else {
				if ok {
					ok = false
					ans += curr
				}
				curr = 0
			}
		}
	}
	fmt.Println(ans)
}

func find(l []int, number int) bool {
	for _, element := range l {
		if element == number {
			return true
		}
	}
	return false
}

func solvePart2() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	matrix := make([][]rune, 150)
	numbers := make([][]int, 150)
	for i := range matrix {
		matrix[i] = make([]rune, 150)
		numbers[i] = make([]int, 150)
		for j := range matrix[i] {
			matrix[i][j] = '.'
			numbers[i][j] = 0
		}
	}

	k := 1
	for scanner.Scan() {
		line := scanner.Text()
		for j := 1; j <= len(line); j++ {
			matrix[k][j] = rune(line[j-1])
		}
		k++
	}

	first := -1
	var curr string
	for i := 1; i <= 140; i++ {
		for j := 1; j <= 140; j++ {
			if first != -1 && j == 1 {
				for k := first; k < first+len(curr); k++ {
					numbers[i-1][k], _ = strconv.Atoi(curr)
				}
				first = -1
				curr = ""
			}
			if '0' <= matrix[i][j] && matrix[i][j] <= '9' {
				if first == -1 {
					first = j
				}
				curr += string(matrix[i][j])
			} else {
				if first != -1 {
					for k := first; k < first+len(curr); k++ {
						numbers[i][k], _ = strconv.Atoi(curr)
					}
				}
				curr = ""
				first = -1
			}
		}
	}

	ans := 0
	for i := 1; i <= 140; i++ {
		for j := 1; j <= 140; j++ {
			if matrix[i][j] == '*' {
				var w []int
				for _, el := range directions {
					num := numbers[el[0]+i][el[1]+j]
					if num != 0 && !find(w, num) {
						w = append(w, num)
					}
				}
				if len(w) == 2 {
					ans += w[0] * w[1]
				}
			}
		}
	}
	fmt.Println(ans)
}

func main() {
	solvePart1()
	solvePart2()
}
