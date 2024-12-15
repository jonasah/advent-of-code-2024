package day14

import (
	"fmt"
	"regexp"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/mathlib"
	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
	"golang.org/x/exp/slices"
)

type Robot struct {
	pos mathlib.Position
	dir mathlib.Direction
}

func Part1(input string, width, height int) int {
	robots := parseInput(input)
	steps := 100

	xc := width / 2
	yc := height / 2
	quadrants := make([]int, 4)

	for _, r := range robots {
		x := (r.pos.X() + r.dir.X()*steps) % width
		if x < 0 {
			x += width
		}

		y := (r.pos.Y() + r.dir.Y()*steps) % height
		if y < 0 {
			y += height
		}

		if x < xc && y < yc {
			quadrants[0]++
		} else if x < xc && y > yc {
			quadrants[1]++
		} else if x > xc && y < yc {
			quadrants[2]++
		} else if x > xc && y > yc {
			quadrants[3]++
		}
	}

	rating := 1
	for _, count := range quadrants {
		rating *= count
	}

	return rating
}

func Part2(input string, width, height int) int {
	robots := parseInput(input)
	steps := 6516

	for _, r := range robots {
		newX := (r.pos.X() + r.dir.X()*steps) % width
		if newX < 0 {
			newX += width
		}

		newY := (r.pos.Y() + r.dir.Y()*steps) % height
		if newY < 0 {
			newY += height
		}

		r.pos = mathlib.NewPosition(newX, newY)
	}

	drawRobots(robots, width, height)

	return steps
}

func parseInput(input string) []*Robot {
	lines := common.GetLines(input)

	r := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	var robots []*Robot
	for _, l := range lines {
		matches := r.FindStringSubmatch(l)
		v := slicelib.Atoi(matches[1:])
		robots = append(robots, &Robot{mathlib.NewPosition(v[0], v[1]), mathlib.NewDirection(v[2], v[3])})
	}

	return robots
}

func drawRobots(robots []*Robot, width, height int) {
	for y := range height {
		for x := range width {
			p := mathlib.NewPosition(x, y)
			containsRobot := slices.ContainsFunc(robots, func(r *Robot) bool { return r.pos == p })
			if containsRobot {
				fmt.Print("o")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}
