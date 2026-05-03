package isatty

import "io"

// IsWriterTerminal return true if the writer is terminal.
func IsWriterTerminal(w io.Writer) bool {
	if f, ok := w.(filelike); ok {
		return IsTerminal(f.Fd())
	}
	return false
}

// IsWriterCygwinTerminal return true if the writer is a cygwin or msys2 terminal.
func IsWriterCygwinTerminal(w io.Writer) bool {
	if f, ok := w.(filelike); ok {
		return IsCygwinTerminal(f.Fd())
	}
	return false
}

type filelike interface {
	Fd() uintptr
}
