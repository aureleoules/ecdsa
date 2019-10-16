package main

import (
	"math/big"
)

type Curve struct {
	A *big.Int
	B *big.Int
	P *big.Int
	N *big.Int
	H *big.Int

	G *Point
}

type Point struct {
	X *big.Int
	Y *big.Int
}

func (c *Curve) GetY(x *big.Int) *big.Int {
	// Construct y equation (y^2 = x^3 + ax + b)
	y := new(big.Int)            // Initialize y
	y.Add(y, x)                  // Add x
	y.Exp(y, big.NewInt(3), nil) // x^3

	ax := c.A     // Initialize ax = curve.A
	ax.Mul(ax, x) // Multiply by x
	y.Add(y, ax)  // Add ax to y equation
	y.Add(y, c.B) // Add b to y equation

	// Initialize pMod
	pMod := new(big.Int)
	pMod.Set(c.P)

	// Compute p mod 4
	pMod.Mod(pMod, big.NewInt(4))

	// Initialize residue
	// Checkout https://en.wikipedia.org/wiki/Quadratic_residue#Pairs_of_residues_and_nonresidues
	res := new(big.Int)
	res.Set(c.P)

	// If p ≡ 3 (mod 4) then
	if pMod.Cmp(big.NewInt(3)) == 0 {
		res.Add(res, big.NewInt(1))
	} else { // if p ≡ 1 (mod 4)
		res.Add(res, big.NewInt(-1))
	}

	res.Div(res, big.NewInt(4))
	y.Exp(y, res, c.P)

	return y

}
