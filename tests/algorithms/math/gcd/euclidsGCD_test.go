package main

import (
	"fmt"
	"testing"

	"github.com/iDeepesh/goNotebook/algorithms/math/gcd"
)

func TestFindGCD(t *testing.T) {
	fmt.Println("Running TestFindGCD")
	n := gcd.FindGCD(10, 5)
	if n != 5 {
		t.Fail()
	}
}

func TestFindGCDOne(t *testing.T) {
	fmt.Println("Running TestFindOne")
	n := gcd.FindGCD(1, 1)
	if n != 1 {
		t.Fail()
	}
}
