package main

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func TestSubBigInt(t *testing.T) {
	testP, _ := rand.Prime(rand.Reader, 1024)
	//hold := new(big.Int).Mul(testP, big.NewInt(1)) // fuckass copy
	hold := new(big.Int).Set(testP)
	testP.Sub(testP, big.NewInt(1))
	if testP.Cmp(hold) != -1 {
		t.Error("testP >= testPHold")
	}
}

func TestFunc(t *testing.T) {

	res := 1
	exp := 1
	if res != exp {
		t.Error("Expected", exp, "but was", res)
	}
}
