package module01

import (
	"fmt"
)

// AnyToDec converts a binary digit to a base by
func AnyToDec(value string, base int) int {
	//	var result int
	var charset = "0123456789ABCDEF"
	res := 0
	multiplier := 1
	for i := len(value) - 1; i >= 0; i-- {
		index := -1
		for j, l := range charset {
			if l == rune(value[i]) {
				index = j
				break
			}
		}

		if index < 0 {
			panic("something went wrong")
		}
		res = res + index*multiplier
		multiplier = multiplier * base
	}
	fmt.Println(res)
	return res
}
