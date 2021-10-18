package debug

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type LogLevel int

const (
	LogLevelNone LogLevel = iota
	LogLevelFatal
	LogLevelError
	LogLevelWarning
	LogLevelInfo
	LogLevelDebug
)

const LogLevelDefault = LogLevelInfo

const (
	red     = 1
	green   = 2
	yellow  = 3
	magenta = 5
	cyan    = 6
)

const fmtColorEscape = "\033[3%dm"
const colorEscapeReset = "\033[0m"

const (
	fmtPrefix     = "[%s]"
	LogFmtDebug   = "[DEBUG]" + colorEscapeReset
	LogFmtInfo    = "[INFO ]" + colorEscapeReset
	LogFmtWarning = "[WARN ]" + colorEscapeReset
	LogFmtError   = "[ERROR]" + colorEscapeReset
	LogFmtFatal   = "[FATAL]" + colorEscapeReset
)

type Logger struct {
	prefix string
	io.Writer
	level        LogLevel
	colorEnabled bool
}

func NewLogger(w io.Writer) *Logger {
	l := &Logger{
		level:        LogLevelDefault,
		colorEnabled: true,
		Writer:       w,
	}

	if l.Writer == nil {
		l.Writer = log.Writer()
	}

	return l
}

func (l *Logger) SetPrefix(s string) {
	l.prefix = s
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) SetColorEnabled(b bool) {
	if runtime.GOOS == "windows" {
		b = false
	}

	l.colorEnabled = b
}

func (l Logger) Info(msg string) {
	if l.level < LogLevelInfo {
		return
	}

	l.print(LogLevelInfo, msg)
}

func (l Logger) Infof(msg string, args ...interface{}) {
	l.Info(fmt.Sprintf(msg, args...))
}

func (l Logger) Warn(msg string) {
	if l.level < LogLevelWarning {
		return
	}

	l.print(LogLevelWarning, msg)
}

func (l Logger) Warnf(msg string, args ...interface{}) {
	l.Warn(fmt.Sprintf(msg, args...))
}

func (l Logger) Error(msg string) {
	if l.level < LogLevelError {
		return
	}

	l.print(LogLevelError, msg)
}

func (l Logger) Errorf(msg string, args ...interface{}) {
	l.Error(fmt.Sprintf(msg, args...))
}

func (l Logger) Fatal(msg string) {
	if l.level < LogLevelFatal {
		return
	}

	l.print(LogLevelFatal, msg)
}

func (l Logger) Fatalf(msg string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(msg, args...))
}

func (l Logger) Debug(msg string) {
	if l.level < LogLevelDebug {
		return
	}

	l.print(LogLevelDebug, msg)
}

func (l Logger) Debugf(msg string, args ...interface{}) {
	l.Debug(fmt.Sprintf(msg, args...))
}

func (l Logger) Write(p []byte) (n int, err error) {
	return l.Writer.Write(p)
}

func (l Logger) format(fmtStr string, fmtInput []byte) []byte {
	return []byte(fmt.Sprintf(fmtStr, string(fmtInput)))
}

func (l *Logger) print(level LogLevel, msg string) {
	if l == nil || l.level < level {
		return
	}

	fmtString := time.Now().Format(time.RFC3339Nano) + " "

	switch level {
	case LogLevelDebug:
		fmtString += fmt.Sprintf(fmtColorEscape, cyan) + LogFmtDebug + " "
	case LogLevelWarning:
		fmtString += fmt.Sprintf(fmtColorEscape, yellow) + LogFmtWarning + " "
	case LogLevelError:
		fmtString += fmt.Sprintf(fmtColorEscape, red) + LogFmtError + " "
	case LogLevelFatal:
		fmtString += fmt.Sprintf(fmtColorEscape, red) + LogFmtFatal + " "
	case LogLevelInfo:
		fallthrough
	case LogLevelNone:
		fallthrough
	default:
		fmtString += fmt.Sprintf(fmtColorEscape, green) + LogFmtInfo + " "
	}

	if l.prefix != "" {
		fmtString += fmt.Sprintf(fmtColorEscape, magenta) + fmt.Sprintf(fmtPrefix, l.prefix) + colorEscapeReset + " %s\n"
	}

	_, err := l.Write(l.format(fmtString, []byte(msg)))

	if err != nil {
		log.Print(err)
	}
}
