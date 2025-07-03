package srand

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Intn returns, uniformly at random, an integer in [0, n).
// It returns an error if secure randomness fails.
func Intn(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("max must be > 0")
	}
	out, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0, err
	}
	return int(out.Int64()), nil
}

func Shuffle[T any](arr []T) error {
	for i := len(arr) - 1; i > 0; i-- {
		j, err := Intn(i + 1)
		if err != nil {
			return err
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	return nil
}

func Perm(n int) ([]int, error) {
	out := make([]int, n)
	for i := range out {
		out[i] = i
	}
	return out, Shuffle(out)
}

func Bool() (bool, error) {
	n, err := Intn(2)
	return n == 1, err
}

func FromSlice[T any](arr []T) (T, error) {
	i, err := Intn(len(arr))
	if err != nil {
		var zero T
		return zero, err
	}
	return arr[i], nil
}

const (
	CharsetAlpha      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetAlphaLower = "abcdefghijklmnopqrstuvwxyz"
	CharsetAlphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetAlphaNum   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	CharsetDigits     = "0123456789"
	CharsetHex        = "0123456789abcdef"
	CharsetBase64     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

func StringFromCharset(n int, charset string) (string, error) {
	if n <= 0 {
		return "", nil
	}
	if len(charset) == 0 {
		return "", fmt.Errorf("empty charset")
	}

	b := make([]byte, n)
	for i := range b {
		idx, err := Intn(len(charset))
		if err != nil {
			return "", err
		}
		b[i] = charset[idx]
	}
	return string(b), nil
}

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}
