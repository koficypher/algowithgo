package module01

import (
	"fmt"
)

//DecToAny converts any decimal number to any base by taking in
//the number to convert and the base
func DecToAny(dec, base int) string {
	var result string
	for dec > 0 {
		rem := dec % base
		dec = dec / base
		result = fmt.Sprintf("%s%s", checkHexRep(rem), result)
	}

	fmt.Println(result)
	return ""
}

func checkHexRep(digit int) string {
	var rep string
	switch digit {
	case 10:
		rep = "A"
	//	fmt.Println(rep)
	case 11:
		rep = "B"
	//	fmt.Println(rep)
	case 12:
		rep = "C"
	//	fmt.Println(rep)
	case 13:
		rep = "D"
	//	fmt.Println(rep)
	case 14:
		rep = "E"
	//	fmt.Println(rep)
	case 15:
		rep = "F"
		//	fmt.Println(rep)
	}

	return rep
}
