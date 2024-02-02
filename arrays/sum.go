package main

func Sum(numbers []int) int {
	sum := 0

	// for i := 0; i < len(numbers); i++ {
	// 	sum += numbers[i]
	// }

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var res []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			res = append(res, Sum(numbers))
		} else {
			tail := Sum(numbers[1:])
			res = append(res, tail)
		}
	}

	return res

}
