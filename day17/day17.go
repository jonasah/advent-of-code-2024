package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
)

type Opcode int

func (o Opcode) String() string {
	return []string{"adv", "bxl", "bst", "jnz", "bxc", "out", "bdv", "cdv"}[o]
}

const (
	adv Opcode = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func Part1(input string) string {
	a, b, c, prog := parseInput(input)
	return strings.Join(slicelib.Itoa(execute(a, b, c, prog)), ",")
}

func Part2(input string) int {
	_, b, c, prog := parseInput(input)

	_, a := findRegisterA(0, len(prog)-1, b, c, prog)
	return a
}

func parseInput(input string) (int, int, int, []int) {
	lines := common.GetLines(input)

	a, _ := strconv.Atoi(strings.Fields(lines[0])[2])
	b, _ := strconv.Atoi(strings.Fields(lines[1])[2])
	c, _ := strconv.Atoi(strings.Fields(lines[2])[2])

	p := slicelib.Atoi(strings.Split(strings.Fields(lines[4])[1], ","))

	return a, b, c, p
}

func execute(a, b, c int, prog []int) []int {
	var output []int

	pc := 0
	for pc < len(prog) {
		opcode := Opcode(prog[pc])
		operand := prog[pc+1]

		switch opcode {
		case adv:
			a >>= combo(operand, a, b, c)
		case bxl:
			b ^= operand
		case bst:
			b = combo(operand, a, b, c) % 8
		case jnz:
			if a != 0 {
				pc = operand - 2
			}
		case bxc:
			b ^= c
		case out:
			output = append(output, combo(operand, a, b, c)%8)
		case bdv:
			c = b >> combo(operand, a, b, c)
		case cdv:
			c = a >> combo(operand, a, b, c)
		default:
			panic(fmt.Sprintf("invalid opcode %d (%s) with operand %d", opcode, opcode, operand))
		}

		pc += 2
	}

	return output
}

func combo(operand, a, b, c int) int {
	return []int{0, 1, 2, 3, a, b, c}[operand]
}

func findRegisterA(base, pos, b, c int, prog []int) (bool, int) {
	if pos < 0 {
		return true, base
	}

	for i := range 8 {
		a := base + i<<(pos*3)

		cmpLen := len(prog) - pos

		output := execute(a, b, c, prog)
		if len(output) < cmpLen {
			continue
		}

		if slices.Equal(output[len(output)-cmpLen:], prog[pos:]) {
			ok, value := findRegisterA(a, pos-1, b, c, prog)
			if ok {
				return ok, value
			}
		}
	}

	return false, -1
}
