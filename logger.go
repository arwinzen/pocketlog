package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information
type Logger struct {
    threshold Level
    output io.Writer
}

// New returns you a logger, ready to log at the required threshold
// Give it a list of configuration functions to tune it to your liking. 
// The default output is StdOut.
func New(threshold Level, opts ...Option) *Logger {
    lgr := &Logger{threshold: threshold, output: os.Stdout}

    for _, configFunc := range opts {
        configFunc(lgr)
    }
    return lgr
}

// logf prints the message to the output
// Add decorations here, if any.
func (l *Logger) logf(format string, args ...any) {
    _, _ = fmt.Printf(format + "\n", args...)
}

// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
    // making sure we can safely write to the output
    if l.output == nil {
        l.output = os.Stdout
    }
    if l.threshold > LevelDebug {
        return 
    }

    l.logf(format, args...)
}

// Info formats and prints a message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any) {
    if l.output == nil {
        l.output = os.Stdout
    }
    if l.threshold > LevelInfo {
        return 
    }

    l.logf(format, args...)
}

// Error formats and prints a message if the log level is error or higher
func (l *Logger) Errorf(format string, args ...any) {
    if l.output == nil {
        l.output = os.Stdout
    }
    if l.threshold > LevelError {
        return 
    }

    l.logf(format, args...)
}
