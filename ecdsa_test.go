package ecdsa

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var secp256k1 Curve

func init() {
	secp256k1.A, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000000", 16)
	secp256k1.B, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000007", 16)

	secp256k1.G.X, _ = new(big.Int).SetString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 16)
	secp256k1.G.Y, _ = new(big.Int).SetString("483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 16)

	secp256k1.P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
	secp256k1.N, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
	secp256k1.H, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000001", 16)
}

func TestIsInfinity(t *testing.T) {
	p := Point{
		X: big.NewInt(0),
		Y: big.NewInt(0),
	}
	assert.Equal(t, true, p.IsInfinity())

	p = Point{
		X: nil,
		Y: nil,
	}
	assert.Equal(t, true, p.IsInfinity())

	p = Point{
		X: big.NewInt(-1),
		Y: big.NewInt(0),
	}
	assert.Equal(t, false, p.IsInfinity())

	p = Point{
		X: big.NewInt(0),
		Y: big.NewInt(-1),
	}
	assert.Equal(t, false, p.IsInfinity())

	p = Point{
		X: big.NewInt(999),
		Y: big.NewInt(-999),
	}
	assert.Equal(t, false, p.IsInfinity())
}

func TestIsOnCurve(t *testing.T) {
	/* Generator Point */
	assert.Equal(t, true, secp256k1.IsOnCurve(secp256k1.G))

	/* Random public key */
	var p Point
	p.X, _ = new(big.Int).SetString("440D851DBCEF5A43A5415B7250C8BEF01D14E1AF93DA31C4F3FFAA2CC3D22B4A", 16)
	p.Y, _ = new(big.Int).SetString("273A11F6571E968A62B63D0F7D2EFD1A52CDA4E7B65EEA5403E43A87A9DEE018", 16)

	assert.Equal(t, true, secp256k1.IsOnCurve(p))

	/* Point at infinity */
	p.X = big.NewInt(0)
	p.Y = big.NewInt(0)
	assert.Equal(t, false, secp256k1.IsOnCurve(p))
	p.X = nil
	p.Y = nil
	assert.Equal(t, false, secp256k1.IsOnCurve(p))
}
