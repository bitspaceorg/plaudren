package utils

import (
	"testing"
)

func AssertEq[T comparable](t *testing.T, expected, result T) {
	t.Helper()
	if expected != result {
		t.Fatalf("Expected: %v Got: %v", expected, result)
	}
}

func AssertNoEq[T comparable](t *testing.T, expected, result T) {
	t.Helper()
	if expected == result {
		t.Fatalf("Expected: %v Got: %v", expected, result)
	}
}

func AssertNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Error is not nil")
	}
}
