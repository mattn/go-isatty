// +build windows
// +build !appengine

package isatty

import (
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

const (
	fileNameInfo uintptr = 2
	fileTypePipe         = 3
)

var (
	kernel32                         = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleMode               = kernel32.NewProc("GetConsoleMode")
	procGetFileInformationByHandleEx = kernel32.NewProc("GetFileInformationByHandleEx")
	procGetFileType                  = kernel32.NewProc("GetFileType")
)

func init() {
	// Check if GetFileInformationByHandleEx is available.
	if procGetFileInformationByHandleEx.Find() != nil {
		procGetFileInformationByHandleEx = nil
	}
}

// IsTerminal return true if the file descriptor is terminal.
func IsTerminal(fd uintptr) bool {
	var st uint32
	r, _, e := syscall.Syscall(procGetConsoleMode.Addr(), 2, fd, uintptr(unsafe.Pointer(&st)), 0)
	return r != 0 && e == 0
}

// IsCygwinTerminal() return true if the file descriptor is a cygwin or msys2
// terminal. This is also always false on this environment.
func IsCygwinTerminal(fd uintptr) bool {
	if procGetFileInformationByHandleEx == nil {
		return false
	}

	// Cygwin/msys's pty is a pipe.
	ft, _, e := syscall.Syscall(procGetFileType.Addr(), 1, fd, 0, 0)
	if ft != fileTypePipe || e != 0 {
		return false
	}

	var buf [2 + syscall.MAX_PATH]uint16
	r, _, e := syscall.Syscall6(procGetFileInformationByHandleEx.Addr(),
		4, fd, fileNameInfo, uintptr(unsafe.Pointer(&buf)),
		uintptr(len(buf)*2), 0, 0)
	if r == 0 || e != 0 {
		return false
	}

	l := *(*uint32)(unsafe.Pointer(&buf))
	name := string(utf16.Decode(buf[2 : 2+l/2]))
	if !strings.HasPrefix(name, `\msys-`) && !strings.HasPrefix(name, `\cygwin-`) {
		return false
	}

	token := strings.Split(name, "-")
	if len(token) < 3 || !strings.HasPrefix(token[2], "pty") {
		return false
	}
	return true
}
