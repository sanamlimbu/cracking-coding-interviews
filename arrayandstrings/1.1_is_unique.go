package arrayandstrings

import (
	"sort"
)

/*
 	Implement an algorithm to determine if a string has all unique characters.
	What if you cannot use additional data structures?
*/

func IsUnique(str string) bool {
	// This array holds the ASCII code of character (rune)
	var arr []int
	for _, c := range str {
		arr = append(arr, int(c))
	}

	// Sort the array
	sort.Slice(arr, func(i, j int) bool {
		return i < j
	})

	for i := 0; (i + 1) < len(arr); i++ {
		if arr[i] == arr[i+1] {
			return false
		}
	}

	return true
}
