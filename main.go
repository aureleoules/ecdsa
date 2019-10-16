package main

import (
	"fmt"
	"log"
	"math/big"
)

var a = int64(0)
var b = int64(7)
var n = new(big.Int)

func init() {
	n.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)

}

func main() {
	log.Println(n)
	x := new(big.Int)
	x.SetString("47B6822336E45D5C12A5631ACA8E09ECEA7E977E260A6EFC02D24F462326EF2B", 16)
	log.Println("x=", x)
	// x.SetInt64(259875)
	y := F(x)

	yHexa := toHexInt(y)
	log.Println(yHexa)
}

// F (x) = y
func F(x *big.Int) *big.Int {
	y := new(big.Int)
	y.Add(y, x)
	y.Exp(y, big.NewInt(3), nil)

	log.Println("y = x^3")

	ax := big.NewInt(a)
	ax.Mul(ax, x)
	y.Add(y, ax)
	log.Println("y = x^3 + ax")
	y.Add(y, big.NewInt(b))
	log.Println("y = x^3 + ax + b")

	n.Add(n, big.NewInt(1))
	n.Div(n, big.NewInt(4))
	p := new(big.Int)
	p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
	y.Exp(y, n, p)

	log.Println("y = sqrt(x^3 + ax + b)")
	return y
}

func toHexInt(n *big.Int) string {
	return fmt.Sprintf("%x", n) // or %X or upper case
}
