package random_test

import (
	"testing"
	"../random"
)

func TestLimitGreatThanSizeRepetitionsTrue(t *testing.T) {
	array, err := random.Random(10, 5, true)
	if err != nil && len(array) != 10 {
		t.Error("Size can be great than limit when repetitions false")
	}
}

func TestLimitGreatThanSizeRepetitionsFalse(t *testing.T) {
	_, err := random.Random(10, 5, false)
	if err == nil  {
		t.Error("Size must be smaller than limit when repetitions false")
	}
}

func TestReturnSizeEqualsParameterWithRepetitionsFalse(t *testing.T) {
	array, err := random.Random(5, 5, false)
	if err != nil && len(array) != 10 {
		t.Error("Error or length unexpected")
	}
}

func TestReturnsSizeEqualsParameterWithRepetitionsTrue(t *testing.T) {
	array, err := random.Random(5, 5, true)
	if err != nil && len(array) != 10 {
		t.Error("Error or length unexpected")
	}
}

func TestSizeWithInvalidZeroWithRepetitionsFalse(t *testing.T) {
	_, err := random.Random(0, 5, false)
	if err == nil  {
		t.Error("Can't be zero")
	}
}

func TestSizeWithInvalidZeroWithRepetitionsTrue(t *testing.T) {
	_, err := random.Random(0, 5, true)
	if err == nil  {
		t.Error("Can't be zero")
	}
}

func TestLimitWithInvalidZeroWithRepetitionsFalse(t *testing.T) {
	_, err := random.Random(5, 0, false)
	if err == nil  {
		t.Error("Can't be zero")
	}
}

func TestLimitWithInvalidZeroWithRepetitionsTrue(t *testing.T) {
	_, err := random.Random(5, 0, true)
	if err == nil  {
		t.Error("Can't be zero")
	}
}
func TestSizeWithInvalidNegativeWithRepetitionsFalse(t *testing.T) {
	_, err := random.Random(-1, 5, false)
	if err == nil  {
		t.Error("Can't be negative")
	}
}

func TestSizeWithInvalidNegativeWithRepetitionsTrue(t *testing.T) {
	_, err := random.Random(-2, 5, true)
	if err == nil  {
		t.Error("Can't be negative")
	}
}

func TestLimitWithInvalidNegativeWithRepetitionsFalse(t *testing.T) {
	_, err := random.Random(5, -3, false)
	if err == nil  {
		t.Error("Can't be negative")
	}
}

func TestLimitWithInvalidNegativeWithRepetitionsTrue(t *testing.T) {
	_, err := random.Random(5, -4, true)
	if err == nil  {
		t.Error("Can't be negative")
	}
}