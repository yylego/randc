package randc_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/randc"
)

func TestIntN(t *testing.T) {
	for _, n := range []int{1, 10, 1000} {
		for range 100 {
			v, err := randc.IntN(n)
			require.NoError(t, err)
			require.GreaterOrEqual(t, v, 0)
			require.Less(t, v, n)
		}
	}
}

func TestIntBetween(t *testing.T) {
	for _, bounds := range [][2]int{{10, 20}, {20, 10}, {-20, -10}, {-10, 10}, {5, 6}} {
		for range 100 {
			v, err := randc.IntBetween(bounds[0], bounds[1])
			require.NoError(t, err)
			require.GreaterOrEqual(t, v, min(bounds[0], bounds[1]))
			require.Less(t, v, max(bounds[0], bounds[1]))
		}
	}
	for _, n := range []int{-5, 0, 5} {
		v, err := randc.IntBetween(n, n)
		require.NoError(t, err)
		require.Equal(t, n, v)
	}
}

func TestDigitString(t *testing.T) {
	for _, n := range []int{1, 6, 32, 128} {
		for range 100 {
			text, err := randc.DigitString(n)
			require.NoError(t, err)
			require.Len(t, text, n)
			for _, digit := range text {
				require.GreaterOrEqual(t, digit, '0')
				require.LessOrEqual(t, digit, '9')
			}
		}
	}
}
