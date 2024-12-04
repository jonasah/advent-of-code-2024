package day4

import (
	"slices"
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
)

func Part1(input string) int {
	rows := common.GetLines(input)

	toSearch := slices.Clone(rows)
	toSearch = append(toSearch, getColumns(rows)...)
	toSearch = append(toSearch, getDiagonals(rows)...)
	toSearch = append(toSearch, getDiagonals(slicelib.Map(rows, reverseString))...)

	wordCount := 0
	for _, s := range toSearch {
		wordCount += strings.Count(s, "XMAS")
		wordCount += strings.Count(s, "SAMX")
	}

	return wordCount
}

func Part2(input string) int {
	lines := common.GetLines(input)

	h := len(lines)
	w := len(lines[0])

	count := 0
	for c := 0; c < w-2; c++ {
		for r := 0; r < h-2; r++ {
			diag1 := string(lines[r][c]) + string(lines[r+1][c+1]) + string(lines[r+2][c+2])
			diag2 := string(lines[r+2][c]) + string(lines[r+1][c+1]) + string(lines[r][c+2])

			if (diag1 == "MAS" || diag1 == "SAM") && (diag2 == "MAS" || diag2 == "SAM") {
				count++
			}
		}
	}

	return count
}

func getColumns(rows []string) []string {
	var cols []string
	for c := range rows[0] {
		var sb strings.Builder
		for r := range rows {
			sb.WriteByte(rows[r][c])
		}

		cols = append(cols, sb.String())
	}
	return cols
}

func getDiagonals(rows []string) []string {
	var diag []string
	for c := range rows[0] {
		var sb strings.Builder
		for r := range rows {
			if c+r < len(rows[r]) {
				sb.WriteByte(rows[r][c+r])
			}
		}

		diag = append(diag, sb.String())
	}

	for r := range rows {
		if r == 0 {
			continue
		}

		var sb strings.Builder
		for c := range rows[0] {
			if c+r < len(rows) {
				sb.WriteByte(rows[r+c][c])
			}
		}

		diag = append(diag, sb.String())
	}

	return diag
}

func reverseString(s string) string {
	var result string
	for _, c := range s {
		result = string(c) + result
	}

	return result
}
