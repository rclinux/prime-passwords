//go:build linux

package main

import "syscall"

// harden marks the process non-dumpable, so a crash never writes passwords
// to a core dump and other programs running as the same user cannot read
// this process's memory (ptrace, /proc/<pid>/mem).
func harden() {
	const prSetDumpable = 4
	_, _, _ = syscall.RawSyscall(syscall.SYS_PRCTL, prSetDumpable, 0, 0)
}
