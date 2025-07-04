# 🔐 srand — Secure Random Utilities for Go

`srand` is a lightweight Go package that provides **crypto-safe random number utilities**.  
Built on top of `crypto/rand`, it's ideal for applications that require **secure, unpredictable randomness**, such as:

- 🧠 Games & Gambling (GLI-compliant RNG)
- 🎲 Lotteries, Token Generators
- 🔐 Authentication Tokens, OTP, Captchas
- 🤖 Security-focused APIs & simulations

That means:

✅ No need to seed manually

✅ Safe for generating tokens, secrets, or regulated game logic

✅ Compliant with standards like GLI 11

---

## ✨ Features

- ✅ Secure `Intn(n)` — like `math/rand.Intn` but cryptographically secure
- 🔁 Secure `Shuffle()` — Fisher–Yates shuffle using crypto
- 🔢 Secure `Perm(n)` — random permutation of integers
- 🎯 `FromSlice()` — pick a random element from a slice
- 🔠 `StringFromCharset()` — generate secure random strings from custom charsets
- 🔢 `Bool()` — secure random true/false
- 🧱 Predefined charsets: `digits`, `alphanumeric`, `base64`, `hex`, etc.

---

## 🚀 Install

```bash
go get github.com/skyrocket-qy/srand
```

## 👨‍💻 Example

### Intn

```go
n, err := srand.Intn(100)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Random number [0,100):", n)
```

### Shuffle

```go
items := []string{"apple", "banana", "cherry", "date"}
err := srand.Shuffle(items)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Shuffled:", items)
```

### Random String

```go
code, _ := srand.StringFromCharset(8, srand.CharsetAlphaNum)
fmt.Println("Verification Code:", code)
```

Bug reports, suggestions, or PRs are appreciated! Let's make secure randomness easier for everyone
