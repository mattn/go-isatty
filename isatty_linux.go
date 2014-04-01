// +build linux

package isatty

import (
	"syscall"
	"unsafe"
)

const ioctlReadTermios = syscall.TCGETS

func IsTerminal(fd uintptr) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, fd, ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}
