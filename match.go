package exam

import "github.com/krelinga/go-match"

func Match[T any](e *E, got T, matcher match.Matcher[T]) bool {
	matched, explanation := matcher.Match(got)
	if !matched {
		e.writef("\n%s", explanation)
	}
	return matched
}

func FilterMatch[T any](e *E, got T, matcher match.Matcher[T]) T {
	matched, explanation := matcher.Match(got)
	if !matched {
		e.writef("\n%s", explanation)
		e.T.FailNow()
	}
	return got
}