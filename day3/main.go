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
	lines := strings.Split(input, "\n")
	s := 0
	for i, line := range lines {
		curNum := ""
		hasSymbol := false
		for j, c := range line + "." {
			if util.IsDigit(c) {
				curNum += string(c)
				for k := i - 1; k <= i+1; k++ {
					if k < 0 || k >= len(lines) {
						continue
					}
					for l := j - 1; l <= j+1; l++ {
						if l < 0 || l >= len(line) {
							continue
						}
						check := rune(lines[k][l])
						if check != '.' && !util.IsDigit(check) {
							hasSymbol = true
						}
					}
				}
			} else {
				if hasSymbol {
					r, _ := strconv.ParseInt(curNum, 10, 32)
					s += int(r)
				}
				curNum = ""
				hasSymbol = false
			}
		}
	}
	return s
}

func prob2(input string) int {
	s := 0
	nums := genNumbers(input)
	for i, line := range strings.Split(input, "\n") {
		for j, c := range line {
			if c != '*' {
				continue
			}
			var r1, r2 *int
			for k := i - 1; k <= i+1; k++ {
				if k < 0 || k >= len(nums) {
					continue
				}
				for l := j - 1; l <= j+1; l++ {
					if l < 0 || l >= len(nums[k]) {
						continue
					}
					r := nums[k][l]
					if r1 == nil {
						r1 = r
					} else if r1 != r && r2 == nil {
						r2 = r
					}
				}
			}
			if r1 != nil && r2 != nil {
				s += *r1 * *r2
				fmt.Println(*r1, *r2)
			}
		}
	}
	return s
}

func genNumbers(input string) [][]*int {
	lines := strings.Split(input, "\n")
	res := make([][]*int, len(lines))
	for i, line := range lines {
		res[i] = make([]*int, len(line))
		cur := new(int)
		for j, c := range line + "." {
			if util.IsDigit(c) {
				*cur *= 10
				*cur += int(c - '0')
				res[i][j] = cur
			} else {
				cur = new(int)
			}
		}
	}
	return res
}
