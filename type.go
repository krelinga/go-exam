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
	e.T.Helper()
	e.Log(args...)
	e.recordError()
}

func (e *E) Errorf(format string, args ...any) {
	e.T.Helper()
	e.Logf(format, args...)
	e.recordError()
}

func (e *E) Fail() {
	e.T.Helper()
	e.recordError()
}

func (e *E) Fatal(args ...any) {
	e.T.Helper()
	e.T.Log(args...)
	e.T.FailNow()
}

func (e *E) Fatalf(format string, args ...any) {
	e.T.Helper()
	e.T.Logf(format, args...)
	e.T.FailNow()
}

func (e *E) Log(args ...any) {
	e.T.Helper()
	e.logLocs()
	e.T.Log(args...)
}

func (e *E) Logf(format string, args ...any) {
	e.T.Helper()
	e.logLocs()
	e.T.Logf(format, args...)
}

func (e *E) Run(name string, f func(e *E)) bool {
	return e.T.Run(name, func(t *testing.T) {
		e = e.clone()
		e.T = t
		f(e)
	})
}

func (e *E) Skip(args ...any) {
	e.T.Helper()
	e.T.Log(args...)
	e.T.SkipNow()
}

func (e *E) Skipf(format string, args ...any) {
	e.T.Helper()
	e.T.Logf(format, args...)
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
