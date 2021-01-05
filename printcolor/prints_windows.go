// +build windows

package printcolor

import (
	"fmt"
	"io"
	"os"
	"syscall"
)

const (
	// color code on windows
	fgBlack2 = iota
	fgBlue2
	fgGreen2
	fgCyan2
	fgRed2
	fgPurple2
	fgYellow2
	fgLightGray2 = 7
	fgWhite2     = 15
)

type windows struct {
	kernel32    *syscall.LazyDLL
	proc        *syscall.LazyProc
	closeHandle *syscall.LazyProc
}

func init() {
	kernel32 := syscall.NewLazyDLL(`kernel32.dll`)
	print = &windows{
		kernel32:    kernel32,
		proc:        kernel32.NewProc(`SetConsoleTextAttribute`),
		closeHandle: kernel32.NewProc(`CloseHandle`),
	}
}

// syscall.Stdout
func (win *windows) Call(w io.Writer, s string, i int) {
	handle, _, _ := win.proc.Call(uintptr(win.GetStd(w)), uintptr(i))
	fmt.Print(s)
	win.closeHandle.Call(handle)
	handle, _, _ = win.proc.Call(uintptr(win.GetStd(w)), uintptr(fgLightGray2))
	win.closeHandle.Call(handle)
}

// Implementation of Windows CMD
func (win *windows) Red(w io.Writer, format string, a ...interface{}) {
	win.Call(w, fmt.Sprintf(format, a...), fgRed2)
}

func (win *windows) Cyan(w io.Writer, format string, a ...interface{}) {
	win.Call(w, fmt.Sprintf(format, a...), fgCyan2)
}

func (win *windows) Blue(w io.Writer, format string, a ...interface{}) {
	win.Call(w, fmt.Sprintf(format, a...), fgBlue2)
}

func (win *windows) White(w io.Writer, format string, a ...interface{}) {
	win.Call(w, fmt.Sprintf(format, a...), fgWhite2)
}

func (win *windows) Black(w io.Writer, format string, a ...interface{}) {
	win.Call(w, fmt.Sprintf(format, a...), fgBlack2)
}

func (win *windows) Green(w io.Writer, format string, a ...interface{}) {
	win.Call(w, fmt.Sprintf(format, a...), fgGreen2)
}

func (win *windows) Yellow(w io.Writer, format string, a ...interface{}) {
	win.Call(w, fmt.Sprintf(format, a...), fgYellow2)
}

func (win *windows) GetStd(w io.Writer) int {
	var std int
	if w == os.Stdout {
		std = int(syscall.Stdout)
	} else if w == os.Stderr {
		std = int(syscall.Stderr)
	}
	return std
}
