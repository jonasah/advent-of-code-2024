package slicelib

import (
	"slices"
	"strconv"
)

func FilterFunc[T any](a []T, f func(T) bool) []T {
	var out []T
	for _, v := range a {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

func FindFunc[T any](a []T, f func(T) bool) (T, bool) {
	i := slices.IndexFunc(a, f)
	if i == -1 {
		var empty T
		return empty, false
	}

	return a[i], true
}

func IndexAll[T comparable](a []T, val T) []int {
	var matches []int
	for i, v := range a {
		if v == val {
			matches = append(matches, i)
		}
	}

	return matches
}

func IndexAllFunc[T comparable](a []T, f func(T) bool) []int {
	var matches []int
	for i, v := range a {
		if f(v) {
			matches = append(matches, i)
		}
	}

	return matches
}

func Map[T, U any](a []T, f func(T) U) []U {
	out := make([]U, 0, len(a))
	for _, v := range a {
		out = append(out, f(v))
	}
	return out
}

func Atoi(a []string) []int {
	l := make([]int, 0, len(a))
	for _, s := range a {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		l = append(l, i)
	}

	return l
}

func Itoa(a []int) []string {
	return Map(a, strconv.Itoa)
}
