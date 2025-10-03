package exam

import (
	"context"
	"fmt"
	"io"
	"strings"
	"sync"
)

type Recorder struct {
	mu        sync.RWMutex
	cleanups  []func()
	logs      strings.Builder
	failed    bool
	failNowed bool
	skipped   bool
	name      string
}

func NewRecorder(name string) *Recorder {
	r := &Recorder{
		name: name,
	}
	return r
}

func (r *Recorder) Finish() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, cleanup := range r.cleanups {
		cleanup()
	}
	r.cleanups = nil
}

func (r *Recorder) Cleanup(f func()) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cleanups = append(r.cleanups, f)
}

func (r *Recorder) Error(args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.failed = true
	msg := fmt.Sprint(args...)
	fmt.Fprintf(&r.logs, "ERROR: %s\n", strings.TrimSpace(msg))
}

func (r *Recorder) Errorf(format string, args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.failed = true
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(&r.logs, "ERROR: %s\n", strings.TrimSpace(msg))
}

func (r *Recorder) Fail() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.failed = true
}

func (r *Recorder) FailNow() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.failed = true
	r.failNowed = true
}

func (r *Recorder) Failed() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.failed
}

func (r *Recorder) Fatal(args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.failed = true
	msg := fmt.Sprint(args...)
	fmt.Fprintf(&r.logs, "FATAL: %s\n", strings.TrimSpace(msg))
	r.failNowed = true
}

func (r *Recorder) Fatalf(format string, args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.failed = true
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(&r.logs, "FATAL: %s\n", strings.TrimSpace(msg))
	r.failNowed = true
}

func (r *Recorder) Helper() {
	// No-op for recorder
}

func (r *Recorder) Log(args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	msg := fmt.Sprint(args...)
	fmt.Fprintf(&r.logs, "LOG: %s\n", strings.TrimSpace(msg))
}

func (r *Recorder) Logf(format string, args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(&r.logs, "LOG: %s\n", strings.TrimSpace(msg))
}

func (r *Recorder) Name() string {
	return r.name
}

func (r *Recorder) Skip(args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	msg := fmt.Sprint(args...)
	fmt.Fprintf(&r.logs, "SKIP: %s\n", strings.TrimSpace(msg))
	r.skipped = true
}

func (r *Recorder) SkipNow() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.skipped = true
}

func (r *Recorder) Skipf(format string, args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(&r.logs, "SKIP: %s\n", strings.TrimSpace(msg))
	r.skipped = true
}

func (r *Recorder) Skipped() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.skipped
}

func (r *Recorder) Context() context.Context {
	// No context in this simple recorder
	return context.Background()
}

func (r *Recorder) Output() io.Writer {
	return newSafeWriter(&r.logs, &r.mu)
}

func (r *Recorder) Logs() string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.logs.String()
}

func (r *Recorder) FailNowed() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.failNowed
}
