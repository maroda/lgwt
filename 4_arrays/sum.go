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
