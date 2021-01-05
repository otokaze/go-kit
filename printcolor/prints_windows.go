package printcolor

import (
	"fmt"
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

func (w *windows) Call(s string, i int) {
	handle, _, _ := w.proc.Call(uintptr(syscall.Stdout), uintptr(i))
	fmt.Print(s)
	w.closeHandle.Call(handle)
	handle, _, _ = w.proc.Call(uintptr(syscall.Stdout), uintptr(fgLightGray2))
	w.closeHandle.Call(handle)
}

// Implementation of Windows CMD
func (w *windows) Red(format string, a ...interface{}) {
	w.Call(fmt.Sprintf(format, a...), fgRed2)
}

func (w *windows) Cyan(format string, a ...interface{}) {
	fmt.Println(fgRed2, fgBlack2, fgWhite2)
	w.Call(fmt.Sprintf(format, a...), fgCyan2)
}

func (w *windows) Blue(format string, a ...interface{}) {
	w.Call(fmt.Sprintf(format, a...), fgBlue2)
}

func (w *windows) White(format string, a ...interface{}) {
	w.Call(fmt.Sprintf(format, a...), fgWhite2)
}

func (w *windows) Black(format string, a ...interface{}) {
	w.Call(fmt.Sprintf(format, a...), fgBlack2)
}

func (w *windows) Green(format string, a ...interface{}) {
	w.Call(fmt.Sprintf(format, a...), fgGreen2)
}

func (w *windows) Yellow(format string, a ...interface{}) {
	w.Call(fmt.Sprintf(format, a...), fgYellow2)
}
