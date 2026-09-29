package main

import (
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
)

// TestScreenshots renders README images into docs/screenshots.
// Run with: SCREENSHOTS=1 go test -run TestScreenshots .
func TestScreenshots(t *testing.T) {
	if os.Getenv("SCREENSHOTS") == "" {
		t.Skip("set SCREENSHOTS=1 to render README images")
	}
	for _, variant := range []fyne.ThemeVariant{theme.VariantDark, theme.VariantLight} {
		shot(t, variant, false)
		shot(t, variant, true)
	}
}

func shot(t *testing.T, variant fyne.ThemeVariant, about bool) {
	a := test.NewTempApp(t)
	a.Settings().SetTheme(&fixedVariant{theme.DefaultTheme(), variant})
	w := test.NewTempWindow(t, nil)
	u := newMainUI(a, w)
	u.status.SetText("New passwords generated at 12:00:00")
	w.Resize(fyne.NewSize(760, 560))
	name := "main"
	if about {
		w.Resize(fyne.NewSize(760, 640))
		u.showAbout()
		name = "about"
	}
	mode := map[fyne.ThemeVariant]string{theme.VariantDark: "dark", theme.VariantLight: "light"}[variant]
	path := filepath.Join("docs", "screenshots", "linux-"+name+"-"+mode+".png")
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, w.Canvas().Capture()); err != nil {
		t.Fatal(err)
	}
}

type fixedVariant struct {
	fyne.Theme
	v fyne.ThemeVariant
}

func (f *fixedVariant) Color(n fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return f.Theme.Color(n, f.v)
}
