package pointer

import "testing"

func TestDouble(t *testing.T) {
	given := 2
	want := 4

	result := given

	double(&result)

	if result != want {
		t.Errorf("double(&%d), %d -> %d, want %d", given, given, result, want)
	}
}

func TestAppendGreeting(t *testing.T) {
	given := "Bob"
	want := "Hi, Bob"

	result := given

	appendGreeting(&result)

	if result != want {
		t.Errorf("appendGreeting(&%s), %s -> %s, want %s", given, given, result, want)
	}
}
