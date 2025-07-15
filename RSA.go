package main

import (
	"crypto/rand"
	"math/big"
)

type RSAClient struct {
	product, encryptExp, decryptExp *big.Int
}

/*
 */
func (c *RSAClient) New() {
	c.product = new(big.Int)
	c.encryptExp = new(big.Int)
	c.decryptExp = new(big.Int)
}

/*
 */
func (c *RSAClient) generateKeys() {
	primeA, _ := rand.Prime(rand.Reader, 1024)
	primeB, _ := rand.Prime(rand.Reader, 1024)
	c.product.Mul(primeA, primeB)

	// phi is the result of Eulers Totient Function, since primeA and primeB are both primes,
	// we can calculate the value as A-1 * B-1
	phi := new(big.Int).Mul(primeA.Sub(primeA, big.NewInt(1)), primeB.Sub(primeB, big.NewInt(1)))

	// industry standard encryption exponent. 4th Fermat Number
	c.encryptExp.Set(big.NewInt(65537))

	// decryption exponent calculated such that e*d ≡ 1 (mod phi)
	c.decryptExp.ModInverse(c.encryptExp, phi)
}

/*
 */
func (c *RSAClient) encrypt(message *big.Int) *big.Int {
	return new(big.Int).Exp(message, c.encryptExp, c.product)
}

/*
 */
func (c *RSAClient) decrypt(encryptedMessage *big.Int) *big.Int {
	return new(big.Int).Exp(encryptedMessage, c.decryptExp, c.product)
}
