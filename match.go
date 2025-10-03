package exam

import "github.com/krelinga/go-match"

func Match[T any](e E, got T, matcher match.Matcher[T]) *Result {
	e.Helper()
	matched, explanation := matcher.Match(got)
	if !matched {
		e.Errorf("\n%s", explanation)
	}
	return NewResult(e, matched)
}
