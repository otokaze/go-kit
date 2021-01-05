package printcolor

import (
	"fmt"
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
	Red(format string, a ...interface{})
	Cyan(format string, a ...interface{})
	Blue(format string, a ...interface{})
	White(format string, a ...interface{})
	Black(format string, a ...interface{})
	Green(format string, a ...interface{})
	Yellow(format string, a ...interface{})
}

func init() { print = &linux{} }

func Red(format string, a ...interface{})    { print.Red(format, a...) }
func Cyan(format string, a ...interface{})   { print.Cyan(format, a...) }
func Blue(format string, a ...interface{})   { print.Blue(format, a...) }
func White(format string, a ...interface{})  { print.White(format, a...) }
func Black(format string, a ...interface{})  { print.Black(format, a...) }
func Green(format string, a ...interface{})  { print.Green(format, a...) }
func Yellow(format string, a ...interface{}) { print.Yellow(format, a...) }

type linux struct{}

// Implementation of Linux terminal
func (*linux) Red(format string, a ...interface{}) {
	var s = fmt.Sprintf(format, a...)
	fmt.Printf("\033[1;%dm%s\033[0m", fgRed, s)
}

func (*linux) Cyan(format string, a ...interface{}) {
	var s = fmt.Sprintf(format, a...)
	fmt.Printf("\033[1;%dm%s\033[0m", fgCyan, s)
}

func (*linux) Blue(format string, a ...interface{}) {
	var s = fmt.Sprintf(format, a...)
	fmt.Printf("\033[1;%dm%s\033[0m", fgBlue, s)
}

func (*linux) White(format string, a ...interface{}) {
	var s = fmt.Sprintf(format, a...)
	fmt.Printf("\033[1;%dm%s\033[0m", fgWhite, s)
}

func (*linux) Black(format string, a ...interface{}) {
	var s = fmt.Sprintf(format, a...)
	fmt.Printf("\033[1;%dm%s\033[0m", fgBlack, s)
}

func (*linux) Green(format string, a ...interface{}) {
	var s = fmt.Sprintf(format, a...)
	fmt.Printf("\033[1;%dm%s\033[0m", fgGreen, s)
}

func (*linux) Yellow(format string, a ...interface{}) {
	var s = fmt.Sprintf(format, a...)
	fmt.Printf("\033[1;%dm%s\033[0m", fgYellow, s)
}
