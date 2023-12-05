package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solvePart1() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	g := strings.Fields(scanner.Text())
	var seeds []int
	for _, el := range g[1:] {
		num, _ := strconv.Atoi(el)
		seeds = append(seeds, num)
	}

	var w [][][]int
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		} else if line[0] == 's' || line[0] == 'f' || line[0] == 'l' || line[0] == 'h' || line[0] == 'w' || line[0] == 't' {
			w = append(w, [][]int{})
		} else {
			numbers := strings.Fields(line)
			var a []int
			for _, val := range numbers {
				num, _ := strconv.Atoi(val)
				a = append(a, num)
			}
			w[len(w)-1] = append(w[len(w)-1], a)
		}
	}
	ans := 99999999999
	for _, el := range seeds {
		for _, a := range w {
			for _, b := range a {
				if el >= b[1] && el < b[1]+b[2] {
					el = b[0] + (el - b[1])
					break
				}
			}
		}
		ans = min(ans, el)
	}
	fmt.Println(ans)
}

// todo: create an optimal solution for part2 / the current solution takes more than 2 minutes to finish executing
func solvePart2() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	g := strings.Fields(scanner.Text())
	var seeds []int
	for _, el := range g[1:] {
		num, _ := strconv.Atoi(el)
		seeds = append(seeds, num)
	}

	var w [][][]int
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		} else if line[0] == 's' || line[0] == 'f' || line[0] == 'l' || line[0] == 'h' || line[0] == 'w' || line[0] == 't' {
			w = append(w, [][]int{})
		} else {
			numbers := strings.Fields(line)
			var a []int
			for _, val := range numbers {
				num, _ := strconv.Atoi(val)
				a = append(a, num)
			}
			w[len(w)-1] = append(w[len(w)-1], a)
		}
	}
	ans := 99999999999
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			el := j
			for _, a := range w {
				for _, b := range a {
					if el >= b[1] && el < b[1]+b[2] {
						el = b[0] + (el - b[1])
						break
					}
				}
			}
			ans = min(ans, el)
		}
	}
	fmt.Println(ans)
}

func main() {
	solvePart1()
	solvePart2()
}
