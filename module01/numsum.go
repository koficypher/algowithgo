package module01

//NumSum returns the sum of all integers in the list provided
func NumSum(list []int) int {
	sum := 0
	for _, i := range list {
		sum += i
	}
	return sum
}

//NumSumAlt is an alternative for solving the summation problem
//but with recursion
func NumSumAlt(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return list[0] + NumSumAlt(list[1:])
}
