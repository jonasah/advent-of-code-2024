package slicelib

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
