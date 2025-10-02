package exam_test

import (
	"testing"

	"github.com/krelinga/go-exam"
	"github.com/krelinga/go-match"
)

func TestMatch(t *testing.T) {
	e := exam.New(t)
	e.Run("Match Success", func(e *exam.E) {
		exam.Match(e, 42, match.Equal(42))
	})
	e.Run("Match Failure", func(e *exam.E) {
		if !*manualFlag {
			e.Skip("skipping test in manual mode")
		}
		exam.Match(e, 42, match.Equal(100))
	})
}
