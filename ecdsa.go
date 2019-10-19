package main

import (
	"fmt"
	"log"
	"math/big"
)

// Curve struct
// y² = x^3 + ax + b
type Curve struct {
	A *big.Int
	B *big.Int
	P *big.Int
	N *big.Int
	H *big.Int

	G Point // Generator Point
}

// Point struct
type Point struct {
	X *big.Int
	Y *big.Int
}

// IsInfinity checks if a point is the point at infinity
func (p *Point) IsInfinity() bool {
	return p.X == nil && p.Y == nil
}

// ScalarMult computes the scalar multiplication of P by n defined by nP = P + P + ... + P where P + P is the point addition of P and P
func (c *Curve) ScalarMult(n *big.Int, P Point) *Point {
	/* Using Montgomery ladder algorithm */
	var R0 Point
	R1 := P

	binary := fmt.Sprintf("%b", n)

	for i := len(binary) - 1; i >= 0; i-- {
		if string(binary[i]) == "1" {
			R1 = c.AddPoints(R0, R1)
			R0 = c.AddPoints(R0, R0)
		} else {
			R0 = c.AddPoints(R0, R1)
			R1 = c.AddPoints(R1, R1)
		}
	}

	return &R0
}

// AddPoints computes the addition of two points on the curve
func (c *Curve) AddPoints(P Point, Q Point) Point {
	R := new(Point)

	/* Checkout https://en.wikipedia.org/wiki/Elliptic_curve_point_multiplication#Point_operations */
	if P.IsInfinity() && Q.IsInfinity() {
		/* Point at infinity */
		// 0 + 0 = 0
		R.X = nil
		R.Y = nil
	} else if P.IsInfinity() {
		/* Point at infinity */
		// 0 + Q = Q
		R.X = Q.X
		R.Y = Q.Y
	} else if Q.IsInfinity() {
		/* Point at infinity */
		// P + 0 = P
		R.X = P.X
		R.Y = P.Y
	} else if P.X.Cmp(big.NewInt(0)) == 0 && c.AddMod(P.Y, Q.Y).Cmp(big.NewInt(0)) == 0 {
		/* Point negation */
		// Check if P.X = 0 and P.Y + Q.Y = 0 ⇒ P.Y = -Q.Y
		R.X = nil
		R.Y = nil
	} else if P.X.Cmp(Q.X) != 0 {
		/* Point addition */
		// Check if P.X != P.Y
		R = c.PointAddition(P, Q)
	} else if P.X.Cmp(Q.X) == 0 && P.Y.Cmp(Q.Y) == 0 && P.Y.Cmp(big.NewInt(0)) != 0 {
		/* Point doubling */
		R = c.PointDoubling(P, Q)
	} else {
		log.Fatal("Not supported")
	}

	return *R
}

// PointAddition computes the point addition of P + Q
func (c *Curve) PointAddition(P, Q Point) *Point {
	R := new(Point)

	/* Point addition */
	// Check if P.X != P.Y

	lambda := c.SubMod(Q.Y, P.Y) // Q.Y - P.Y
	q := c.SubMod(Q.X, P.X)      // Q.X - P.X
	q = c.InvMod(q)

	lambda = c.MultMod(lambda, q) // λ = (Q.Y - P.Y) / (Q.X - P.X)

	/* Compute x */
	// x = λ² - P.X - Q.X
	R.X = c.MultMod(lambda, lambda)
	R.X = c.SubMod(R.X, P.X)
	R.X = c.SubMod(R.X, Q.X)

	/* Compute y */
	// y = λ(P.X - R.X) - P.Y
	R.Y = c.SubMod(P.X, R.X)
	R.Y = c.MultMod(lambda, R.Y)
	R.Y = c.SubMod(R.Y, P.Y)

	return R
}

// PointDoubling computes the point doubling of P + Q
func (c *Curve) PointDoubling(P, Q Point) *Point {
	/* Point Doubling */
	/* The operation is the same as the Point Addition, except lambda is different */
	// Check if P.X == P.Y and P.Y = Q.Y and P.Y != 0
	// ⇒ P and Q are coincident

	R := new(Point)

	/* Compute λ */
	// λ = (3*(P.X)² + a) / (2*P.Y)
	// a = curve.A
	lambda := c.MultMod(P.X, P.X)
	lambda = c.MultMod(lambda, big.NewInt(3))
	lambda = c.AddMod(lambda, c.A)

	lambda = c.MultMod(lambda, c.InvMod(c.MultMod(P.Y, big.NewInt(2))))

	/* Point addition from here */

	/* Compute x */
	// x = λ² - P.X - Q.X
	R.X = c.MultMod(lambda, lambda)
	R.X = c.SubMod(R.X, P.X)
	R.X = c.SubMod(R.X, Q.X)

	/* Compute y */
	// y = λ(P.X - R.X) - P.Y
	R.Y = c.SubMod(P.X, R.X)
	R.Y = c.MultMod(lambda, R.Y)
	R.Y = c.SubMod(R.Y, P.Y)

	return R
}

// GetY returns y given a x
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

// AddMod computes (x + y) mod curve.P
func (c *Curve) AddMod(x, y *big.Int) *big.Int {
	z := new(big.Int)
	z.Add(x, y)
	z.Mod(z, c.P)
	return z
}

// SubMod computes (x - y) mod curve.P
func (c *Curve) SubMod(x, y *big.Int) *big.Int {
	z := new(big.Int)
	z.Sub(x, y)
	z.Mod(z, c.P)
	return z
}

// MultMod computes (x * y) mod curve.P
func (c *Curve) MultMod(x, y *big.Int) *big.Int {
	z := new(big.Int)
	z.Mul(x, y)
	z.Mod(z, c.P)
	return z
}

// DivMod computes (x / y) mod curve.P
func (c *Curve) DivMod(x, y *big.Int) *big.Int {
	z := new(big.Int)
	z.Div(x, y)
	z.Mod(z, c.P)
	return z
}

// InvMod computes (1/X) mod curve.P
func (c *Curve) InvMod(x *big.Int) *big.Int {
	z := new(big.Int)
	z.ModInverse(x, c.P)
	return z
}

// IsOnCurve checks if P is on elliptic curve
func (c *Curve) IsOnCurve(P Point) bool {
	if P.IsInfinity() {
		return false
	}

	y2 := c.MultMod(P.Y, P.Y) // y²

	eq := c.MultMod(P.X, P.X)              // x^2
	eq = c.MultMod(eq, P.X)                // x^3
	eq = c.AddMod(eq, c.MultMod(c.A, P.X)) // x^3 + ax
	eq = c.AddMod(eq, c.B)                 // x^3 + ax + b

	// y² = x^3 + ax + b
	return y2.Cmp(eq) == 0
}
