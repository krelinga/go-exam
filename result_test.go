package exam_test

import (
	"strings"
	"testing"

	"github.com/krelinga/go-exam"
)

func TestResult(t *testing.T) {
	t.Run("NewResult_Ok", func(t *testing.T) {
		recorder, cleanup := exam.NewRecorder("test")
		defer cleanup()

		result := exam.NewResult(recorder, true)

		if !result.Ok() {
			t.Error("expected result to be ok when created with ok=true")
		}
	})

	t.Run("NewResult_Failed", func(t *testing.T) {
		recorder, cleanup := exam.NewRecorder("test")
		defer cleanup()

		result := exam.NewResult(recorder, false)

		if result.Ok() {
			t.Error("expected result to not be ok when created with ok=false")
		}
	})

	t.Run("Log_WhenFailed", func(t *testing.T) {
		recorder, cleanup := exam.NewRecorder("test")
		defer cleanup()

		result := exam.NewResult(recorder, false)
		returnedResult := result.Log("test message")

		// Check that it returns the same result for chaining
		if returnedResult != result {
			t.Error("expected Log to return the same result for chaining")
		}

		logs := recorder.GetLogs()
		if !strings.Contains(logs, "LOG: test message\n") {
			t.Errorf("expected logs to contain log message when failed, got %q", logs)
		}
	})

	t.Run("Log_WhenOk", func(t *testing.T) {
		recorder, cleanup := exam.NewRecorder("test")
		defer cleanup()

		result := exam.NewResult(recorder, true)
		result.Log("test message")

		logs := recorder.GetLogs()
		if strings.Contains(logs, "test message") {
			t.Errorf("expected logs to not contain message when ok, got %q", logs)
		}
	})

	t.Run("Logf_WhenFailed", func(t *testing.T) {
		recorder, cleanup := exam.NewRecorder("test")
		defer cleanup()

		result := exam.NewResult(recorder, false)
		returnedResult := result.Logf("test message %d", 42)

		// Check that it returns the same result for chaining
		if returnedResult != result {
			t.Error("expected Logf to return the same result for chaining")
		}

		logs := recorder.GetLogs()
		if !strings.Contains(logs, "LOG: test message 42\n") {
			t.Errorf("expected logs to contain formatted log message when failed, got %q", logs)
		}
	})

	t.Run("Logf_WhenOk", func(t *testing.T) {
		recorder, cleanup := exam.NewRecorder("test")
		defer cleanup()

		result := exam.NewResult(recorder, true)
		result.Logf("test message %d", 42)

		logs := recorder.GetLogs()
		if strings.Contains(logs, "test message 42") {
			t.Errorf("expected logs to not contain formatted message when ok, got %q", logs)
		}
	})

	t.Run("Fatal_WhenFailed", func(t *testing.T) {
		recorder, cleanup := exam.NewRecorder("test")
		defer cleanup()

		result := exam.NewResult(recorder, false)
		returnValue := result.Fatal()

		if returnValue {
			t.Error("expected Fatal to return false when result is failed")
		}

		if !recorder.GetFailNowed() {
			t.Error("expected recorder to be fail-nowed after Fatal on failed result")
		}
	})

	t.Run("Fatal_WhenOk", func(t *testing.T) {
		recorder, cleanup := exam.NewRecorder("test")
		defer cleanup()

		result := exam.NewResult(recorder, true)
		returnValue := result.Fatal()

		if !returnValue {
			t.Error("expected Fatal to return true when result is ok")
		}

		if recorder.GetFailNowed() {
			t.Error("expected recorder to not be fail-nowed after Fatal on ok result")
		}
	})

	t.Run("Ok_ReturnsTrueForSuccessfulResult", func(t *testing.T) {
		recorder, cleanup := exam.NewRecorder("test")
		defer cleanup()

		result := exam.NewResult(recorder, true)

		if !result.Ok() {
			t.Error("expected Ok to return true for successful result")
		}
	})

	t.Run("Ok_ReturnsFalseForFailedResult", func(t *testing.T) {
		recorder, cleanup := exam.NewRecorder("test")
		defer cleanup()

		result := exam.NewResult(recorder, false)

		if result.Ok() {
			t.Error("expected Ok to return false for failed result")
		}
	})

	t.Run("Chaining", func(t *testing.T) {
		recorder, cleanup := exam.NewRecorder("test")
		defer cleanup()

		result := exam.NewResult(recorder, false)

		// Test method chaining
		finalResult := result.Log("first message").Logf("second message %s", "test")

		if finalResult != result {
			t.Error("expected chained methods to return the same result")
		}

		logs := recorder.GetLogs()
		if !strings.Contains(logs, "LOG: first message\n") {
			t.Errorf("expected logs to contain first message, got %q", logs)
		}
		if !strings.Contains(logs, "LOG: second message test\n") {
			t.Errorf("expected logs to contain second message, got %q", logs)
		}
	})
}
