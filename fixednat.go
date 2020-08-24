package fixed56

import (
	"fmt"
	"math"
	"math/big"
)

type Fixed struct {
	val *big.Int
}

var rawFixedOne = rawfixed(1)
var fixedFracMask = rawfixed(fracMask)
var fixedOne = rawfixed(oneValue)
var fixedHalfValue = rawfixed(halfValue)
var fixedRoundValue = rawfixed(halfValue)

func rawfixed(x int64) Fixed {
	return Fixed{big.NewInt(x) }
}

func (x Fixed) fixed56() int64 {
	return x.val.Int64()
}

func fixed(v int64) Fixed {
	x := big.NewInt(v)
	return Fixed{x.Lsh(x,fracBits) }
}

func sign(x Fixed) int {
	return x.val.Sign()
}

func from(f float64) Fixed {
	a := math.Float64bits(f)
	if a != 0 {
		v := int64(a & ((uint64(1) << 52) - 1))
		e := int((a>>52)&((1<<11)-1)) - 1023 - 52 + fracBits
		v |= int64(1) << 52
		x := big.NewInt(v)
		if e < 0 {
			x = x.Rsh(x, uint(-e))
		} else {
			x = x.Lsh(x, uint(e))
		}
		if int64(a) < 0 {
			return Fixed{x.Neg(x) }
		}
		return Fixed{x }
	}
	return Fixed{new(big.Int) }
}

func integer(x Fixed) int64 {
	s := x.val.Sign()
	q := new(big.Int).Abs(x.val)
	return q.Rsh(q,fracBits).Int64() * int64(s)
}

func fraction(x Fixed) int64 {
	q := new(big.Int).Abs(x.val)
	return q.And(q, fixedFracMask.val).Int64()
}

func round(x Fixed) int64 {
	y := new(big.Int).Add(x.val, fixedHalfValue.val)
	return y.Rsh(y,fracBits).Int64()
}

func one() Fixed {
	return Fixed{new(big.Int).Set(fixedOne.val) }
}

func zero() Fixed {
	return Fixed{new(big.Int) }
}

func iszero(x Fixed) bool {
	return x.val.BitLen() == 0
}

func mul(x, y Fixed) Fixed {
	a := new(big.Int)
	a = a.Mul(x.val,y.val)
	a = a.Add(a, fixedRoundValue.val)
	a = a.Rsh(a,fracBits)
	return Fixed{a }
}

func mulx(x Fixed, yy ...Fixed) Fixed {
	a := new(big.Int).Set(x.val)
	for _,y := range yy {
		a.Mul(a,y.val).Add(a, fixedHalfValue.val).Rsh(a,fracBits)
	}
	return Fixed{a }
}

func muladd1(x, y Fixed) Fixed {
	a := new(big.Int).Mul(x.val,y.val)
	return Fixed{a.Add(a, fixedHalfValue.val).Rsh(a,fracBits).Add(a, fixedOne.val) }
}

func div(x, y Fixed) Fixed {
	a := new(big.Int).Lsh(x.val,fracBits+1)
	return Fixed{a.Div(a,y.val).Add(a, rawFixedOne.val).Rsh(a,1) }
}

func divadd1(x, y Fixed) Fixed {
	a := new(big.Int).Lsh(x.val,fracBits+1)
	return Fixed{a.Div(a,y.val).Add(a, rawFixedOne.val).Rsh(a,1).Add(a, fixedOne.val) }
}

func inv(x Fixed) Fixed {
	a := new(big.Int).Lsh(fixedOne.val,fracBits+1)
	return Fixed{a.Div(a,x.val).Add(a, rawFixedOne.val).Rsh(a,1) }
}

func add(x, y Fixed) Fixed {
	return Fixed{new(big.Int).Add(x.val,y.val) }
}

func sub(x, y Fixed) Fixed {
	return Fixed{new(big.Int).Sub(x.val,y.val) }
}

func addx(x Fixed, y ...Fixed) Fixed {
	z := new(big.Int).Set(x.val)
	for _, a := range y {
		z.Add(z,a.val)
	}
	return Fixed{z }
}

func subx(x Fixed, y ...Fixed) Fixed {
	z := new(big.Int).Set(x.val)
	for _, a := range y {
		z.Sub(z,a.val)
	}
	return Fixed{z }
}

func (x Fixed) lsh(n uint) {
	x.val.Lsh(x.val,n)
}

func (x Fixed) rsh(n uint) {
	x.val.Rsh(x.val,n)
}

func cmpabs(x Fixed, e Fixed) int {
	return x.val.CmpAbs(e.val)
}

var formatString = fmt.Sprintf("%%d+%%0%dx/%d", fracBits/4, fracBits)

func format(x Fixed) string {
	return fmt.Sprintf(formatString, integer(x), fraction(x))
}

func float(x Fixed) float64 {
	f, _ :=  new(big.Float).SetInt(x.val).Float64()
	return f/float64(oneValue)
}
