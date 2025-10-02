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

func TestMustMatch(t *testing.T) {
	e := exam.New(t)
	e.Run("MustMatch Success", func(e *exam.E) {
		got := exam.MustMatch(e, "hello", match.Equal("hello"))
		if got != "hello" {
			e.Fatalf("expected got to be 'hello', got %q", got)
		}
	})
	e.Run("MustMatch Failure", func(e *exam.E) {
		if !*manualFlag {
			e.Skip("skipping test in manual mode")
		}
		exam.MustMatch(e, "hello", match.NotEqual("hello"))
		e.Fatal("expected MustMatch to fail and stop execution")
	})
}
