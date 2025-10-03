package exam

import (
	"io"
	"sync"
)

type safeWriter struct {
	mu *sync.RWMutex
	w  io.Writer
}

func newSafeWriter(w io.Writer, mu *sync.RWMutex) *safeWriter {
	return &safeWriter{
		mu: mu,
		w:  w,
	}
}

func (sw *safeWriter) Write(p []byte) (n int, err error) {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	return sw.w.Write(p)
}
