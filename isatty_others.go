// +build !windows
// +build !appengine
// +build !ppc64le

package isatty

// IsCygwinTerminal() return true if the file descriptor is a cygwin or msys2
// terminal. This is also always false on this environment.
func IsCygwinTerminal(fd uintptr) bool {
	return false
}
