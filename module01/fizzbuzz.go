package module01

import (
	"fmt"
)

//FizzBuzz receives a number N as argument and prints
//a list of numbers from 1 to N replacing its mod of 3 and 5
//with fizz and buzz respectively
func FizzBuzz(number int) {
	for i := 1; i <= number; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}

	}
}

//FizzBuzzAlt - alternative solving the fizzbuzz problem
func FizzBuzzAlt(number int) {
	for i := 1; i <= number; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}

		// if i != number {
		// 	fmt.Println(", ")
		// }
	}

	fmt.Println()
}
