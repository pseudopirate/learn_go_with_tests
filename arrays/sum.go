package main

func Sum(numbers []int) (sum int) {
	for _, d := range numbers {
		sum += d
	}

	return sum
}

func SumAll(numbers ...[]int) (result []int) {
	for _, d := range numbers {
		result = append(result, Sum(d))
	}

	return result
}

func SumAllTails(numbers ...[]int) (result []int) {
	var sums []int
	for _, numbers := range numbers {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
