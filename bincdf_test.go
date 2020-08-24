package fixed56

import (
	"fmt"
	"testing"
)

func TestFixed_BinCDF(t *testing.T) {
	//  0.005131128987239594
	fmt.Printf("%.18f\n", BinCDF(34, From(0.9),25).Float())
}

func TestFixed_BinCDF1(t *testing.T) {
	//  0.000069568970065445
	fmt.Printf("%.18f\n", BinCDF(340, From(0.82), 250).Float())
}

func TestFixed_BinCDF2(t *testing.T) {
	acc := accuracy{Epsilon: 1e-8}
	for i, tc := range bincdfTestCases[125:] {
		p := From(tc.p)
		got := BinCDF(tc.n,p,tc.x)
		if ok := acc.update(got, tc.cdf); !ok {
			t.Errorf("%d: BinCDF(%v,%v,%v) => got %v|%v, want %v|%v",  i, tc.n, tc.p, tc.x, got, got.Float(), From(tc.cdf), tc.cdf)
		}
	}
	t.Log(acc)
}
