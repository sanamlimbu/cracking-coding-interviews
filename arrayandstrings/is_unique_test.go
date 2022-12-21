package arrayandstrings

import "testing"

func TestIsUnique(t *testing.T) {

	if IsUnique("44") {
		t.Error("wanted false go true")
	}

	if IsUnique("117") {
		t.Error("wanted false fot true")
	}

	if !IsUnique("132") {
		t.Error("wanted true got false")
	}
}

func TestIsUnique2(t *testing.T) {

	if IsUnique2("44") {
		t.Error("wanted false go true")
	}

	if IsUnique2("117") {
		t.Error("wanted false fot true")
	}

	if !IsUnique2("132") {
		t.Error("wanted true got false")
	}
}
