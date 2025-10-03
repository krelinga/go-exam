package exam

import "testing"

func NewResult(t *testing.T, ok bool) *Result {
	return &Result{
		t: t,
		failed: !ok,
	}
}

type Result struct {
	t      *testing.T
	failed bool
}

func (r *Result) Log(args ...any) *Result {
	if r.failed {
		r.t.Log(args...)
	}
	return r
}

func (r *Result) Logf(format string, args ...any) *Result {
	if r.failed {
		r.t.Logf(format, args...)
	}
	return r
}

func (r *Result) Fatal() bool {
	if r.failed {
		r.t.FailNow()
	}
	return !r.failed
}

func (r *Result) Ok() bool {
	return !r.failed
}