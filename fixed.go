package fixed56

import (
	"errors"
)

var ErrOverflow = errors.New("overflow")

func (x Fixed) String() string {
	return format(x)
}

func New(val int64) Fixed {
	return fixed(val)
}

func From(val float64) Fixed {
	return from(val)
}

func One() Fixed {
	return one()
}

func Zero() Fixed {
	return zero()
}

func (x Fixed) Floor() int64 {
	return integer(x)
}

func (x Fixed) Round() int64 {
	return round(x)
}

func (x Fixed) Float() float64 {
	return float(x)
}

func (x Fixed) Mul(y Fixed) Fixed {
	return mul(x,y)
}

func BinCDF(n int64, p Fixed, x int64) Fixed {
	if x < 0 {
		return zero()
	} else if x >= n {
		return one()
	} else {
		return incomplete(n-x, x+1, oneValue-p.fixed56())
	}
}

