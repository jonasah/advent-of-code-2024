package slicelib_test

import (
	"strconv"
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
	"github.com/stretchr/testify/assert"
)

func Test_IndexAll(t *testing.T) {
	out := slicelib.IndexAll([]int{3, 6, 2, 8, 3, 3}, 3)
	assert.Equal(t, []int{0, 4, 5}, out)
}

func Test_Map(t *testing.T) {
	out := slicelib.Map([]int{2, 4, 5}, func(i int) string { return strconv.Itoa(i) })
	assert.Equal(t, []string{"2", "4", "5"}, out)
}
