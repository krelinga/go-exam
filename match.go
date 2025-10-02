package exam

import "github.com/krelinga/go-match"

func Match[T any](e *E, got T, matcher match.Matcher[T]) {
	matched, explanation := matcher.Match(got)
	if !matched {
		e.Errorf("\n%s", explanation)
	}
}