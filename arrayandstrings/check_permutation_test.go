package arrayandstrings

import (
	"testing"
)

func TestCheckPermutation(t *testing.T) {
	if !CheckPermutation("dog", "god") {
		t.Error("'dog' and 'god' are permutation of eachother but failed")
	}

	if CheckPermutation("hello", "llgeo") {
		t.Error("'hello' and 'llgeo' are not permutation of eachother but passed")
	}

	if CheckPermutation("newyork", "dallas") {
		t.Error("'newyork' and 'dallas' are not permutation of eachother but passed")
	}
}

func TestCheckPermutation2(t *testing.T) {
	if !CheckPermutation2("dog", "god") {
		t.Error("'dog' and 'god' are permutation of eachother but failed")
	}

	if CheckPermutation2("hello", "llgeo") {
		t.Error("'hello' and 'llgeo' are not permutation of eachother but passed")
	}

	if CheckPermutation2("newyork", "dallas") {
		t.Error("'newyork' and 'dallas' are not permutation of eachother but passed")
	}
}
