package main

import (
	"math/big"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {

	c := new(RSAClient)
	c.New()
	c.generateKeys()
	origMessage := big.NewInt(9283094820398)
	m := c.encrypt(origMessage)
	d := c.decrypt(m)

	if origMessage.Cmp(d) != 0 {
		t.Error("Expected", origMessage, "but was", d)
	}
}
