package exam_test

import (
	"testing"

	"github.com/krelinga/go-exam"
	"github.com/krelinga/go-match"
)

func TestMatch(t *testing.T) {
	e := exam.New(t)
	e.Run("Match Success", func(e *exam.E) {
		got := exam.Match(e, 42, match.Equal(42))
		if !got {
			e.Fatal("expected match to succeed")
		}
	})
	e.Run("Match Failure", func(e *exam.E) {
		if !*manualFlag {
			e.Skip("skipping test in manual mode")
		}
		got := exam.Match(e, 42, match.Equal(100))
		if got {
			e.Fatal("expected match to fail")
		}
	})
}

func TestFilterMatch(t *testing.T) {
	e := exam.New(t)
	e.Run("FilterMatch Success", func(e *exam.E) {
		got := exam.FilterMatch(e, "hello", match.Equal("hello"))
		if got != "hello" {
			e.Fatalf("expected got to be 'hello', got %q", got)
		}
	})
	e.Run("FilterMatch Failure", func(e *exam.E) {
		if !*manualFlag {
			e.Skip("skipping test in manual mode")
		}
		exam.FilterMatch(e, "hello", match.NotEqual("hello"))
		e.Fatal("expected FilterMatch to fail and stop execution")
	})
}
