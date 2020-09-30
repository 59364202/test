package array

import ()

func RemoveStringArrayByValue(array []string, x ...string) []string {
	for _, value := range x {
		i := FindStringInArray(array, value)
		if i >= 0 {
			array = append(array[:i], array[i+1:]...)
		}
	}

	return array
}
func FindStringInArray(array []string, x string) int {
	for i, v := range array {
		if v == x {
			return i
		}
	}
	return -1
}

func RemoveIntArrayByIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
