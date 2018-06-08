package cribaeratostenes

import (
	"testing"
	"log"
)

func TestCriba(t *testing.T) {
	n := 10

	expected := "2 3 5 7 "

	result := Criba(n)
	if expected != result {
		log.Fatalf("Expected %s but got %s ", expected, result)
	}

}
