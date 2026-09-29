// Package gen makes cryptographically secure random strings.
//
// All randomness comes from crypto/rand, which reads the operating system's
// secure random generator (getrandom on Linux, BCryptGenRandom on Windows,
// arc4random on macOS). Characters are picked with rand.Int, so every
// character in a set has exactly equal odds. Nothing is stored or logged.
package gen

import (
	"crypto/rand"
	"math"
	"math/big"
)

// Character sets, matching GRC's Perfect Passwords page.
const (
	Hex       = "0123456789ABCDEF"
	Alnum     = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	Printable = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
)

// Kind describes one of the outputs shown to the user.
type Kind struct {
	Name    string
	Charset string
	Length  int
}

// Bits returns the entropy of one generated string, in bits.
func (k Kind) Bits() float64 {
	return float64(k.Length) * math.Log2(float64(len(k.Charset)))
}

// The three outputs.
var (
	HexKind       = Kind{"64 hexadecimal characters", Hex, 64}
	PrintableKind = Kind{"63 printable ASCII characters", Printable, 63}
	AlnumKind     = Kind{"63 letters and digits", Alnum, 63}
	Kinds         = []Kind{HexKind, PrintableKind, AlnumKind}
)

// String returns length characters drawn uniformly from charset.
func String(charset string, length int) (string, error) {
	max := big.NewInt(int64(len(charset)))
	out := make([]byte, length)
	for i := range out {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		out[i] = charset[n.Int64()]
	}
	return string(out), nil
}

// Generate returns a new string for k.
func (k Kind) Generate() (string, error) {
	return String(k.Charset, k.Length)
}
