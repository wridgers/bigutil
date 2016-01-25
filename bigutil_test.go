package bigutil_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/wridgers/bigutil"
)

func TestSqrt(t *testing.T) {

	var cases = []struct {
		number string
		square string
	}{
		{"2", "4"},
		{"1.4142131805419921875", "2"},
	}

	for _, c := range cases {
		n, success := new(big.Float).SetPrec(20).SetString(c.number)

		if !success {
			t.Errorf("Failed to create big.Float from string: %s\n", c.number)
		}

		s, success := new(big.Float).SetPrec(20).SetString(c.square)

		if !success {
			t.Errorf("Failed to create big.Float from string: %s\n", c.square)
		}

		sqrt := bigutil.Sqrt(s)

		fmt.Printf("sqrt(%.50f) == %.50f\n", s, sqrt)

		if sqrt.Cmp(n) != 0 {
			t.Error("sometthing is wrong")
		}
	}
}
