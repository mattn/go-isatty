// Package isatty implements interface to isatty.
//
// Prefer IsTerminal/IsCygwinTerminal with an explicit file descriptor, or
// IsTerminalFile/IsCygwinTerminalFile when you have an *os.File (or any
// value with Fd() uintptr) from an io.Reader/io.Writer (#85, #80).
package isatty
