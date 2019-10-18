package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {

	A, _ := new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000000", 16)
	B, _ := new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000007", 16)
	P, _ := new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
	N, _ := new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
	H := big.NewInt(1)

	Gx, _ := new(big.Int).SetString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 16)
	Gy, _ := new(big.Int).SetString("483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 16)

	bitcoinCurve := Curve{
		A: A,
		B: B,
		P: P,
		N: N,
		H: H,
		G: &Point{
			X: Gx,
			Y: Gy,
		},
	}

	x := new(big.Int)
	x.SetString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 16)

	y := bitcoinCurve.GetY(x)
	log.Println(y)
	log.Println(toHexInt(y))

}

func toHexInt(n *big.Int) string {
	return fmt.Sprintf("%x", n) // or %X or upper case
}
