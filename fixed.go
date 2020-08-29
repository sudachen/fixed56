package fixed56

import (
	"errors"
)

var ErrOverflow = errors.New("overflow")

func (x Fixed) String() string {
	return x.format()
}

func New(val int64) Fixed {
	return fixed(val)
}

func From(val float64) Fixed {
	return from(val)
}

func One() Fixed {
	return fixedOne
}

func Zero() Fixed {
	return Fixed{}
}

func (x Fixed) Floor() int64 {
	return x.floor()
}

func (x Fixed) Ceil() int64 {
	return x.ceil()
}

func (x Fixed) Round() int64 {
	return x.round()
}

func (x Fixed) Float() float64 {
	return x.float()
}

func (x Fixed) Mul(y Fixed) Fixed {
	return mul(x,y)
}

func (x Fixed) Div(y Fixed) Fixed {
	return div(x,y)
}

func (x Fixed) Add(y Fixed) Fixed {
	return add(x,y)
}

func (x Fixed) Sub(y Fixed) Fixed {
	return sub(x,y)
}

func BinCDF(n int64, p Fixed, x int64) Fixed {
	if x < 0 {
		return Zero()
	} else if x >= n {
		return One()
	} else {
		return incomplete(n-x, x+1, oneValue-p.fixed56())
	}
}

