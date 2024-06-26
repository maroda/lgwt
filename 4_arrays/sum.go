package main

func Sum(numbers []int) int {
	sum := 0
	//	for i := 0; i < 5; i++ {
	//		sum += numbers[i]
	//	}
	for _, number := range numbers { // range returns an index and the value, `_` means ignore the index
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int { // variadic function can take arbitrary number of input args
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	//	lengthOfNumbers := len(numbersToSum)
	//	sums := make([]int, lengthOfNumbers)
	//
	//	for i, numbers := range numbersToSum {
	//		sums[i] = Sum(numbers)
	//	}

	return sums
}
