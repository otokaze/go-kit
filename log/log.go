package log

import (
	"fmt"
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

func Error(format string, a ...interface{}) {
	ts := time.Now().Local().Format("2006-01-02 15:04:05")
	format = fmt.Sprintf("%s %s %s %s\n", errPrefix, ts, caller(1), format)
	printcolor.Fred(os.Stderr, format, a...)
}
func Warn(format string, a ...interface{}) {
	ts := time.Now().Local().Format("2006-01-02 15:04:05")
	format = fmt.Sprintf("%s %s %s %s\n", warnPrefix, ts, caller(1), format)
	printcolor.Fyellow(os.Stdout, format, a...)
}
func Info(format string, a ...interface{}) {
	ts := time.Now().Local().Format("2006-01-02 15:04:05")
	format = fmt.Sprintf("%s %s %s %s\n", infoPrefix, ts, caller(1), format)
	fmt.Fprintf(os.Stdout, format, a...)
}

func caller(depth int) string {
	_, file, line, _ := runtime.Caller(depth)
	idx := strings.LastIndexByte(file, '/')
	return file[idx+1:] + ":" + strconv.Itoa(line)
}
