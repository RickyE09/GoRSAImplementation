package main

import (
	"fmt"
	"math/big"
)

func main() {
	c := new(RSAClient)
	c.New()
	c.generateKeys()
	fmt.Println("Original message example: 9283094820398")
	m := c.encrypt(big.NewInt(9283094820398))
	fmt.Println("encrypted:", m.String())
	d := c.decrypt(m)
	fmt.Println("decrypted:", d)
}
