package utils

import (
	"testing"

)

func AssertEq[T comparable](t *testing.T, expected T, result T) {
	if expected != result {
		t.Fatalf("Expected: %v Got: %v", expected, result)
	}
}

func AssertNoEq[T comparable](t *testing.T, expected T, result T) {
	if expected == result {
		t.Fatalf("Expected: %v Got: %v", expected, result)
	}
}


func AssertNoErr(t *testing.T,err error) {
	if(err != nil ){
		t.Fatalf("Error is not nil");
	}
}
