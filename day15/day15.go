package day15

import (
	"fmt"
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/mathlib"
)

type Tile rune
type Move rune

const (
	robot    Tile = '@'
	box      Tile = 'O'
	wall     Tile = '#'
	empty    Tile = '.'
	boxLeft  Tile = '['
	boxRight Tile = ']'

	left  Move = '<'
	right Move = '>'
	up    Move = '^'
	down  Move = 'v'
)

var moveToDirMap = map[Move]mathlib.Direction{
	left:  mathlib.Left,
	right: mathlib.Right,
	up:    mathlib.Up,
	down:  mathlib.Down,
}

type Map [][]Tile

func (m Map) tile(p mathlib.Position) Tile {
	return m[p.Y()][p.X()]
}

func (m Map) setTile(p mathlib.Position, v Tile) {
	m[p.Y()][p.X()] = v
}

func Part1(input string) int {
	return solve(input, 1)
}

func Part2(input string) int {
	return solve(input, 2)
}

func solve(input string, part int) int {
	m, robotPos, moves := parseInput(input, part)

	for _, move_ := range moves {
		d := moveToDirMap[move_]

		if canMove(m, robotPos, d) {
			move(m, robotPos, d)
			robotPos = robotPos.Move(d)
		}
	}

	sum := 0
	for y, row := range m {
		for x, t := range row {
			if t == box || t == boxLeft {
				sum += 100*y + x
			}
		}
	}

	return sum
}

func parseInput(input string, part int) (Map, mathlib.Position, []Move) {
	parts := strings.Split(input, "\n\n")

	var m Map
	var robotPos mathlib.Position
	for y, l := range common.GetLines(parts[0]) {
		var row []Tile
		for x, t := range []Tile(l) {
			if t == robot {
				robotPos = mathlib.NewPosition(x*part, y)
				t = empty
			}

			if part == 1 {
				row = append(row, t)
			} else {
				if t == wall || t == empty {
					row = append(row, t, t)
				} else {
					row = append(row, boxLeft, boxRight)
				}
			}
		}

		m = append(m, row)
	}

	moves := []Move(strings.ReplaceAll(parts[1], "\n", ""))

	return m, robotPos, moves
}

func canMove(m Map, p mathlib.Position, d mathlib.Direction) bool {
	return tryMove(m, p, d, false)
}

func move(m Map, p mathlib.Position, d mathlib.Direction) {
	tryMove(m, p, d, true)
}

func tryMove(m Map, p mathlib.Position, d mathlib.Direction, doMove bool) bool {
	next := p.Move(d)

	if m.tile(next) == wall {
		return false
	}

	if m.tile(next) == empty {
		return true
	}

	if m.tile(next) == box {
		canMove := tryMove(m, next, d, doMove)

		if canMove && doMove {
			m.setTile(next, empty)
			m.setTile(next.Move(d), box)
		}

		return canMove
	}

	// found two-wide box

	if d == mathlib.Right {
		boxLeftPos := next
		boxRightPos := boxLeftPos.Move(mathlib.Right)
		canMove := tryMove(m, boxRightPos, d, doMove)

		if canMove && doMove {
			m.setTile(boxLeftPos, empty)
			m.setTile(boxRightPos, boxLeft)
			m.setTile(boxRightPos.Move(mathlib.Right), boxRight)
		}

		return canMove
	}

	if d == mathlib.Left {
		boxRightPos := next
		boxLeftPos := boxRightPos.Move(mathlib.Left)
		canMove := tryMove(m, boxLeftPos, d, doMove)

		if canMove && doMove {
			m.setTile(boxRightPos, empty)
			m.setTile(boxLeftPos, boxRight)
			m.setTile(boxLeftPos.Move(mathlib.Left), boxLeft)
		}

		return canMove
	}

	// moving up or down

	if m.tile(next) == boxLeft {
		boxLeftPos := next
		boxRightPos := boxLeftPos.Move(mathlib.Right)
		canMove := tryMove(m, boxLeftPos, d, doMove) && tryMove(m, boxRightPos, d, doMove)

		if canMove && doMove {
			m.setTile(boxLeftPos, empty)
			m.setTile(boxRightPos, empty)
			m.setTile(boxLeftPos.Move(d), boxLeft)
			m.setTile(boxRightPos.Move(d), boxRight)
		}

		return canMove
	}

	// m.tile(next) == boxRight
	boxRightPos := next
	boxLeftPos := boxRightPos.Move(mathlib.Left)
	canMove := tryMove(m, boxLeftPos, d, doMove) && tryMove(m, boxRightPos, d, doMove)

	if canMove && doMove {
		m.setTile(boxLeftPos, empty)
		m.setTile(boxRightPos, empty)
		m.setTile(boxLeftPos.Move(d), boxLeft)
		m.setTile(boxRightPos.Move(d), boxRight)
	}

	return canMove
}

func printMap(m Map, robotPos mathlib.Position, move rune) {
	var sb strings.Builder
	for y, row := range m {
		if y == robotPos.Y() {
			sb.WriteString(string(row[:robotPos.X()]))
			sb.WriteString(string(robot))
			sb.WriteString(string(row[robotPos.X()+1:]))
		} else {
			sb.WriteString(string(row))
		}
		sb.WriteString("\n")
	}

	fmt.Println("move", string(move))
	fmt.Println(sb.String())
}
