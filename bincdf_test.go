package fixed56

import (
	"fmt"
	"testing"
)

func TestFixed_BinCDFa(t *testing.T) {
	//  0.005131128987239601876
    //  https://www.wolframalpha.com/input/?i=N%5BCDF%5BBinomialDistribution%5B34%2C0.9%5D%2C25%5D%2C30%5D
	fmt.Printf("%.17f ?= 0.00513112898723960\n", BinCDF(34, From(0.9),25).Float())
}

func TestFixed_BinCDFb(t *testing.T) {
	//  0.000069568970065445156
	//  https://www.wolframalpha.com/input/?i=N%5BCDF%5BBinomialDistribution%5B340%2C0.82%5D%2C250%5D%2C20%5D
	fmt.Printf("%.17f ?= 0.00006956897006545\n", BinCDF(340, From(0.82), 250).Float())
}

func TestFixed_BinCDFc(t *testing.T) {
	//  0.97799731700834905720
	//  https://www.wolframalpha.com/input/?i=N%5BCDF%5BBinomialDistribution%5B3400%2C0.72%5D%2C2500%5D%2C20%5D
	fmt.Printf("%.17f ?= 0.97799731700834905\n", BinCDF(3400, From(0.72), 2500).Float())
}

func TestFixed_BinCDFd(t *testing.T) {
	//  0.50230327513576946695
	//  https://www.wolframalpha.com/input/?i=N%5BCDF%5BBinomialDistribution%5B30000%2C0.5%5D%2C15000%5D%2C20%5D
	fmt.Printf("%.17f ?= 0.50230327513576946\n", BinCDF(30000, From(0.5), 15000).Float())
}

func TestFixed_BinCDF2(t *testing.T) {
	acc := accuracy{Epsilon: 1e-10}
	for i, tc := range bincdfTestCases {
		p := From(tc.p)
		got := BinCDF(tc.n,p,tc.x)
		if ok := acc.update(got, tc.cdf); !ok {
			t.Errorf("%d: BinCDF(%v,%v,%v) => got %v|%v, want %v|%v",  i, tc.n, tc.p, tc.x, got, got.Float(), From(tc.cdf), tc.cdf)
		}
	}
	t.Log(acc)
}

func BenchmarkFixed_BinCDF(b *testing.B) {
	for i:=0; i< b.N; i++ {
		tc := bincdfTestCases[i%len(bincdfTestCases)]
		BinCDF(tc.n,From(tc.p),tc.x)
	}
}

