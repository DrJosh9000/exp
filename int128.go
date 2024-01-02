package exp

import (
	"cmp"
	"math/bits"
)

type Uint128 struct{ hi, lo uint64 }

func NewUint128(hi, lo uint64) Uint128 { return Uint128{hi: hi, lo: lo} }

func (x Uint128) Add(y Uint128) (z Uint128) {
	l, c := bits.Add64(x.lo, y.lo, 0)
	z.lo = l
	z.hi = x.hi + y.hi + c
	return z
}

func (x Uint128) Sub(y Uint128) (z Uint128) {
	l, b := bits.Sub64(x.lo, y.lo, 0)
	z.lo = l
	z.hi, _ = bits.Sub64(x.hi, y.hi, b)
	return z
}

func (x Uint128) Mul(y Uint128) (z Uint128) {
	z.hi, z.lo = bits.Mul64(x.lo, y.lo)
	z.hi += x.hi*y.lo + x.lo*y.hi
	return z
}

func (x Uint128) Cmp(y Uint128) int {
	if x.hi != y.hi {
		return cmp.Compare(x.hi, y.hi)
	}
	return cmp.Compare(x.lo, y.lo)
}

// func (x Uint128) Div(y uint64) (z Uint128) {

// }

func (x Uint128) Lsh(k int) Uint128 {
	x.hi <<= k
	x.hi += x.lo >> (64 - k)
	x.lo <<= k
	return x
}

func (x Uint128) Rsh(k int) Uint128 {
	x.lo >>= k
	x.lo += x.hi << (64 - k)
	x.hi >>= k
	return x
}
