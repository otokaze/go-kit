package printcolor

import (
	"fmt"
	"io"
	"os"
)

var print Print

const (
	// color code on linux
	fgBlack  = 30
	bgBlack  = 40
	fgRed    = 31
	bgRed    = 41
	fgGreen  = 32
	bgGreen  = 42
	fgYellow = 33
	bgYellow = 43
	fgBlue   = 34
	bgBlue   = 44
	fgPurple = 35
	bgPurple = 45
	fgCyan   = 36
	bgCyan   = 46
	fgWhite  = 37
	bgWhite  = 47
)

type Print interface {
	Red(w io.Writer, format string, a ...interface{})
	Cyan(w io.Writer, format string, a ...interface{})
	Blue(w io.Writer, format string, a ...interface{})
	White(w io.Writer, format string, a ...interface{})
	Black(w io.Writer, format string, a ...interface{})
	Green(w io.Writer, format string, a ...interface{})
	Yellow(w io.Writer, format string, a ...interface{})
}

func init() { print = &linux{} }

func Red(format string, a ...interface{})    { print.Red(os.Stdout, format, a...) }
func Cyan(format string, a ...interface{})   { print.Cyan(os.Stdout, format, a...) }
func Blue(format string, a ...interface{})   { print.Blue(os.Stdout, format, a...) }
func White(format string, a ...interface{})  { print.White(os.Stdout, format, a...) }
func Black(format string, a ...interface{})  { print.Black(os.Stdout, format, a...) }
func Green(format string, a ...interface{})  { print.Green(os.Stdout, format, a...) }
func Yellow(format string, a ...interface{}) { print.Yellow(os.Stdout, format, a...) }

func Fred(w io.Writer, format string, a ...interface{})    { print.Red(w, format, a...) }
func Fcyan(w io.Writer, format string, a ...interface{})   { print.Cyan(w, format, a...) }
func Fblue(w io.Writer, format string, a ...interface{})   { print.Blue(w, format, a...) }
func Fwhite(w io.Writer, format string, a ...interface{})  { print.White(w, format, a...) }
func Fblack(w io.Writer, format string, a ...interface{})  { print.Black(w, format, a...) }
func Fgreen(w io.Writer, format string, a ...interface{})  { print.Green(w, format, a...) }
func Fyellow(w io.Writer, format string, a ...interface{}) { print.Yellow(w, format, a...) }

type linux struct{}

// Implementation of Linux terminal
func (*linux) Red(w io.Writer, format string, a ...interface{}) {
	var s = fmt.Sprintf(format, a...)
	fmt.Fprintf(w, "\033[1;%dm%s\033[0m", fgRed, s)
}

func (*linux) Cyan(w io.Writer, format string, a ...interface{}) {
	var s = fmt.Sprintf(format, a...)
	fmt.Fprintf(w, "\033[1;%dm%s\033[0m", fgCyan, s)
}

func (*linux) Blue(w io.Writer, format string, a ...interface{}) {
	var s = fmt.Sprintf(format, a...)
	fmt.Fprintf(w, "\033[1;%dm%s\033[0m", fgBlue, s)
}

func (*linux) White(w io.Writer, format string, a ...interface{}) {
	var s = fmt.Sprintf(format, a...)
	fmt.Fprintf(w, "\033[1;%dm%s\033[0m", fgWhite, s)
}

func (*linux) Black(w io.Writer, format string, a ...interface{}) {
	var s = fmt.Sprintf(format, a...)
	fmt.Fprintf(w, "\033[1;%dm%s\033[0m", fgBlack, s)
}

func (*linux) Green(w io.Writer, format string, a ...interface{}) {
	var s = fmt.Sprintf(format, a...)
	fmt.Fprintf(w, "\033[1;%dm%s\033[0m", fgGreen, s)
}

func (*linux) Yellow(w io.Writer, format string, a ...interface{}) {
	var s = fmt.Sprintf(format, a...)
	fmt.Fprintf(w, "\033[1;%dm%s\033[0m", fgYellow, s)
}
