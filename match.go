package exam

import "github.com/krelinga/go-match"

func Match[T any](e *E, got T, matcher match.Matcher[T]) bool {
	e.Helper()
	matched, explanation := matcher.Match(got)
	if !matched {
		e.Errorf("\n%s", explanation)
	}
	return matched
}

func FilterMatch[T any](e *E, got T, matcher match.Matcher[T]) T {
	e.Helper()
	Match(e.WithMust(), got, matcher)
	return got
}