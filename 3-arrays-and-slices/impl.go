package arrays_and_slices

func Sum(arr []int) int {
	sum := 0

	for _, number := range arr {
		sum += number
	}

	return sum
}

func SumAll(arrays ...[]int) []int {
	var result []int

	for _, arr := range arrays {
		result = append(result, Sum(arr))
	}

	return result
}

func SumAllTails(arrays ...[]int) []int {
	var result []int

	for _, arr := range arrays {
		var tail int

		if len(arr) == 0 {
			tail = 0
		} else {
			tail = Sum(arr[1:])
		}

		result = append(result, tail)
	}

	return result
}
