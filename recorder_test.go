package exam_test

import (
	"context"
	"strings"
	"sync"
	"testing"

	"github.com/krelinga/go-exam"
)

func TestNewRecorder(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test-recorder")
	defer cleanup()

	if recorder.Name() != "test-recorder" {
		t.Errorf("expected name to be 'test-recorder', got %q", recorder.Name())
	}

	if recorder.Failed() {
		t.Error("expected new recorder to not be failed")
	}

	if recorder.Skipped() {
		t.Error("expected new recorder to not be skipped")
	}
}

func TestRecorderCleanup(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")

	var cleanupCalled bool
	recorder.Cleanup(func() {
		cleanupCalled = true
	})

	cleanup()

	if !cleanupCalled {
		t.Error("expected cleanup function to be called")
	}
}

func TestRecorderError(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")
	defer cleanup()

	recorder.Error("test error")

	if !recorder.Failed() {
		t.Error("expected recorder to be failed after Error")
	}

	logs := recorder.GetLogs()
	if !strings.Contains(logs, "ERROR: test error\n") {
		t.Errorf("expected logs to contain error message, got %q", logs)
	}
}

func TestRecorderErrorf(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")
	defer cleanup()

	recorder.Errorf("test error %d", 42)

	if !recorder.Failed() {
		t.Error("expected recorder to be failed after Errorf")
	}

	logs := recorder.GetLogs()
	if !strings.Contains(logs, "ERROR: test error 42\n") {
		t.Errorf("expected logs to contain formatted error message, got %q", logs)
	}
}

func TestRecorderFail(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")
	defer cleanup()

	recorder.Fail()

	if !recorder.Failed() {
		t.Error("expected recorder to be failed after Fail")
	}

	if recorder.GetFailNowed() {
		t.Error("expected recorder to not be fail-nowed after Fail")
	}
}

func TestRecorderFailNow(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")
	defer cleanup()

	recorder.FailNow()

	if !recorder.Failed() {
		t.Error("expected recorder to be failed after FailNow")
	}

	if !recorder.GetFailNowed() {
		t.Error("expected recorder to be fail-nowed after FailNow")
	}
}

func TestRecorderFatal(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")
	defer cleanup()

	recorder.Fatal("fatal error")

	if !recorder.Failed() {
		t.Error("expected recorder to be failed after Fatal")
	}

	if !recorder.GetFailNowed() {
		t.Error("expected recorder to be fail-nowed after Fatal")
	}

	logs := recorder.GetLogs()
	if !strings.Contains(logs, "FATAL: fatal error\n") {
		t.Errorf("expected logs to contain fatal message, got %q", logs)
	}
}

func TestRecorderFatalf(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")
	defer cleanup()

	recorder.Fatalf("fatal error %s", "test")

	if !recorder.Failed() {
		t.Error("expected recorder to be failed after Fatalf")
	}

	if !recorder.GetFailNowed() {
		t.Error("expected recorder to be fail-nowed after Fatalf")
	}

	logs := recorder.GetLogs()
	if !strings.Contains(logs, "FATAL: fatal error test\n") {
		t.Errorf("expected logs to contain formatted fatal message, got %q", logs)
	}
}

func TestRecorderLog(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")
	defer cleanup()

	recorder.Log("test log")

	if recorder.Failed() {
		t.Error("expected recorder to not be failed after Log")
	}

	logs := recorder.GetLogs()
	if !strings.Contains(logs, "LOG: test log\n") {
		t.Errorf("expected logs to contain log message, got %q", logs)
	}
}

func TestRecorderLogf(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")
	defer cleanup()

	recorder.Logf("test log %d", 42)

	if recorder.Failed() {
		t.Error("expected recorder to not be failed after Logf")
	}

	logs := recorder.GetLogs()
	if !strings.Contains(logs, "LOG: test log 42\n") {
		t.Errorf("expected logs to contain formatted log message, got %q", logs)
	}
}

func TestRecorderSkip(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")
	defer cleanup()

	recorder.Skip("skip reason")

	if !recorder.Skipped() {
		t.Error("expected recorder to be skipped after Skip")
	}

	logs := recorder.GetLogs()
	if !strings.Contains(logs, "SKIP: skip reason\n") {
		t.Errorf("expected logs to contain skip message, got %q", logs)
	}
}

func TestRecorderSkipNow(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")
	defer cleanup()

	recorder.SkipNow()

	if !recorder.Skipped() {
		t.Error("expected recorder to be skipped after SkipNow")
	}
}

func TestRecorderSkipf(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")
	defer cleanup()

	recorder.Skipf("skip reason %s", "test")

	if !recorder.Skipped() {
		t.Error("expected recorder to be skipped after Skipf")
	}

	logs := recorder.GetLogs()
	if !strings.Contains(logs, "SKIP: skip reason test\n") {
		t.Errorf("expected logs to contain formatted skip message, got %q", logs)
	}
}

func TestRecorderContext(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")
	defer cleanup()

	ctx := recorder.Context()
	if ctx != context.Background() {
		t.Error("expected context to be background context")
	}
}

func TestRecorderOutput(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")
	defer cleanup()

	output := recorder.Output()
	if output == nil {
		t.Error("expected output to not be nil")
	}

	// Test writing to output
	output.Write([]byte("test output"))
	logs := recorder.GetLogs()
	if !strings.Contains(logs, "test output") {
		t.Errorf("expected logs to contain output, got %q", logs)
	}
}

func TestRecorderConcurrency(t *testing.T) {
	recorder, cleanup := exam.NewRecorder("test")
	defer cleanup()

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

	logs := recorder.GetLogs()
	if logs == "" {
		t.Error("expected logs to contain messages from concurrent operations")
	}
}
