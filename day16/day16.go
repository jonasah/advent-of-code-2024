package day16

import (
	"maps"
	"math"
	"slices"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/mathlib"
)

const (
	start = 'S'
	end   = 'E'
	wall  = '#'
)

type Maze struct {
	start, end mathlib.Position
	walls      map[mathlib.Position]bool
}

type Path struct {
	pos     mathlib.Position
	dir     mathlib.Direction
	visited map[mathlib.Position]bool
	score   int
}

func Part1(input string) int {
	result, _ := solve(input)
	return result
}

func Part2(input string) int {
	_, result := solve(input)
	return result
}

func solve(input string) (int, int) {
	maze := parseMap(input)

	paths := []Path{
		{
			pos:     maze.start,
			dir:     mathlib.Right,
			visited: make(map[mathlib.Position]bool),
			score:   0,
		},
	}

	cache := map[mathlib.Direction]map[mathlib.Position]int{
		mathlib.Right: make(map[mathlib.Position]int),
		mathlib.Left:  make(map[mathlib.Position]int),
		mathlib.Up:    make(map[mathlib.Position]int),
		mathlib.Down:  make(map[mathlib.Position]int),
	}

	bestScore := math.MaxInt
	var bestPathsVisited map[mathlib.Position]bool

	for len(paths) > 0 {
		path := paths[0]
		paths = paths[1:]

		bestScoreHere, hasVisited := cache[path.dir][path.pos]
		if path.score > bestScore || hasVisited && path.score > bestScoreHere {
			continue
		}

		cache[path.dir][path.pos] = path.score

		path.visited[path.pos] = true

		if path.pos == maze.end {
			if path.score < bestScore {
				bestScore = path.score
				bestPathsVisited = make(map[mathlib.Position]bool)
			}

			maps.Copy(bestPathsVisited, path.visited)
			continue
		}

		candidates := []struct {
			dir  mathlib.Direction
			cost int
		}{
			{path.dir, 1},
			{path.dir.TurnLeft(), 1001},
			{path.dir.TurnRight(), 1001},
		}

		for _, c := range candidates {
			next := path.pos.Move(c.dir)
			if !path.visited[next] && !maze.walls[next] {
				paths = append(paths, Path{
					pos:     next,
					dir:     c.dir,
					visited: maps.Clone(path.visited),
					score:   path.score + c.cost,
				})
			}
		}

		slices.SortFunc(paths, func(a, b Path) int { return est(a, maze) - est(b, maze) })
	}

	return bestScore, len(bestPathsVisited)
}

func parseMap(input string) Maze {
	lines := common.GetLines(input)

	maze := Maze{walls: make(map[mathlib.Position]bool)}
	for y, l := range lines {
		for x, t := range l {
			p := mathlib.NewPosition(x, y)

			switch t {
			case 'S':
				maze.start = p
			case 'E':
				maze.end = p
			case '#':
				maze.walls[p] = true
			}
		}
	}

	return maze
}

func est(path Path, maze Maze) int {
	e := mathlib.ManhattanDistance(path.pos, maze.end)

	if path.dir == mathlib.Right {
		e += 1000
	}
	if path.dir == mathlib.Left {
		e += 2000
	}
	if path.dir == mathlib.Up {
		e += 0
	}
	if path.dir == mathlib.Down {
		e += 2000
	}

	return e
}
