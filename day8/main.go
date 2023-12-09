package main

import (
	"fmt"
	"strings"

	"github.com/eliiasg/adventofgo2023/util"
)

func parseInput(lines []string) map[string][2]string {
	res := make(map[string][2]string)
	for i := 2; i < len(lines); i++ {
		sp := strings.Split(strings.TrimSpace(lines[i]), " = ")
		rs := strings.Split(sp[1][1:len(sp[1])-1], ", ")
		res[sp[0]] = [2]string{rs[0], rs[1]}
	}
	return res
}

func main() {
	fmt.Println("problem 1 result:")
	fmt.Println(prob1(util.GetInput("1.txt")))
	fmt.Println("problem 2 result:")
	fmt.Println(prob2(util.GetInput("1.txt")))
}

func prob1(input string) int {
	return -1
}

func prob2(input string) int {
	lines := strings.Split(input, "\n")
	xmap := parseInput(lines)
	nodes := make([]string, 0, len(xmap))
	for k := range xmap {
		if k[2] == 'A' {
			nodes = append(nodes, k)
		}
	}
	ln := len(nodes)
	inst := strings.TrimSpace(lines[0])
	instl := len(inst)
	var c byte
	var allZ bool
	var node string
	i := 0
	j := 0
	for {
		c = inst[i%instl]
		allZ = true
		for j = 0; j < ln; j++ {
			node = nodes[j]
			if c == 'L' {
				nodes[j] = xmap[node][0]
			} else if c == 'R' {
				nodes[j] = xmap[node][1]
			}
			if node[2] != 'Z' {
				allZ = false
			}
		}
		i++
		if allZ {
			return i
		}
	}
}
