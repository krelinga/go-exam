package exam_test

import (
	"context"
	"strings"
	"sync"
	"testing"

	"github.com/krelinga/go-exam"
)

func TestRecorder(t *testing.T) {
	t.Run("NewRecorder", func(t *testing.T) {
		recorder := exam.NewRecorder("test-recorder")
		defer recorder.Finish()

		if recorder.Name() != "test-recorder" {
			t.Errorf("expected name to be 'test-recorder', got %q", recorder.Name())
		}

		if recorder.Failed() {
			t.Error("expected new recorder to not be failed")
		}

		if recorder.Skipped() {
			t.Error("expected new recorder to not be skipped")
		}
	})

	t.Run("RecorderCleanup", func(t *testing.T) {
		recorder := exam.NewRecorder("test")

		var cleanupCalled bool
		recorder.Cleanup(func() {
			cleanupCalled = true
		})

		recorder.Finish()

		if !cleanupCalled {
			t.Error("expected cleanup function to be called")
		}
	})

	t.Run("RecorderError", func(t *testing.T) {
		recorder := exam.NewRecorder("test")
		defer recorder.Finish()

		recorder.Error("test error")

		if !recorder.Failed() {
			t.Error("expected recorder to be failed after Error")
		}

		logs := recorder.Logs()
		if !strings.Contains(logs, "ERROR: test error\n") {
			t.Errorf("expected logs to contain error message, got %q", logs)
		}
	})

	t.Run("RecorderErrorf", func(t *testing.T) {
		recorder := exam.NewRecorder("test")
		defer recorder.Finish()

		recorder.Errorf("test error %d", 42)

		if !recorder.Failed() {
			t.Error("expected recorder to be failed after Errorf")
		}

		logs := recorder.Logs()
		if !strings.Contains(logs, "ERROR: test error 42\n") {
			t.Errorf("expected logs to contain formatted error message, got %q", logs)
		}
	})

	t.Run("RecorderFail", func(t *testing.T) {
		recorder := exam.NewRecorder("test")
		defer recorder.Finish()

		recorder.Fail()

		if !recorder.Failed() {
			t.Error("expected recorder to be failed after Fail")
		}

		if recorder.FailNowed() {
			t.Error("expected recorder to not be fail-nowed after Fail")
		}
	})

	t.Run("RecorderFailNow", func(t *testing.T) {
		recorder := exam.NewRecorder("test")
		defer recorder.Finish()

		recorder.FailNow()

		if !recorder.Failed() {
			t.Error("expected recorder to be failed after FailNow")
		}

		if !recorder.FailNowed() {
			t.Error("expected recorder to be fail-nowed after FailNow")
		}
	})

	t.Run("RecorderFatal", func(t *testing.T) {
		recorder := exam.NewRecorder("test")
		defer recorder.Finish()

		recorder.Fatal("fatal error")

		if !recorder.Failed() {
			t.Error("expected recorder to be failed after Fatal")
		}

		if !recorder.FailNowed() {
			t.Error("expected recorder to be fail-nowed after Fatal")
		}

		logs := recorder.Logs()
		if !strings.Contains(logs, "FATAL: fatal error\n") {
			t.Errorf("expected logs to contain fatal message, got %q", logs)
		}
	})

	t.Run("RecorderFatalf", func(t *testing.T) {
		recorder := exam.NewRecorder("test")
		defer recorder.Finish()

		recorder.Fatalf("fatal error %s", "test")

		if !recorder.Failed() {
			t.Error("expected recorder to be failed after Fatalf")
		}

		if !recorder.FailNowed() {
			t.Error("expected recorder to be fail-nowed after Fatalf")
		}

		logs := recorder.Logs()
		if !strings.Contains(logs, "FATAL: fatal error test\n") {
			t.Errorf("expected logs to contain formatted fatal message, got %q", logs)
		}
	})

	t.Run("RecorderLog", func(t *testing.T) {
		recorder := exam.NewRecorder("test")
		defer recorder.Finish()

		recorder.Log("test log")

		if recorder.Failed() {
			t.Error("expected recorder to not be failed after Log")
		}

		logs := recorder.Logs()
		if !strings.Contains(logs, "LOG: test log\n") {
			t.Errorf("expected logs to contain log message, got %q", logs)
		}
	})

	t.Run("RecorderLogf", func(t *testing.T) {
		recorder := exam.NewRecorder("test")
		defer recorder.Finish()

		recorder.Logf("test log %d", 42)

		if recorder.Failed() {
			t.Error("expected recorder to not be failed after Logf")
		}

		logs := recorder.Logs()
		if !strings.Contains(logs, "LOG: test log 42\n") {
			t.Errorf("expected logs to contain formatted log message, got %q", logs)
		}
	})

	t.Run("RecorderSkip", func(t *testing.T) {
		recorder := exam.NewRecorder("test")
		defer recorder.Finish()

		recorder.Skip("skip reason")

		if !recorder.Skipped() {
			t.Error("expected recorder to be skipped after Skip")
		}

		logs := recorder.Logs()
		if !strings.Contains(logs, "SKIP: skip reason\n") {
			t.Errorf("expected logs to contain skip message, got %q", logs)
		}
	})

	t.Run("RecorderSkipNow", func(t *testing.T) {
		recorder := exam.NewRecorder("test")
		defer recorder.Finish()

		recorder.SkipNow()

		if !recorder.Skipped() {
			t.Error("expected recorder to be skipped after SkipNow")
		}
	})

	t.Run("RecorderSkipf", func(t *testing.T) {
		recorder := exam.NewRecorder("test")
		defer recorder.Finish()

		recorder.Skipf("skip reason %s", "test")

		if !recorder.Skipped() {
			t.Error("expected recorder to be skipped after Skipf")
		}

		logs := recorder.Logs()
		if !strings.Contains(logs, "SKIP: skip reason test\n") {
			t.Errorf("expected logs to contain formatted skip message, got %q", logs)
		}
	})

	t.Run("RecorderContext", func(t *testing.T) {
		recorder := exam.NewRecorder("test")
		defer recorder.Finish()

		ctx := recorder.Context()
		if ctx != context.Background() {
			t.Error("expected context to be background context")
		}
	})

	t.Run("RecorderOutput", func(t *testing.T) {
		recorder := exam.NewRecorder("test")
		defer recorder.Finish()

		output := recorder.Output()
		if output == nil {
			t.Error("expected output to not be nil")
		}

		// Test writing to output
		output.Write([]byte("test output"))
		logs := recorder.Logs()
		if !strings.Contains(logs, "test output") {
			t.Errorf("expected logs to contain output, got %q", logs)
		}
	})

	t.Run("RecorderConcurrency", func(t *testing.T) {
		recorder := exam.NewRecorder("test")
		defer recorder.Finish()

		var wg sync.WaitGroup
		numGoroutines := 10

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				recorder.Logf("goroutine %d", id)
				recorder.Error("error from goroutine")
				recorder.Skip("skip from goroutine")
			}(i)
		}

		wg.Wait()

		if !recorder.Failed() {
			t.Error("expected recorder to be failed after concurrent errors")
		}

		if !recorder.Skipped() {
			t.Error("expected recorder to be skipped after concurrent skips")
		}

		logs := recorder.Logs()
		if logs == "" {
			t.Error("expected logs to contain messages from concurrent operations")
		}
	})
}
