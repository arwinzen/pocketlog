package pocketlog_test

import (
	"os"
	"testing"

	"github.com/arwinzen/pocketlog"
)

// testwriter is a struct that implements io.writer
// we use it to validate that we can write to a specific output
type testWriter struct {
    contents string
}

// write implements the io.writer interface
func (tw *testWriter) Write(p []byte) (n int, err error) {
    tw.contents = tw.contents + string(p)
    return len(p), nil
}


func ExampleLogger_Debugf() {
    debugLogger := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(os.Stdout)) 
    debugLogger.Debugf("Hello, %s", "world")
    // Output: Hello, world
}

const (
    debugMessage = "Why write I still all one, ever the same,"
    infoMessage = "And keep invention in a noted weed,"
    errorMessage = "That every word doth almost tell my name,"
)

func TestLogger_DebugfInfofErrorf(t *testing.T) {
    type testCase struct {
        level pocketlog.Level
        expectedMessage string        
    }

    tt := map[string]testCase{
        "Debug": {
            level: pocketlog.LevelDebug,
            expectedMessage: debugMessage + "\n" + infoMessage + "\n" + errorMessage + "\n",
        },
        "Info": {
            level: pocketlog.LevelInfo,
            expectedMessage: infoMessage + "\n" + errorMessage + "\n",
        },
        "Error": {
            level: pocketlog.LevelError,
            expectedMessage: errorMessage + "\n",
        },
    }
    for name, tc := range tt {
        t.Run(name, func(t *testing.T) {
            tw := &testWriter{}

            testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

            testedLogger.Debugf(debugMessage)
            testedLogger.Infof(infoMessage)
            testedLogger.Errorf(errorMessage)

            if tw.contents != tc.expectedMessage {
                t.Errorf("invalid contents, expected %q, got %q", tc.expectedMessage, tw.contents)
            }
        })
    }
}
