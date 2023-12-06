package main

import (
	"fmt"
	"strings"

	"github.com/eliiasg/adventofgo2023/util"
)

func main() {
	fmt.Println("problem 1 result:")
	fmt.Println(prob1(util.GetInput("1.txt")))
	fmt.Println("problem 2 result:")
	fmt.Println(prob2(util.GetInput("1.txt")))
}

func prob1(input string) int {
	s := 0
	for _, line := range strings.Split(input, "\n") {
		r := value(line)
		if r != 0 {
			r = 1 << (r - 1)
		}
		s += r
	}
	return s
}

func value(line string) int {
	s := 0
	line = line[9:]
	tmp := strings.Split(line, " |")
	winning, own := tmp[0], tmp[1]
	for i := 0; i < len(winning); i += 3 {
		num := winning[i : i+3]
		if strings.Contains(own, num) {
			s++
		}
	}
	return s
}

func prob2(input string) int {
	lines := strings.Split(input, "\n")
	values := make([]int, len(lines))
	amt := make([]int, len(lines))
	for i, line := range lines {
		values[i] = value(line)
		amt[i] = 1
	}
	for i := 0; i < len(amt); i++ {
		for j := i + 1; j <= i+values[i]; j++ {
			amt[j] += amt[i]
		}
	}
	s := 0
	for _, e := range amt {
		s += e
	}
	return s
}
