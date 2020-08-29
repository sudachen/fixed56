package fixed56

import (
	"fmt"
	"math"
	"testing"
)

func TestFixed_BinCDFa(t *testing.T) {
	//  0.005131128987239601876
    //  https://www.wolframalpha.com/input/?i=N%5BCDF%5BBinomialDistribution%5B34%2C0.9%5D%2C25%5D%2C30%5D
	v := BinCDF(34, From(0.9),25)
	v2 := bincdf_( 34, 0.9, 25)
	fmt.Printf("cdf(34,0.9,25) => 0.00513112898723960 | fixed: %.17f, float64: %.17f\n", v.Float(), v2)
}

func TestFixed_BinCDFb(t *testing.T) {
	//  0.000069568970065445156
	//  https://www.wolframalpha.com/input/?i=N%5BCDF%5BBinomialDistribution%5B340%2C0.82%5D%2C250%5D%2C20%5D
	v := BinCDF(340, From(0.82), 250)
	v2 := bincdf_( 340, 0.82, 250)
	fmt.Printf("cdf(340,0.82,250) => 0.00006956897006545 | fixed: %.17f, float64: %.17f\n", v.Float(), v2)
}

func TestFixed_BinCDFc(t *testing.T) {
	//  0.97799731700834905720
	//  https://www.wolframalpha.com/input/?i=N%5BCDF%5BBinomialDistribution%5B3400%2C0.72%5D%2C2500%5D%2C20%5D
	v := BinCDF(3400, From(0.72), 2500)
	v2 := bincdf_( 3400, 0.72, 2500)
	fmt.Printf("cdf(3400,0.72,2500) => 0.97799731700834905 | fixed: %.17f, float64: %.17f\n", v.Float(), v2)
}

func TestFixed_BinCDFd(t *testing.T) {
	//  0.50230327513576946695
	//  https://www.wolframalpha.com/input/?i=N%5BCDF%5BBinomialDistribution%5B30000%2C0.5%5D%2C15000%5D%2C20%5D
	v := BinCDF(30000, From(0.5), 15000)
	v2 := bincdf_(30000, 0.5, 15000)
	fmt.Printf("cdf(30000, 0.5, 15000) => 0.50230327513576946 | fixed: %.17f, float64: %17f\n", v.Float(), v2)
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

var bincdfResultFix Fixed

func Benchmark_Fixed_BinCDF(b *testing.B) {
	for i:=0; i< b.N; i++ {
		tc := bincdfTestCases[i%len(bincdfTestCases)]
		bincdfResultFix = BinCDF(tc.n,From(tc.p),tc.x)
	}

	bincdfResultFix.lo++
}

var bincdfResultFlt float64

func Benchmark_Float_BinCDF(b *testing.B) {
	for i:=0; i< b.N; i++ {
		tc := bincdfTestCases[i%len(bincdfTestCases)]
		bincdfResultFlt = bincdf_(tc.n,tc.p,tc.x)
	}

	bincdfResultFlt++
}

func incomplete_(a, b, x float64) float64 {
	// Iₓ(a,b) = (xᵃ*(1-x)ᵇ)/(a*B(a,b)) * (1/(1+(d₁/(1+(d₂/(1+...))))))
	// (xᵃ*(1-x)ᵇ)/B(a,b) = exp(lgamma(a+b) - lgamma(a) - lgamma(b) + a*log(x) + b*log(1-x))
	// d_{2m+1} = -(a+m)(a+b+m)x/((a+2m)(a+2m+1))
	// d_{2m}   = m(b-m)x/((a+2m-1)(a+2m))
	bt := 0.

	if 0 < x && x < 1 {
		lgamma_ := func(x float64) float64{
			a,s := math.Lgamma(x)
			return a*float64(s)
		}
		bt = math.Exp(lgamma_(a+b)-lgamma_(a)-lgamma_(b)+a*math.Log(x)+b*math.Log(1-x))
	} else if x < 0 || x > 1 {
		panic(ErrOverflow)
	}

	bcfx := func() float64 {
		if math.Abs(bt) < 1e-17 { return 0 }
		h := bcf_(x, a, b)
		return bt*h/a
	}

	if x > (a+1)/(a+b+2) {
		// symmetry transform
		// 1 - bt/b*bcf(1-x,b,a)
		x,a,b = 1-x,b,a
		return 1-bcfx()
	}
	return bcfx()
}

func bcf_(x, a, b float64) float64 {
	const e = 1e-14
	const iters = 300

	nonzero := func(z float64) float64 {
		if math.Abs(z) < 1e-16 {
			return 1e-16
		}
		return z
	}

	c := 1.
	d := 1./(nonzero(1.-x*(a+b)/(a+1.)))
	h := d
	del := 0.

	for m := int64(1); m < iters; m++ {
		fm := float64(m)
		amm := a+fm+fm

		// d_{2m} = n = m(b-m)x/((a+2m-1)(a+2m))
		//n := div(mulx(fm,fixed(b-m),xx),mul(fixed(a+m+m-1),amm))
		n := fm*(b-fm)*x/((amm-1.)*amm)
		d = 1./(nonzero(1.+n*d))
		c = nonzero(1 + n/c)
		h = h*d*c

		// d_{2m+1} = n = -(a+m)(a+b+m)x/((a+2m)(a+2m+1))
		n = -(a+fm)*(a+b+fm)*x/(amm*(amm+1.))
		d = 1/(nonzero(1+n*d))
		c = nonzero(1 + n/c)

		del = d*c
		h = h*del

		if math.Abs(del-1) < e {
			return h
		}
	}
	panic(ErrOverflow)
	return h
}

func bincdf_(n int64, p float64, x int64) float64 {
	if x < 0 {
		return 0
	} else if x >= n {
		return 1
	} else {
		return incomplete_(float64(n-x), float64(x+1), 1-p)
	}
}
