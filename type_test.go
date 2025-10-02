package exam_test

import (
	"flag"
	"testing"

	"github.com/krelinga/go-exam"
)

var manualFlag = flag.Bool("manual", false, "run manual tests")

func TestType(t *testing.T) {
	if !*manualFlag {
		t.Skip("skipping test in manual mode")
	}

	t.Run("WithLoc", func(t *testing.T) {
		e := exam.New(t).WithLoc(exam.Here())
		e.Error("an error with location")
	})

	t.Run("WithMust", func(t *testing.T) {
		e := exam.New(t).WithFatal()
		e.Error("an error with must=true")
		e.Log("this should not appear because of must=true")
	})

	t.Run("Run", func(t *testing.T) {
		e := exam.New(t)
		tests := []struct {
			name string
			loc  exam.Loc
		}{
			{"test1", exam.Here()},
			{"test2", exam.Here()},
		}
		for _, tt := range tests {
			e.WithLoc(tt.loc).Run(tt.name, func(e *exam.E) {
				e.Logf("inside subtest %s with location", tt.name)
			})
		}
	})
}
