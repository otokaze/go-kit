package log

import (
	"fmt"
	"os"
	"time"

	"github.com/otokaze/go-kit/printcolor"
)

var (
	errPrefix  = "[ERROR]"
	infoPrefix = "[INFO]"
	warnPrefix = "[WANR]"
)

func Error(format string, a ...interface{}) {
	ts := time.Now().Local().Format("2006-01-02 15:04:05")
	format = fmt.Sprintf("%s %s %s\n", errPrefix, ts, format)
	printcolor.Fred(os.Stderr, format, a...)
}
func Warn(format string, a ...interface{}) {
	ts := time.Now().Local().Format("2006-01-02 15:04:05")
	format = fmt.Sprintf("%s %s %s\n", warnPrefix, ts, format)
	printcolor.Fyellow(os.Stdout, format, a...)
}
func Info(format string, a ...interface{}) {
	ts := time.Now().Local().Format("2006-01-02 15:04:05")
	format = fmt.Sprintf("%s %s %s\n", infoPrefix, ts, format)
	fmt.Fprintf(os.Stdout, format, a...)
}
