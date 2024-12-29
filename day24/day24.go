package day24

import (
	"cmp"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
)

type Gate string

const (
	and Gate = "AND"
	or  Gate = "OR"
	xor Gate = "XOR"
)

type Wire struct {
	gate   Gate
	in     []string
	output int
}

func (w *Wire) calcOutput(m map[string]*Wire) {
	if w.output != -1 {
		return
	}

	w0 := m[w.in[0]]
	w0.calcOutput(m)

	w1 := m[w.in[1]]
	w1.calcOutput(m)

	switch w.gate {
	case and:
		w.output = w0.output & w1.output
	case or:
		w.output = w0.output | w1.output
	case xor:
		w.output = w0.output ^ w1.output
	default:
		panic(w.gate)
	}
}

func Part1(input string) int {
	return simulate(parseInput(input))
}

func Part2(input string) string {
	wires, x, y := parseInput(input)

	// graph(wires)

	swaps := []string{"fhc", "z06", "qhj", "z11", "hqk", "z35", "ggt", "mwh"}
	for i := 0; i < len(swaps); i += 2 {
		w0 := swaps[i]
		w1 := swaps[i+1]
		wires[w0].gate, wires[w1].gate = wires[w1].gate, wires[w0].gate
		wires[w0].in, wires[w1].in = wires[w1].in, wires[w0].in
	}

	// graph(wires)

	for range 10000 {
		z := simulate(wires, x, y)

		d := x + y - z
		if d != 0 {
			panic(fmt.Sprintf("Invalid answer: %d + %d != %d (d=%d)", x, y, z, d))
		}

		x = rand.Intn(1 << 45)
		y = rand.Intn(1 << 45)
	}

	slices.Sort(swaps)
	return strings.Join(swaps, ",")
}

func parseInput(input string) (map[string]*Wire, int, int) {
	parts := strings.Split(input, "\n\n")

	wires := map[string]*Wire{}
	xy := make(map[byte]int, 2)

	for _, l := range common.GetLines(parts[0]) {
		f := strings.FieldsFunc(l, func(r rune) bool { return r == ':' || r == ' ' })
		name := f[0]
		wires[name] = &Wire{output: -1}

		v, _ := strconv.Atoi(f[1])
		bit, _ := strconv.Atoi(name[1:])
		xy[name[0]] += v << bit
	}

	for _, l := range common.GetLines(parts[1]) {
		f := strings.Fields(l)
		name := f[4]
		w := &Wire{gate: Gate(f[1]), in: []string{f[0], f[2]}, output: -1}
		slices.Sort(w.in)
		wires[name] = w
	}

	return wires, xy['x'], xy['y']
}

func simulate(wires map[string]*Wire, x, y int) int {
	for name, w := range wires {
		if name[0] == 'x' {
			bit, _ := strconv.Atoi(name[1:])
			w.output = (x >> bit) & 1
			continue
		}

		if name[0] == 'y' {
			bit, _ := strconv.Atoi(name[1:])
			w.output = (y >> bit) & 1
			continue
		}

		w.output = -1
	}

	result := 0
	for name, w := range wires {
		if name[0] != 'z' {
			continue
		}

		w.calcOutput(wires)
		bit, _ := strconv.Atoi(name[1:])
		result += w.output << bit
	}

	return result
}

func graph(wires map[string]*Wire) {
	colors := map[Gate]string{
		and: "red",
		or:  "blue",
		xor: "green",
	}

	var sb strings.Builder
	sb.WriteString("digraph {\n")

	var xy, z []string
	for name, w := range wires {
		if name[0] == 'x' || name[0] == 'y' {
			xy = append(xy, name)
		}
		if name[0] == 'z' {
			z = append(z, name)
		}

		if len(w.in) == 0 {
			continue
		}

		sb.WriteString(fmt.Sprintf("  { %s } -> %s [color=%s]\n", strings.Join(w.in, " "), name, colors[w.gate]))
	}

	sb.WriteString(fmt.Sprintf("  subgraph xy { rank=same; %s }\n", strings.Join(xy, " ")))

	slices.SortFunc(z, func(a, b string) int { return -cmp.Compare(a, b) })
	sb.WriteString(fmt.Sprintf("  subgraph z { rank=same; ordering=out; %s }\n", strings.Join(z, " -> ")))

	sb.WriteString("}")

	os.WriteFile("graph.gv", []byte(sb.String()), 0600)
}
