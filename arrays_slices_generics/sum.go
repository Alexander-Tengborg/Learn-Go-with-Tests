package arraysslices

func Reduce[A, B any](arr []A, fn func(B, A) B, initialValue B) B {
	result := initialValue
	for _, val := range arr {
		result = fn(result, val)
	}

	return result
}

func Sum(numbers []int) int {
	sum := func(result int, val int) int {
		return result + val
	}

	return Reduce(numbers, sum, 0)
}

func SumAllTails(numbers ...[]int) []int {
	sumTail := func(result []int, value []int) []int {
		if len(value) > 1 {
			return append(result, Sum(value[1:]))
		}
		return append(result, 0)
	}

	return Reduce(numbers, sumTail, []int{})
}

func Find[T any](numbers []T, fn func(T) bool) (T, bool) {
	for _, num := range numbers {
		if fn(num) {
			return num, true
		}
	}

	var zero T
	return zero, false
}
