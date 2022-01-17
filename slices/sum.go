package main

func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}
	return
}

func SumAll(slicesToSum ...[]int) (sums []int) {
	lengthofSlicesToSum := len(slicesToSum)
	sums = make([]int, lengthofSlicesToSum)

	for i, numbers := range slicesToSum {
		sums[i] = Sum(numbers)
	}
	return
}
