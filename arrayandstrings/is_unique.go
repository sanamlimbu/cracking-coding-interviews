package arrayandstrings

import (
	"sort"
)

/*
 	Implement an algorithm to determine if a string has all unique characters.
	What if you cannot use additional data structures?
*/

// Time complexity:  0(n log n)
func IsUnique(str string) bool {
	var arr []int
	for _, c := range str {
		arr = append(arr, int(c))
	}

	// Sort the array
	sort.Ints(arr)

	for i := 0; (i + 1) < len(arr); i++ {
		if arr[i] == arr[i+1] {
			return false
		}
	}

	return true
}

// Time complexity: 0(n)
func IsUnique2(str string) bool {
	if len(str) > 128 {
		return false
	}
	// Assuming string has ASCII characters i.e 8-bit ASCII encoding, 128 characters
	charSet := [128]bool{}

	for i := 0; i < len(str); i++ {
		val := str[i]
		if charSet[val] {
			return false
		}
		charSet[val] = true
	}

	return true
}
