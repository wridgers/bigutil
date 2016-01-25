package bigutil

import "math/big"

// Sqrt returns the square root n.
func Sqrt(n *big.Float) *big.Float {
	prec := n.Prec()

	x := new(big.Float).SetPrec(prec).SetInt64(1)
	z := new(big.Float).SetPrec(prec).SetInt64(1)

	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)
	t := new(big.Float).SetPrec(prec)

	for {
		z.Copy(x)

		t.Mul(x, x)
		t.Sub(t, n)
		t.Quo(t, x)
		t.Mul(t, half)
		x.Sub(x, t)

		if x.Cmp(z) == 0 {
			break
		}
	}

	return x
}

// Log returns the natural logarithm of x.
func Log(x *big.Float) *big.Float {

	// Look here: http://stackoverflow.com/questions/739532/logarithm-of-a-bigdecimal
	return x
}
