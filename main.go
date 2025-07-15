package main

import (
	"crypto/rand"
	"math/big"
)

func generateKeys(product, encryptExp, decryptExp *big.Int) {
	primeA, _ := rand.Prime(rand.Reader, 1024)
	primeB, _ := rand.Prime(rand.Reader, 1024)
	product.Mul(primeA, primeB)

	// phi is the result of Eulers Totient Function, since primeA and primeB are both primes,
	// we can calculate the value as A-1 * B-1
	phi := new(big.Int).Mul(primeA.Sub(primeA, big.NewInt(1)), primeB.Sub(primeB, big.NewInt(1)))

	// industry standard encryption exponent. 4th Fermat Number
	encryptExp.Set(big.NewInt(65537))

	// decryption exponent calculated such that e*d ≡ 1 (mod phi)
	decryptExp.ModInverse(encryptExp, phi)
}

func main() {
	var product, encryptExp, decryptExp big.Int
	generateKeys(&product, &encryptExp, &decryptExp)

}
