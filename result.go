package exam

func NewResult(e E, ok bool) *Result {
	return &Result{
		e:      e,
		failed: !ok,
	}
}

type Result struct {
	e      E
	failed bool
}

func (r *Result) Log(args ...any) *Result {
	r.e.Helper()
	if r.failed {
		r.e.Log(args...)
	}
	return r
}

func (r *Result) Logf(format string, args ...any) *Result {
	r.e.Helper()
	if r.failed {
		r.e.Logf(format, args...)
	}
	return r
}

func (r *Result) Fatal() bool {
	r.e.Helper()
	if r.failed {
		r.e.FailNow()
	}
	return !r.failed
}

func (r *Result) Ok() bool {
	return !r.failed
}
