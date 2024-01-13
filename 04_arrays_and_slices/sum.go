package sum

func Sum(numbers []int) int {
	add := func(res, x int) int { return res + x }
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	fn := func(res, x []int) []int {
		res = append(res, Sum(x))
		return res
	}
	return Reduce(numbersToSum, fn, []int{})
}

func SumAllTails(numbersToSum ...[]int) []int {
	fn := func(res, x []int) []int {
		if len(x) == 0 {
			res = append(res, 0)
		} else {
			tail := x[1:]
			res = append(res, Sum(tail))
		}
		return res
	}
	return Reduce(numbersToSum, fn, []int{})
}

func Reduce[T any](collection []T, accumulator func(T, T) T, initialValue T) T {
	var result = initialValue
	for _, n := range collection {
		result = accumulator(result, n)
	}
	return result
}
