package module01

import (
	"fmt"
)

//ReverseString reverses a string
func ReverseString(word string) string {
	var res string
	for i := len(word) - 1; i >= 0; i-- {
		//	fmt.Println(string(word[i]))
		res = res + string(word[i])
	}
	//fmt.Println(res)
	return res
}

//ReverseStringAlt is an alternative for reversing strings
func ReverseStringAlt(word string) string {
	var res string
	for _, i := range word {
		res = string(i) + res
	}
	fmt.Println(res)
	return res
}
