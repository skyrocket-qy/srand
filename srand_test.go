package srand

import (
	"testing"
)

func TestIntn(t *testing.T) {
	for i := 1; i <= 100; i++ {
		n, err := Intn(10)
		if err != nil {
			t.Fatalf("Intn error: %v", err)
		}
		if n < 0 || n >= 10 {
			t.Fatalf("Intn returned out of range value: %d", n)
		}
	}

	_, err := Intn(0)
	if err == nil {
		t.Fatal("Expected error for Intn(0)")
	}
	_, err = Intn(-5)
	if err == nil {
		t.Fatal("Expected error for Intn(-5)")
	}
}

func TestShuffle(t *testing.T) {
	var empty []int
	err := Shuffle(empty)
	if err != nil {
		t.Fatalf("Shuffle empty slice error: %v", err)
	}

	single := []int{42}
	err = Shuffle(single)
	if err != nil {
		t.Fatalf("Shuffle single element slice error: %v", err)
	}

	arr := []int{1, 2, 3, 4, 5}
	err = Shuffle(arr)
	if err != nil {
		t.Fatalf("Shuffle error: %v", err)
	}
	if len(arr) != 5 {
		t.Fatalf("Shuffle changed array size")
	}
}

func TestPerm(t *testing.T) {
	out, err := Perm(0)
	if err != nil {
		t.Fatalf("Perm(0) error: %v", err)
	}
	if len(out) != 0 {
		t.Fatalf("Perm(0) returned non-empty slice")
	}

	perm, err := Perm(5)
	if err != nil {
		t.Fatalf("Perm error: %v", err)
	}
	if len(perm) != 5 {
		t.Fatalf("Perm returned wrong length")
	}
	found := map[int]bool{}
	for _, v := range perm {
		found[v] = true
	}
	for i := range 5 {
		if !found[i] {
			t.Fatalf("Perm missing value: %d", i)
		}
	}
}

func TestBool(t *testing.T) {
	count := 0
	for range 100 {
		b, err := Bool()
		if err != nil {
			t.Fatalf("Bool error: %v", err)
		}
		if b {
			count++
		}
	}
	if count == 0 || count == 100 {
		t.Fatalf("Bool distribution seems broken: %d trues", count)
	}
}

func TestFromSlice(t *testing.T) {
	var arr []int
	v, err := FromSlice(arr)
	if err == nil {
		t.Fatal("Expected error for empty slice")
	}
	var zero int
	if v != zero {
		t.Fatalf("Expected zero value for empty slice, got %v", v)
	}

	words := []string{"a", "b", "c"}
	for range 10 {
		w, err := FromSlice(words)
		if err != nil {
			t.Fatalf("FromSlice error: %v", err)
		}
		found := false
		for _, word := range words {
			if word == w {
				found = true
			}
		}
		if !found {
			t.Fatalf("FromSlice returned unknown value: %s", w)
		}
	}
}

func TestStringFromCharset(t *testing.T) {
	s, err := StringFromCharset(0, CharsetAlphaNum)
	if err != nil || s != "" {
		t.Fatalf("Expected empty string and nil error, got %q, %v", s, err)
	}

	_, err = StringFromCharset(5, "")
	if err == nil {
		t.Fatal("Expected error for empty charset")
	}

	s, err = StringFromCharset(16, CharsetAlphaNum)
	if err != nil {
		t.Fatalf("StringFromCharset error: %v", err)
	}
	if len(s) != 16 {
		t.Fatalf("StringFromCharset returned wrong length: %d", len(s))
	}
}

func TestBytes(t *testing.T) {
	b, err := Bytes(32)
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	if len(b) != 32 {
		t.Fatalf("Bytes returned wrong length: %d", len(b))
	}
}
