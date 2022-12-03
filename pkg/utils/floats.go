package utils

func Abs[T float32 | float64](a T) T {
	if a < 0 {
		a = a * -1
	}
	return a
}

func AlmostEqual[T float32 | float64](a, b T) bool {
	const e = 1e-6

	diff := Abs(a - b)

	return diff < e
}
