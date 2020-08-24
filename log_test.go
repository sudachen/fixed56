package fixed56

import (
	"math"
	"math/rand"
	"testing"
)

func TestRange_Log(t *testing.T) {
	rand.Seed(42)
	step := oneValue >> 12
	acc := accuracy{Epsilon: 1 << 5}
	for i := oneValue; i < 27*oneValue; i += step {
		a := randomFixed(int64(i))
		y := log(a)
		want := math.Log(float56(a))
		if ok := acc.update(y, want); !ok {
			t.Errorf("log(%v|%v) => %v|%v, want %v|%v", rawfixed(a), float56(a), y, float(y), From(want), want)
			t.FailNow()
		}
	}
	t.Log(acc)
}

func TestRange_Log2(t *testing.T) {
	rand.Seed(42)
	acc := accuracy{}
	maxarg := int64(100000)
	step := maxarg/200
	for i := int64(0); i < 10000; i += step {
		a := i + rand.Int63n(step)
		y := rawfixed(ilog56(a))
		want := math.Log(float64(a))
		acc.Epsilon = float56(1 << min64(max64(int64(y.val.BitLen()-24), 2),62))
		if ok := acc.update(y, want); !ok {
			t.Errorf("log(%v) => %v|%v, want %v|%v", a, y, float(y), From(want), want)
			t.FailNow()
		}
	}
	t.Log(acc)
}
