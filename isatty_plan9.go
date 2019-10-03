// +build plan9

package isatty

import (
	"syscall"
)

func IsTerminal(fd uintptr) bool {
	path, err := syscall.Fd2path(fd)
	if err != nil {
		return false
	}
	return path == "/dev/cons" || path == "/mnt/term/dev/cons"
}
