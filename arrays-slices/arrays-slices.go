package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberSlices ...[]int) []int {
	var sums []int
	for _, numbers := range numberSlices {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
