package isatty

// Fder is implemented by *os.File and other types that expose a file descriptor.
type Fder interface {
	Fd() uintptr
}

// IsTerminalFile reports whether f is a terminal, when f implements Fd().
// If f is nil or does not expose a file descriptor, it returns false.
func IsTerminalFile(f Fder) bool {
	if f == nil {
		return false
	}
	return IsTerminal(f.Fd())
}

// IsCygwinTerminalFile reports whether f is a Cygwin/MSYS2 terminal, when f
// implements Fd(). If f is nil or does not expose a file descriptor, it returns false.
func IsCygwinTerminalFile(f Fder) bool {
	if f == nil {
		return false
	}
	return IsCygwinTerminal(f.Fd())
}
