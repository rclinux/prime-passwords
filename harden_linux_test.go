//go:build linux

package main

import (
	"syscall"
	"testing"
)

func TestHardenDisablesDumps(t *testing.T) {
	harden()
	const prGetDumpable = 3
	v, _, errno := syscall.RawSyscall(syscall.SYS_PRCTL, prGetDumpable, 0, 0)
	if errno != 0 || v != 0 {
		t.Fatalf("dumpable = %d (errno %v), want 0", v, errno)
	}
}
