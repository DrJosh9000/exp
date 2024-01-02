package exp

import "testing"

func TestUint128_Lsh(t *testing.T) {
	tests := []struct {
		input, want Uint128
		k           int
	}{
		{
			input: Uint128{hi: 1, lo: 0},
			k:     1,
			want:  Uint128{hi: 2, lo: 0},
		},
		{
			input: Uint128{hi: 0, lo: 1},
			k:     1,
			want:  Uint128{hi: 0, lo: 2},
		},
		{
			input: Uint128{hi: 0, lo: 0x8000000000000000},
			k:     1,
			want:  Uint128{hi: 1, lo: 0},
		},
		{
			input: Uint128{hi: 0x8000000000000000, lo: 0},
			k:     1,
			want:  Uint128{hi: 0, lo: 0},
		},
		{
			input: Uint128{hi: 0, lo: 0xffffffffffffffff},
			k:     1,
			want:  Uint128{hi: 1, lo: 0xfffffffffffffffe},
		},
		{
			input: Uint128{hi: 0xffffffffffffffff, lo: 0xffffffffffffffff},
			k:     1,
			want:  Uint128{hi: 0xffffffffffffffff, lo: 0xfffffffffffffffe},
		},
	}

	for _, test := range tests {
		if got := test.input.Lsh(test.k); got != test.want {
			t.Errorf("(%v).Lsh(%d) = %v, want %v", test.input, test.k, got, test.want)
		}
	}
}
