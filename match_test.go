package exam_test

import (
	"testing"

	"github.com/krelinga/go-exam"
	"github.com/krelinga/go-match"
)

func TestMatch(t *testing.T) {
	t.Run("Match_Success", func(t *testing.T) {
		recorder, cleanup := exam.NewRecorder("test")
		defer cleanup()

		matcher := match.Equal(42)
		result := exam.Match(recorder, 42, matcher)

		if !result.Ok() {
			t.Error("expected match to be successful for equal values")
		}
	})

	t.Run("Match_Failure", func(t *testing.T) {
		recorder, cleanup := exam.NewRecorder("test")
		defer cleanup()

		matcher := match.Equal(100)
		result := exam.Match(recorder, 42, matcher)

		if result.Ok() {
			t.Error("expected match to fail for unequal values")
		}
	})
}
