
package arrays

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}

	return
}
