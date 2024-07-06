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

func Reduce[T, P any](collection []T, fn func(acc P, curr T) P, initial P) P {
	for _, item := range collection {
		initial = fn(initial, item)
	}
	return initial
}

type Amount float64

type Transaction struct {
	From, To string
	Sum      Amount
}

func BalanceFor(transactions []Transaction, name string) Amount {
	adjustBalance := func(b Amount, t Transaction) Amount {
		if t.From == name {
			return b - t.Sum
		}
		if t.To == name {
			return b + t.Sum
		}
		return b
	}
	return Reduce(transactions, adjustBalance, 0.0)
}
