//go:build windows

package main

import (
	"os"
	"syscall"
)

// attachConsole connects a windowed (-H windowsgui) build to the console it
// was started from, so terminal flags like --cli print their output.
func attachConsole() {
	const attachParentProcess = ^uintptr(0) // (DWORD)-1
	r, _, _ := syscall.NewLazyDLL("kernel32.dll").NewProc("AttachConsole").Call(attachParentProcess)
	if r == 0 {
		return // started from Explorer: no console to attach to
	}
	if f, err := os.OpenFile("CONOUT$", os.O_WRONLY, 0); err == nil {
		os.Stdout = f
		os.Stderr = f
	}
}
