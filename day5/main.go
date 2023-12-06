package main

import (
	"fmt"
	"math"
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
	lines := strings.Split(input, "\n")
	seedStrs := strings.Split(lines[0][7:], " ")
	values := make([]int64, len(seedStrs))
	for i, str := range seedStrs {
		values[i], _ = strconv.ParseInt(str, 10, 64)
	}
	for _, section := range strings.Split(input, "\n\n")[1:] {
		changed := make([]bool, len(values))
		for _, line := range strings.Split(section, "\n")[1:] {
			dat := strings.Split(line, " ")
			dstStart, _ := strconv.ParseInt(dat[0], 10, 64)
			srcStart, _ := strconv.ParseInt(dat[1], 10, 64)
			amt, _ := strconv.ParseInt(dat[2], 10, 64)
			for i, val := range values {
				if val < srcStart || val > srcStart+amt || changed[i] {
					continue
				}
				values[i] = val - srcStart + dstStart
				changed[i] = true
			}
		}
	}
	mn := int64(math.MaxInt64)
	for _, val := range values {
		if val < mn {
			mn = val
		}
	}
	return int(mn)
}

func prob2(input string) int {
	lines := strings.Split(input, "\n")
	seedStrs := strings.Split(lines[0][7:], " ")
	values := make([]int64, 0)
	valset := make(map[int64]bool)
	for i := 0; i < len(seedStrs); i += 2 {
		start, _ := strconv.ParseInt(seedStrs[i], 10, 32)
		amt, _ := strconv.ParseInt(seedStrs[i+1], 10, 32)
		for j := start; j <= start+amt; j++ {
			if valset[j] {
				continue
			}
			values = append(values, j)
			valset[j] = true
		}
		fmt.Println(i)
	}
	for i, section := range strings.Split(input, "\n\n")[1:] {
		changed := make([]bool, len(values))
		for j, line := range strings.Split(section, "\n")[1:] {
			fmt.Println(i, j)
			dat := strings.Split(line, " ")
			dstStart, _ := strconv.ParseInt(dat[0], 10, 64)
			srcStart, _ := strconv.ParseInt(dat[1], 10, 64)
			amt, _ := strconv.ParseInt(dat[2], 10, 64)
			for i, val := range values {
				if val < srcStart || val > srcStart+amt || changed[i] {
					continue
				}
				values[i] = val - srcStart + dstStart
				changed[i] = true
			}
		}
	}
	mn := int64(math.MaxInt64)
	for _, val := range values {
		if val < mn {
			mn = val
		}
	}
	return int(mn)
}
