package main

func Sum(numbers []int) (sum int) {
	return Reduce(numbers, func(acc int, curr int) int {
		return acc + curr
	}, 0)
}

func SumAllTails(numberSlices ...[]int) []int {
	var sums []int
	return Reduce(numberSlices, func(acc []int, curr []int) []int {
		if len(curr) == 0 {
			return append(acc, 0)
		} else {
			tail := curr[1:]
			return append(acc, Sum(tail))
		}
	}, sums)
}

func Reduce[T any, P any](collection []T, fn func(acc P, curr T) P, initial P) P {
	for _, item := range collection {
		initial = fn(initial, item)
	}
	return initial
}

type Transaction struct {
	From, To string
	Sum      float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	var balance float64
	for _, t := range transactions {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	return balance
}
