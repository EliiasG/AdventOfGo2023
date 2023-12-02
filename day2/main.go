package main

import (
	"fmt"
	"strconv"
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
	for i, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		var ln string
		if i < 10 {
			ln = line[8:]
		} else {
			ln = line[9:]
		}
		if isValid(ln) {
			s += (i + 1)
		}
	}
	return s
}

func isValid(line string) bool {
	sections := strings.Split(line, "; ")
	for _, section := range sections {
		colors := strings.Split(section, ", ")
		for _, color := range colors {
			res := strings.Split(color, " ")
			amt, col := res[0], res[1]
			a, _ := strconv.ParseInt(amt, 10, 32)
			if col == "red" && a > 12 {
				return false
			}
			if col == "green" && a > 13 {
				return false
			}
			if col == "blue" && a > 14 {
				return false
			}
		}
	}
	return true
}

func power(line string) int {
	var r, g, b int
	sections := strings.Split(line, "; ")
	for _, section := range sections {
		colors := strings.Split(section, ", ")
		for _, color := range colors {
			res := strings.Split(color, " ")
			if len(res) != 2 {
				return -1
			}
			amt, col := res[0], res[1]
			a, _ := strconv.ParseInt(amt, 10, 32)
			if col == "red" {
				r = max(r, int(a))
			}
			if col == "green" {
				g = max(g, int(a))
			}
			if col == "blue" {
				b = max(b, int(a))
			}
		}
	}
	return r * g * b
}

func prob2(input string) int {
	s := 0
	for i, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		var ln string
		if i < 9 {
			ln = line[8:]
		} else if i < 99 {
			ln = line[9:]
		} else {
			ln = line[10:]
		}
		s += power(ln)
	}
	return s
}
