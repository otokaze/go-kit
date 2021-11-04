package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/otokaze/go-kit/printcolor"
)

var (
	errPrefix  = "[ERROR]"
	infoPrefix = "[INFO]"
	warnPrefix = "[WANR]"
)

func Error(format string, a ...interface{}) { Errorw(os.Stderr, format, a...) }
func Warn(format string, a ...interface{})  { Warnw(os.Stderr, format, a...) }
func Info(format string, a ...interface{})  { Infow(os.Stdout, format, a...) }

func Errorw(w io.Writer, format string, a ...interface{}) {
	ts := time.Now().Local().Format("2006-01-02 15:04:05")
	format = fmt.Sprintf("%s %s %s %s\n", errPrefix, ts, caller(2), format)
	printcolor.Fred(w, format, a...)
}
func Warnw(w io.Writer, format string, a ...interface{}) {
	ts := time.Now().Local().Format("2006-01-02 15:04:05")
	format = fmt.Sprintf("%s %s %s %s\n", warnPrefix, ts, caller(2), format)
	printcolor.Fyellow(w, format, a...)
}
func Infow(w io.Writer, format string, a ...interface{}) {
	ts := time.Now().Local().Format("2006-01-02 15:04:05")
	format = fmt.Sprintf("%s %s %s %s\n", infoPrefix, ts, caller(2), format)
	fmt.Fprintf(w, format, a...)
}

func caller(depth int) string {
	_, file, line, _ := runtime.Caller(depth)
	idx := strings.LastIndexByte(file, '/')
	return file[idx+1:] + ":" + strconv.Itoa(line)
}
