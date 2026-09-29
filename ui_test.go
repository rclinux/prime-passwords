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

func TestCopyAndClear(t *testing.T) {
	clipboardClearAfter = 50 * time.Millisecond
	t.Cleanup(func() { clipboardClearAfter = 30 * time.Second })

	u := newTestUI(t)
	cb := u.app.Clipboard()
	for _, r := range u.rows {
		test.Tap(r.copyBtn)
		if got := cb.Content(); got != r.value {
			t.Fatalf("clipboard = %q, want %q", got, r.value)
		}
	}
	time.Sleep(200 * time.Millisecond)
	if got := cb.Content(); got != "" {
		t.Errorf("clipboard not cleared: %q", got)
	}
}

func TestCopyDoesNotClearOtherContent(t *testing.T) {
	clipboardClearAfter = 50 * time.Millisecond
	t.Cleanup(func() { clipboardClearAfter = 30 * time.Second })

	u := newTestUI(t)
	cb := u.app.Clipboard()
	test.Tap(u.rows[0].copyBtn)
	cb.SetContent("something the user copied later")
	time.Sleep(200 * time.Millisecond)
	if got := cb.Content(); got != "something the user copied later" {
		t.Errorf("clipboard overwritten: %q", got)
	}
}
