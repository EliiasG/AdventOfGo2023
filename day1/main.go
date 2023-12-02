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
	var s int64
	for _, line := range strings.Split(input, "\n") {
		f := strings.IndexAny(line, "0123456789")
		l := strings.LastIndexAny(line, "0123456789")
		r, _ := strconv.ParseInt(string(line[f])+string(line[l]), 10, 64)
		s += r
	}
	return int(s)
}

var values = map[string]int{
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func prob2(input string) int {
	s := 0
	for _, line := range strings.Split(input, "\n") {
		keys := make([]string, 0, len(values))
		for k := range values {
			keys = append(keys, k)
		}
		s += values[firstKeyword(line, keys)]*10 + values[lastKeyword(line, keys)]
	}
	return s
}

func firstKeyword(s string, keywords []string) string {
	idx := len(s)
	out := ""
	for _, keyword := range keywords {
		r := strings.Index(s, keyword)
		if r == -1 || r > idx {
			continue
		}
		idx = r
		out = keyword
	}
	return out
}

func lastKeyword(s string, keywords []string) string {
	idx := 0
	out := ""
	for _, keyword := range keywords {
		r := strings.LastIndex(s, keyword)
		if r == -1 || r < idx {
			continue
		}
		idx = r
		out = keyword
	}
	return out
}
