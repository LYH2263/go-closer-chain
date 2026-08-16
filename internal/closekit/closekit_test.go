package closekit

import (
	"errors"
	"testing"
)

type countingCloser struct {
	err   error
	calls *int
}

func (c countingCloser) Close() error {
	*c.calls++
	return c.err
}

func TestCloseAllContinues(t *testing.T) {
	var n int
	boom := errors.New("boom")
	err := CloseAll(
		countingCloser{err: boom, calls: &n},
		countingCloser{err: nil, calls: &n},
	)
	if !errors.Is(err, boom) {
		t.Fatalf("got %v", err)
	}
	if n != 2 {
		t.Fatalf("closes=%d want 2", n)
	}
}
