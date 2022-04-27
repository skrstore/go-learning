package math

import "testing"

func TestAdd(t *testing.T) {

	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	{1, 3, 4},
	{3, 4, 7},
	{22, 3, 25},
	{12, 3, 16},
}

func TestAddTable(t *testing.T) {

	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func TestSubtract(t *testing.T) {

	got := Subtract(4, 6)
	want := -2

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
