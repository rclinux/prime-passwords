package gen

import (
	"strings"
	"testing"
)

func TestPrintableSet(t *testing.T) {
	if len(Printable) != 94 {
		t.Fatalf("printable set has %d chars, want 94", len(Printable))
	}
	for c := byte('!'); c <= '~'; c++ {
		if !strings.ContainsRune(Printable, rune(c)) {
			t.Errorf("printable set missing %q", c)
		}
	}
	if len(Alnum) != 62 || len(Hex) != 16 {
		t.Fatalf("alnum=%d hex=%d, want 62 and 16", len(Alnum), len(Hex))
	}
}

func TestLengthsAndCharsets(t *testing.T) {
	for _, k := range Kinds {
		for i := 0; i < 1000; i++ {
			s, err := k.Generate()
			if err != nil {
				t.Fatal(err)
			}
			if len(s) != k.Length {
				t.Fatalf("%s: length %d, want %d", k.Name, len(s), k.Length)
			}
			for _, c := range s {
				if !strings.ContainsRune(k.Charset, c) {
					t.Fatalf("%s: %q not in charset", k.Name, c)
				}
			}
		}
	}
}

func TestNoRepeats(t *testing.T) {
	seen := map[string]bool{}
	for i := 0; i < 10000; i++ {
		s, err := AlnumKind.Generate()
		if err != nil {
			t.Fatal(err)
		}
		if seen[s] {
			t.Fatal("repeated output")
		}
		seen[s] = true
	}
}

// TestUniform runs a chi-square test on each character set. The limits are
// the 0.9999 quantile for the matching degrees of freedom, so a correct
// generator fails about once in 10,000 runs.
func TestUniform(t *testing.T) {
	limits := map[int]float64{15: 44.3, 61: 111.0, 93: 152.6}
	for _, k := range Kinds {
		n := len(k.Charset)
		const samples = 200000
		s, err := String(k.Charset, samples)
		if err != nil {
			t.Fatal(err)
		}
		counts := map[rune]int{}
		for _, c := range s {
			counts[c]++
		}
		if len(counts) != n {
			t.Fatalf("%s: saw %d distinct chars, want %d", k.Name, len(counts), n)
		}
		want := float64(samples) / float64(n)
		chi := 0.0
		for _, got := range counts {
			d := float64(got) - want
			chi += d * d / want
		}
		if chi > limits[n-1] {
			t.Errorf("%s: chi-square %.1f exceeds %.1f", k.Name, chi, limits[n-1])
		}
	}
}

func TestBits(t *testing.T) {
	if b := HexKind.Bits(); b != 256 {
		t.Errorf("hex bits = %v, want 256", b)
	}
	if b := PrintableKind.Bits(); b < 412 || b > 414 {
		t.Errorf("printable bits = %v, want ~413", b)
	}
}
