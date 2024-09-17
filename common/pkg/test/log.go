package test

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	ctrl "sigs.k8s.io/controller-runtime"
)

// NewTestLogger returns a logger that writes to the test log.
func NewTestLogger(t *testing.T) logr.Logger {
	stdr.SetVerbosity(8)
	logger := log.New(&testLogWriter{t}, "TEST: ", 0)
	return stdr.New(logger)
}

// ContextWithLogger returns a context with a logger that writes to the test log.
func ContextWithLogger(t *testing.T) context.Context {
	logger := NewTestLogger(t)
	return ctrl.LoggerInto(context.Background(), logger)
}

type testLogWriter struct {
	t *testing.T
}

func (w *testLogWriter) Write(p []byte) (n int, err error) {
	w.t.Log(strings.TrimSpace(string(p)))
	return len(p), nil
}