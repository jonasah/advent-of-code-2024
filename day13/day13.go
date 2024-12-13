package day13

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
)

type Machine struct {
	ax, ay, bx, by, px, py float64
}

func Part1(input string) int {
	return solve(input, 0)
}

func Part2(input string) int {
	return solve(input, 10000000000000)
}

func solve(input string, offset float64) int {
	machines := parseInput(input, offset)

	tokens := 0
	for _, m := range machines {
		det := m.ax*m.by - m.bx*m.ay
		numA := (m.by*m.px - m.bx*m.py) / det
		numB := (-m.ay*m.px + m.ax*m.py) / det

		if numA < 0 || numB < 0 {
			continue
		}

		_, fracA := math.Modf(numA)
		_, fracB := math.Modf(numB)

		if fracA < 1e-9 && fracB < 1e-9 {
			tokens += int(numA*3 + numB)
		}
	}

	return tokens
}

func parseInput(input string, offset float64) []Machine {
	inputs := strings.Split(input, "\n\n")

	r := regexp.MustCompile(`X.(\d+), Y.(\d+)`)

	var machines []Machine
	for _, i := range inputs {
		lines := common.GetLines(i)

		a := r.FindStringSubmatch(lines[0])
		ax, _ := strconv.ParseFloat(a[1], 64)
		ay, _ := strconv.ParseFloat(a[2], 64)

		b := r.FindStringSubmatch(lines[1])
		bx, _ := strconv.ParseFloat(b[1], 64)
		by, _ := strconv.ParseFloat(b[2], 64)

		prize := r.FindStringSubmatch(lines[2])
		px, _ := strconv.ParseFloat(prize[1], 64)
		py, _ := strconv.ParseFloat(prize[2], 64)

		machines = append(machines, Machine{ax, ay, bx, by, px + offset, py + offset})
	}

	return machines
}
