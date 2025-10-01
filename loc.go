package exam

import (
	"fmt"
	"runtime"
)

// Loc represents a line of code in a file.
type Loc struct {
	File string
	Line int
}

// String returns a readable string representation of the location.
func (l Loc) String() string {
	if l.File == "" && l.Line == 0 {
		return "<uninitialized>"
	}
	return fmt.Sprintf("%s:%d", l.File, l.Line)
}

// Here returns a Loc instance with the file and line number where Here() was called from.
func Here() Loc {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return Loc{File: "unknown", Line: 0}
	}
	return Loc{File: file, Line: line}
}
