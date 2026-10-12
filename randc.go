// Package randc provides random integers and digit strings based on crypto/rand.
// 包 randc 基于 crypto/rand 提供安全随机整数和数字字符串。
package randc

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// IntN returns a random int in [0, n). It requires n > 0.
// IntN 返回 [0, n) 内的安全随机整数，要求 n > 0。
func IntN(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("randc: n must be positive: %d", n)
	}
	v, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0, fmt.Errorf("randc: random int: %w", err)
	}
	return int(v.Int64()), nil
}

// IntBetween returns a random int in [min(a, b), max(a, b)). Matching bounds return a.
// IntBetween 返回两端点之间的安全随机整数，含小端点、不含大端点；端点相等时返回该值。
func IntBetween(a, b int) (int, error) {
	if a == b {
		return a, nil
	}
	if a > b {
		a, b = b, a
	}
	// Use big.Int so subtraction does not overflow across signed bounds.
	// 用 big.Int 计算跨度，避免整数端点相减溢出。
	base := big.NewInt(int64(a))
	span := new(big.Int).Sub(big.NewInt(int64(b)), base)
	v, err := rand.Int(rand.Reader, span)
	if err != nil {
		return 0, fmt.Errorf("randc: random int: %w", err)
	}
	return int(v.Add(v, base).Int64()), nil
}

// DigitString returns n random ASCII digits, retaining leading zeros. It requires n > 0.
// DigitString 返回恰好 n 位的安全随机数字字符串，保留前导零，要求 n > 0。
func DigitString(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("randc: digit count must be positive: %d", n)
	}
	span := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(n)), nil)
	v, err := rand.Int(rand.Reader, span)
	if err != nil {
		return "", fmt.Errorf("randc: random digits: %w", err)
	}
	return fmt.Sprintf("%0*s", n, v.String()), nil
}
