package pocketlog_test

import (
	"github.com/arwinzen/pocketlog"
)

func ExampleLogger_Debugf() {
    debugLogger := pocketlog.New(pocketlog.LevelDebug, nil) 
    debugLogger.Debugf("Hello, %s", "world")
    // Output: Hello, world
}
