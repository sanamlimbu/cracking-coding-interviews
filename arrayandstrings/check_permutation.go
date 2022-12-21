package arrayandstrings

import (
	"sort"
)

/*
	Given two strings, write a method to decide if one is a permutation of other.
*/

// Solution 1 : Sort the strings
func CheckPermutation(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	// convert both strings to slices
	var arr1, arr2 []int

	for _, c := range s1 {
		arr1 = append(arr1, int(c))
	}

	for _, c := range s2 {
		arr2 = append(arr2, int(c))
	}

	// sort both slices
	sort.Ints(arr1)
	sort.Ints(arr2)

	// compare slices
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

// Soluion 2 : Check if two strings have identical charcter count

func CheckPermutation2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var counts [128]int // ASCII assumption

	for _, c := range s1 {
		counts[c]++
	}

	for _, c := range s2 {
		counts[c]--
		if counts[c] < 0 {
			return false
		}
	}

	return true
}
