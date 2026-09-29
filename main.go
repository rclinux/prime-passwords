// Prime Passwords makes cryptographically secure random passwords on your
// own computer, in the spirit of GRC's "Perfect Passwords" page.
//
// With no arguments it opens the desktop window. With a flag it prints to
// the terminal instead:
//
//	prime-passwords --cli     all three passwords
//	prime-passwords --hex     64 hex characters only
//	prime-passwords --ascii   63 printable ASCII characters only
//	prime-passwords --alnum   63 letters and digits only
package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/rclinux/prime-passwords/internal/gen"
)

const version = "0.1.0"

//go:embed assets/icon.png
var iconPNG []byte

func main() {
	harden()
	if len(os.Args) > 1 {
		attachConsole()
	}

	cli := flag.Bool("cli", false, "print all three passwords and exit")
	hex := flag.Bool("hex", false, "print 64 hex characters and exit")
	ascii := flag.Bool("ascii", false, "print 63 printable ASCII characters and exit")
	alnum := flag.Bool("alnum", false, "print 63 letters and digits and exit")
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	switch {
	case *showVersion:
		fmt.Println("Prime Passwords", version)
	case *hex:
		printOne(gen.HexKind)
	case *ascii:
		printOne(gen.PrintableKind)
	case *alnum:
		printOne(gen.AlnumKind)
	case *cli:
		for i, k := range gen.Kinds {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("%s (%.0f bits):\n", k.Name, k.Bits())
			printOne(k)
		}
	default:
		runGUI()
	}
}

func printOne(k gen.Kind) {
	s, err := k.Generate()
	if err != nil {
		fmt.Fprintln(os.Stderr, "prime-passwords: random generator failed:", err)
		os.Exit(1)
	}
	fmt.Println(s)
}
