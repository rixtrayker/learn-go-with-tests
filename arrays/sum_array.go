
package arrays

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}

	return
}
