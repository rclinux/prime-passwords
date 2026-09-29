package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/rclinux/prime-passwords/internal/gen"
)

// clipboardClearAfter is a variable so tests can shorten it.
var clipboardClearAfter = 30 * time.Second

const aboutText = `## How these passwords are made

Every character comes from your operating system's secure random number
generator — the same source used to make encryption keys. Each character in
a set is exactly as likely as every other, and each one is picked
independently of the rest.

Because of that, **any part of a password is just as random as the whole**.
If a device only accepts 20 characters, take any 20 in a row.

Nothing is sent over a network, and nothing is saved or logged. Each press
of **Refresh** makes brand-new passwords unrelated to the ones before.

## Which one to use

- **64 hexadecimal characters** — exactly 256 bits. Use where a raw key is
  wanted, such as a WPA/WPA2 Wi-Fi pre-shared key entered in hex.
- **63 printable ASCII characters** — the strongest, about 413 bits. Use as a
  Wi-Fi passphrase or shared secret where any keyboard character is allowed.
- **63 letters and digits** — about 375 bits. Use for devices or sites that
  do not accept symbols.

## Keyboard

**F5** or **Ctrl+R** makes new passwords.

## Clipboard

**Copy** puts a password on the clipboard and clears it again after 30
seconds, as long as nothing else has been copied since.

---

Inspired by Steve Gibson's "Perfect Passwords" page at GRC.com.
Prime Passwords is a separate project, not affiliated with GRC.

Version ` + version + ` · MIT License`

type row struct {
	kind    gen.Kind
	value   string
	text    *widget.Label
	copyBtn *widget.Button
}

// mainUI holds the widgets tests need to reach.
type mainUI struct {
	app        fyne.App
	win        fyne.Window
	rows       []*row
	status     *widget.Label
	refreshBtn *widget.Button
	aboutBtn   *widget.Button
}

func runGUI() {
	a := app.NewWithID("io.github.rclinux.primepasswords")
	w := a.NewWindow("Prime Passwords")
	newMainUI(a, w)
	w.Resize(fyne.NewSize(760, 620))
	w.CenterOnScreen()
	w.ShowAndRun()
}

func newMainUI(a fyne.App, w fyne.Window) *mainUI {
	u := &mainUI{app: a, win: w}
	icon := fyne.NewStaticResource("icon.png", iconPNG)
	a.SetIcon(icon)
	w.SetIcon(icon)

	u.status = widget.NewLabel("")
	u.status.Alignment = fyne.TextAlignCenter
	u.status.Importance = widget.LowImportance

	cards := container.NewVBox()
	for _, k := range gen.Kinds {
		r := &row{kind: k}
		r.text = widget.NewLabel("")
		r.text.TextStyle = fyne.TextStyle{Monospace: true}
		r.text.Wrapping = fyne.TextWrapBreak
		r.text.Selectable = true

		r.copyBtn = widget.NewButtonWithIcon("Copy", theme.ContentCopyIcon(), func() {
			u.copy(r.value)
		})
		heading := widget.NewLabelWithStyle(
			fmt.Sprintf("%s  ·  %.0f bits", k.Name, k.Bits()),
			fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
		top := container.NewBorder(nil, nil, nil, r.copyBtn, heading)
		cards.Add(widget.NewCard("", "", container.NewVBox(top, r.text)))
		u.rows = append(u.rows, r)
	}

	u.refreshBtn = widget.NewButtonWithIcon("Refresh", theme.ViewRefreshIcon(), u.refresh)
	u.refreshBtn.Importance = widget.HighImportance
	u.aboutBtn = widget.NewButtonWithIcon("About", theme.InfoIcon(), u.showAbout)

	logo := canvas.NewImageFromResource(icon)
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(56, 56))
	title := widget.NewLabelWithStyle("Prime Passwords", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	title.SizeName = theme.SizeNameHeadingText
	subtitle := widget.NewLabel("Cryptographically secure passwords, made on this computer.")
	header := container.NewBorder(nil, nil, logo, nil, container.NewVBox(title, subtitle))

	buttons := container.NewBorder(nil, nil, u.aboutBtn, u.refreshBtn, u.status)

	w.SetContent(container.NewPadded(container.NewBorder(
		container.NewVBox(header, widget.NewSeparator()),
		buttons, nil, nil,
		container.NewVScroll(cards),
	)))

	w.Canvas().AddShortcut(&desktop.CustomShortcut{KeyName: fyne.KeyR, Modifier: fyne.KeyModifierShortcutDefault},
		func(fyne.Shortcut) { u.refresh() })
	w.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		if e.Name == fyne.KeyF5 {
			u.refresh()
		}
	})

	u.refresh()
	return u
}

// refresh replaces all three passwords with new ones.
func (u *mainUI) refresh() {
	for _, r := range u.rows {
		s, err := r.kind.Generate()
		if err != nil {
			dialog.ShowError(fmt.Errorf("random generator failed: %w", err), u.win)
			return
		}
		r.value = s
		r.text.SetText(s)
	}
	u.status.SetText("New passwords generated at " + time.Now().Format("15:04:05"))
}

// copy puts s on the clipboard and clears it after clipboardClearAfter,
// unless something else has been copied in the meantime.
func (u *mainUI) copy(s string) {
	cb := u.app.Clipboard()
	cb.SetContent(s)
	u.status.SetText(fmt.Sprintf("Copied — clipboard clears in %d seconds", int(clipboardClearAfter.Seconds())))
	time.AfterFunc(clipboardClearAfter, func() {
		fyne.Do(func() {
			if cb.Content() == s {
				cb.SetContent("")
				u.status.SetText("Clipboard cleared")
			}
		})
	})
}

func (u *mainUI) showAbout() {
	body := widget.NewRichTextFromMarkdown(aboutText)
	body.Wrapping = fyne.TextWrapWord
	scroll := container.NewVScroll(body)
	scroll.SetMinSize(fyne.NewSize(560, 420))
	dialog.ShowCustom("About Prime Passwords", "Close", scroll, u.win)
}
