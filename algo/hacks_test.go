package algo

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestIntegerPredict(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			var signal []int
			period := 10 + rand.Intn(50)
			for j := 0; j < period; j++ {
				signal = append(signal, rand.Intn(101)-50)
			}

			// t.Logf("signal[:period] = %v", signal)

			slope := rand.Intn(101) - 50
			for j := 1; j <= 10; j++ {
				for _, y := range signal[:period] {
					signal = append(signal, y+j*slope)
				}
			}

			// t.Logf("period = %d", period)
			// t.Logf("slope = %d", slope)

			start := 5*period + rand.Intn(period)
			for j := start; j < len(signal); j++ {
				got := IntegerPredict(signal[:start], j)
				want := signal[j]
				if got != want {
					t.Errorf("IntegerPredict(signal[:%d], %d) = %d, want %d", start, j, got, want)
				}
			}
		})
	}
}
