package slicelib

func Map[T, U any](a []T, f func(T) U) []U {
	out := make([]U, 0, len(a))
	for _, v := range a {
		out = append(out, f(v))
	}
	return out
}
