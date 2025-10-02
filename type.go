package exam

import (
	"fmt"
	"testing"
)

func New(t *testing.T) *E {
	return &E{T: t}
}

type E struct {
	*testing.T
	must bool
	locs []Loc
}

// Wrap testing.T methods

func (e *E) Error(args ...any) {
	e.write(args...)
	e.recordError()
}

func (e *E) Errorf(format string, args ...any) {
	e.writef(format, args...)
	e.recordError()
}

func (e *E) Fail() {
	e.recordError()
}

func (e *E) Fatal(args ...any) {
	e.write(args...)
	e.T.FailNow()
}

func (e *E) Fatalf(format string, args ...any) {
	e.writef(format, args...)
	e.T.FailNow()
}

func (e *E) Log(args ...any) {
	e.write(args...)
}

func (e *E) Logf(format string, args ...any) {
	e.writef(format, args...)
}

func (e *E) Run(name string, f func(e *E)) bool {
	return e.T.Run(name, func(t *testing.T) {
		e = e.clone()
		e.T = t
		f(e)
	})
}

func (e *E) Skip(args ...any) {
	e.write(args...)
	e.T.SkipNow()
}

func (e *E) Skipf(format string, args ...any) {
	e.writef(format, args...)
	e.T.SkipNow()
}

func (e *E) WithMust() *E {
	e = e.clone()
	e.must = true
	return e
}

func (e *E) WithLoc(loc Loc) *E {
	e = e.clone()
	e.locs = append(e.locs, loc)
	return e
}

func (e *E) clone() *E {
	return &E{T: e.T, must: e.must, locs: append([]Loc(nil), e.locs...)}
}

func (e *E) logLocs() {
	for _, l := range e.locs {
		fmt.Fprintln(e.T.Output(), l)
	}
}

func (e *E) recordError() {
	if e.must {
		e.T.FailNow()
	} else {
		e.T.Fail()
	}
}

func (e *E) write(args ...any) {
	e.logLocs()
	x := fmt.Sprint(args...)
	fmt.Fprintf(e.T.Output(), "%s: %s\n", hereOffset(2), x)
}

func (e *E) writef(format string, args ...any) {
	e.logLocs()
	x := fmt.Sprintf(format, args...)
	fmt.Fprintf(e.T.Output(), "%s: %s\n", hereOffset(2), x)
}