package pocketlog

import "io"

// Option defines an functional option to our Logger
type Option func(*Logger)

// WithOutput returns a configuration function that sets the output of the logs
func WithOutput(output io.Writer) Option {
    return func(lgr *Logger) {
        lgr.output = output
    }
}
