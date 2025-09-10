package srand

import (
	"crypto/rand"
	"encoding/base64"
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
	charsetLen := len(charset)
	if charsetLen == 0 {
		return "", fmt.Errorf("empty charset")
	}

	// max is the largest multiple of charsetLen that is <= 256
	// All random bytes >= max will be rejected to avoid modulo bias.
	max := 256 - (256 % charsetLen)

	b := make([]byte, n)

	// To minimize calls to rand.Read, we'll read a larger buffer of random
	// bytes at once.
	buffer := make([]byte, n+(n/2)) // Read 1.5x the needed bytes

	b_idx := 0
	for {
		_, err := rand.Read(buffer)
		if err != nil {
			return "", err
		}

		for _, r_byte := range buffer {
			// Check if the byte is in the acceptable range
			if int(r_byte) < max {
				b[b_idx] = charset[int(r_byte)%charsetLen]
				b_idx++
				if b_idx == n {
					return string(b), nil
				}
			}
		}
	}
}

// String returns a random alphanumeric string of length n.
// It is a convenience wrapper around StringFromCharset, using CharsetAlphaNum.
func String(n int) (string, error) {
	return StringFromCharset(n, CharsetAlphaNum)
}

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}

func GenerateRandomToken(n int) (string, error) {
	tokenBytes := make([]byte, n)

	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(tokenBytes), nil
}
