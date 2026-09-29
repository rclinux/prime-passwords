package main

import (
	"testing"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
)

func newTestUI(t *testing.T) *mainUI {
	a := test.NewTempApp(t)
	w := a.NewWindow("Prime Passwords")
	t.Cleanup(w.Close)
	return newMainUI(a, w)
}

func values(u *mainUI) []string {
	var out []string
	for _, r := range u.rows {
		out = append(out, r.text.Text)
	}
	return out
}

func TestStartsFilled(t *testing.T) {
	u := newTestUI(t)
	for i, r := range u.rows {
		if len(r.text.Text) != r.kind.Length || r.text.Text != r.value {
			t.Errorf("row %d: shown %q, stored %q", i, r.text.Text, r.value)
		}
	}
}

func TestRefreshButton(t *testing.T) {
	u := newTestUI(t)
	before := values(u)
	test.Tap(u.refreshBtn)
	after := values(u)
	for i := range before {
		if before[i] == after[i] {
			t.Errorf("row %d unchanged after Refresh", i)
		}
	}
}

func TestRefreshKeys(t *testing.T) {
	u := newTestUI(t)
	before := values(u)
	u.win.Canvas().OnTypedKey()(&fyne.KeyEvent{Name: fyne.KeyF5})
	if values(u)[0] == before[0] {
		t.Error("F5 did not refresh")
	}
}

// captureClears replaces afterFunc so scheduled clears run only when the
// test calls them.
func captureClears(t *testing.T) *[]func() {
	var pending []func()
	afterFunc = func(d time.Duration, f func()) {
		if d != 30*time.Second {
			t.Errorf("clear delay = %v, want 30s", d)
		}
		pending = append(pending, f)
	}
	t.Cleanup(func() { afterFunc = func(d time.Duration, f func()) { time.AfterFunc(d, f) } })
	return &pending
}

func TestCopyAndClear(t *testing.T) {
	pending := captureClears(t)
	u := newTestUI(t)
	cb := u.app.Clipboard()
	for _, r := range u.rows {
		test.Tap(r.copyBtn)
		if got := cb.Content(); got != r.value {
			t.Fatalf("clipboard = %q, want %q", got, r.value)
		}
	}
	// Only the last copy is still on the clipboard, so only its clear acts.
	for _, f := range *pending {
		f()
	}
	if got := cb.Content(); got != "" {
		t.Errorf("clipboard not cleared: %q", got)
	}
}

func TestCopyDoesNotClearOtherContent(t *testing.T) {
	pending := captureClears(t)
	u := newTestUI(t)
	cb := u.app.Clipboard()
	test.Tap(u.rows[0].copyBtn)
	cb.SetContent("something the user copied later")
	for _, f := range *pending {
		f()
	}
	if got := cb.Content(); got != "something the user copied later" {
		t.Errorf("clipboard overwritten: %q", got)
	}
}
