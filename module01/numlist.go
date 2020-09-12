package module01

//NumList returns true when an integer is in the list
//and fals when otherwise
func NumList(list []int, num int) bool {
	for _, i := range list {
		if i == num {
			return true
		}
	}
	return false
}
