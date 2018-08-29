//+build go1.11,js

package isatty

// IsTerminal returns always false if GOOS=js
func IsTerminal(fd uintptr) bool {
	return false
}

// IsCygwinTerminal returns always false if GOOS=js
func IsCygwinTerminal(fd uintptr) bool {
	return false
}
